package publicapi

import (
	"encoding/json"
	"errors"
	wire "github.com/arctex-one/client-go/generated"
	"testing"
)

func TestReadEffectQuarantinesUnknownValuesWithoutLosingEvidence(t *testing.T) {
	base := wire.FinancialEffect{EffectRef: "effect", PostingLineRef: "posting", JournalEntryRef: "journal", AccountRef: "account", Money: wire.Money{Amount: "100.000000000000000001", Currency: "USDC"}, Asset: wire.Asset{Currency: "USDC", Network: "ethereum_sepolia", Kind: "TOKEN_CONTRACT"}, Origin: json.RawMessage(`{"type":"TRANSFER","transfer_ref":"transfer"}`), LineType: "MAIN", CommittedAt: "2026-10-25T00:00:00Z"}
	raw, _ := json.Marshal(base)
	known, err := ReadEffect(raw)
	if err != nil || known.Money.Amount != base.Money.Amount {
		t.Fatalf("exact known financial evidence lost: %v", err)
	}
	for _, change := range []func(*wire.FinancialEffect){func(e *wire.FinancialEffect) { e.LineType = "FUTURE_LINE" }, func(e *wire.FinancialEffect) { e.Asset.Kind = "FUTURE_ASSET" }, func(e *wire.FinancialEffect) {
		e.Origin = json.RawMessage(`{"type":"FUTURE_ORIGIN","future_ref":"new"}`)
	}} {
		copy := base
		change(&copy)
		copy.AdditionalFields = map[string]json.RawMessage{"future_field": json.RawMessage(`{"retained":true}`)}
		raw, _ = json.Marshal(copy)
		retained, err := ReadEffect(raw)
		if !errors.Is(err, ErrUnrecognizedFinancialValue) || retained == nil || len(retained.AdditionalFields["future_field"]) == 0 {
			t.Fatal("unknown financial values must halt interpretation and retain full evidence")
		}
	}
}
