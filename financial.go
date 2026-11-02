// Package publicapi contains consumer guards for the generated draft wire client.
package publicapi

import (
	"encoding/json"
	"errors"
	"fmt"
	wire "github.com/arctex-one/client-go/generated"
	"regexp"
	"time"
)

var ErrUnrecognizedFinancialValue = errors.New("unrecognized financial value; retain raw evidence and reconcile before applying it")
var decimalString = regexp.MustCompile(`^-?(0|[1-9][0-9]*)(\.[0-9]+)?$`)

// ReadEffect preserves the full generated representation even when a new enum
// value prevents interpretation. Callers must retain the raw record and error;
// they must not post a guessed/default financial outcome. This checks wire
// interpretation only, never authorizes posting or proves source completeness.
func ReadEffect(raw []byte) (*wire.FinancialEffect, error) {
	var effect wire.FinancialEffect
	if err := json.Unmarshal(raw, &effect); err != nil {
		return nil, err
	}
	unknown := func(field, value string) (*wire.FinancialEffect, error) {
		return &effect, fmt.Errorf("%w: %s=%q", ErrUnrecognizedFinancialValue, field, value)
	}
	switch effect.LineType {
	case "MAIN", "FEE", "SPREAD", "REVERSAL":
	default:
		return unknown("line_type", effect.LineType)
	}
	switch effect.Asset.Kind {
	case "NATIVE", "TOKEN_CONTRACT":
	default:
		return unknown("asset.kind", effect.Asset.Kind)
	}
	var origin struct {
		Type     string `json:"type"`
		Deposit  string `json:"deposit_ref"`
		Transfer string `json:"transfer_ref"`
		Reason   string `json:"reason_code"`
	}
	if err := json.Unmarshal(effect.Origin, &origin); err != nil {
		return &effect, err
	}
	validOrigin := false
	switch origin.Type {
	case "DEPOSIT":
		validOrigin = origin.Deposit != ""
	case "TRANSFER":
		validOrigin = origin.Transfer != ""
	case "ADJUSTMENT":
		validOrigin = origin.Reason != ""
	default:
		return unknown("origin.type", origin.Type)
	}
	if effect.EffectRef == "" || effect.AccountRef == "" || effect.PostingLineRef == "" || effect.JournalEntryRef == "" || !validOrigin || effect.Asset.Network == "" || effect.Asset.Currency == "" || effect.Money.Currency != effect.Asset.Currency || !decimalString.MatchString(effect.Money.Amount) {
		return &effect, errors.New("incomplete financial evidence")
	}
	if _, err := time.Parse(time.RFC3339Nano, effect.CommittedAt); err != nil {
		return &effect, err
	}
	return &effect, nil
}
