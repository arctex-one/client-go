package publicapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	wire "github.com/arctex-one/client-go/generated"
)

// These portable consumer checks run both here and in the exported SDK with no
// MC modules/services available. Server concurrency/recovery has separate gates.
func TestClientKeepsOriginalSubmissionAndPublicErrorEvidence(t *testing.T) {
	var mu sync.Mutex
	var first []byte
	var calls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/transfers" || r.Method != http.MethodPost || r.Header.Get("Idempotency-Key") != "persisted-before-post" || r.Header.Get("Authorization") != "Bearer test-only" {
			t.Error("client changed submission authority/identity")
			w.WriteHeader(400)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}
		mu.Lock()
		defer mu.Unlock()
		calls++
		if calls == 1 {
			first = body
			w.WriteHeader(503)
			_, _ = w.Write([]byte(`{"code":"FUTURE_DEPENDENCY","detail":"outcome uncertain"}`))
			return
		}
		if !bytes.Equal(body, first) {
			t.Error("client changed original request on retry")
		}
		w.WriteHeader(202)
		_, _ = w.Write([]byte(`{"transfer_ref":"same-transfer","version":"18446744073709551615","state":"FUTURE_STATE","future_field":{"retained":true}}`))
	}))
	defer server.Close()
	client := wire.Client{BaseURL: server.URL, Bearer: "test-only", HTTP: server.Client()}
	reference := "funding-obligation-original"
	request := wire.SubmitTransfer{QuoteRef: "original-quote", ClientReference: &reference}
	_, err := client.SubmitTransfer(t.Context(), nil, "persisted-before-post", request)
	var response *wire.ResponseError
	if !errors.As(err, &response) || response.Status != 503 || !bytes.Contains(response.Body, []byte("FUTURE_DEPENDENCY")) {
		t.Fatalf("public error evidence lost: %v", err)
	}
	mu.Lock()
	observedCalls := calls
	mu.Unlock()
	if observedCalls != 1 {
		t.Fatal("client must return a temporary HTTP response without an application-level retry")
	}
	result, err := client.SubmitTransfer(t.Context(), nil, "persisted-before-post", request)
	if err != nil {
		t.Fatal(err)
	}
	if result.TransferRef != "same-transfer" || result.Version != "18446744073709551615" || result.State != "FUTURE_STATE" || len(result.AdditionalFields["future_field"]) == 0 {
		t.Fatal("forward-compatible response evidence lost")
	}
}

func TestClientPreservesUnpricedQuotesPresenceAndDecimalSpellings(t *testing.T) {
	var quote wire.Quote
	raw := []byte(`{"quote_ref":"blocked","state":"BLOCKED","readiness":"LIQUIDITY_UNAVAILABLE","block_reason":"provider outage","settlement_mode":"FIX_SOURCE","fee_treatment":"ADD_TO_SOURCE","instructions":{"source_account_ref":"account","source_asset":{"network":"monad_testnet","currency":"USDC","kind":"TOKEN_CONTRACT","contract_address":"0x3333333333333333333333333333333333333333"},"destination":{"asset":{"network":"ethereum_sepolia","currency":"USDC","kind":"TOKEN_CONTRACT","contract_address":"0x4444444444444444444444444444444444444444"},"address":"0x2222222222222222222222222222222222222222"}}}`)
	if err := json.Unmarshal(raw, &quote); err != nil {
		t.Fatal(err)
	}
	if quote.Principal != nil || quote.PayerTotalDebit != nil || quote.RecipientAmount != nil || quote.Fees != nil || quote.ExpiresAt != nil || quote.PricedAt != nil {
		t.Fatal("unpriced economics became executable zero amounts")
	}
	out, err := json.Marshal(quote)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(out, []byte(`"principal"`)) || bytes.Contains(out, []byte(`"fees"`)) {
		t.Fatal("omitted quote economics must remain omitted")
	}
	for _, amount := range []string{"100", "100.00", "0.000000000000000001"} {
		request := wire.QuoteRequest{Amount: amount}
		body, err := json.Marshal(request)
		if err != nil {
			t.Fatal(err)
		}
		var decoded wire.QuoteRequest
		if err = json.Unmarshal(body, &decoded); err != nil {
			t.Fatal(err)
		}
		if decoded.Amount != amount || decoded.SourceAsset.ContractAddress != nil {
			t.Fatal("decimal spelling or field presence changed")
		}
	}
	for _, origin := range []string{`{"type":"DEPOSIT","deposit_ref":"deposit"}`, `{"type":"TRANSFER","transfer_ref":"transfer"}`, `{"type":"ADJUSTMENT","reason_code":"reviewed_correction"}`} {
		effect := wire.FinancialEffect{Origin: json.RawMessage(origin), Money: wire.Money{Amount: "-100.00", Currency: "USDC"}}
		out, err := json.Marshal(effect)
		if err != nil {
			t.Fatal(err)
		}
		var recovered wire.FinancialEffect
		if err = json.Unmarshal(out, &recovered); err != nil {
			t.Fatal(err)
		}
		if recovered.Money.Amount != "-100.00" || !bytes.Equal(recovered.Origin, []byte(origin)) {
			t.Fatal("financial origin or decimal amount changed")
		}
	}
}

func TestClientDoesNotRedirectAuthenticatedSubmissions(t *testing.T) {
	for _, code := range []int{301, 302, 303, 307, 308} {
		t.Run(http.StatusText(code), func(t *testing.T) {
			var forwarded atomic.Int32
			target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				forwarded.Add(1)
				_, _ = w.Write([]byte(`{"transfer_ref":"redirected-operation"}`))
			}))
			defer target.Close()
			origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Location", target.URL+"/v1/transfers")
				w.WriteHeader(code)
			}))
			defer origin.Close()
			for _, configured := range []bool{false, true} {
				client := wire.Client{BaseURL: origin.URL, Bearer: "test-only"}
				var policyCalls atomic.Int32
				if configured {
					client.HTTP = origin.Client()
					client.HTTP.CheckRedirect = func(*http.Request, []*http.Request) error { policyCalls.Add(1); return nil }
				}
				_, err := client.SubmitTransfer(t.Context(), nil, "original-key", wire.SubmitTransfer{QuoteRef: "original-quote"})
				var response *wire.ResponseError
				if !errors.As(err, &response) || response.Status != code {
					t.Errorf("redirect must remain an explicit response: %v", err)
				}
				if forwarded.Load() != 0 {
					t.Error("submission or tenant credentials reached the redirect target")
				}
				if policyCalls.Load() != 0 {
					t.Error("SDK must refuse redirects even if a shared client permits them")
				}
				if configured {
					// A shallow client copy must not overwrite the caller's shared policy.
					before := policyCalls.Load()
					_ = client.HTTP.CheckRedirect(nil, nil)
					if policyCalls.Load() != before+1 {
						t.Error("SDK mutated the supplied HTTP client")
					}
				}
			}
		})
	}
}

func TestClientRejectsOversizedResponseEvenWithValidJSONPrefix(t *testing.T) {
	const limit = 4 << 20
	type responseCase struct {
		name      string
		status    int
		oversized bool
	}
	for _, test := range []responseCase{{"oversized accepted", 202, true}, {"oversized dependency failure", 503, true}, {"exact limit accepted", 202, false}, {"exact limit dependency failure", 503, false}} {
		t.Run(test.name, func(t *testing.T) {
			prefix := `{"transfer_ref":"bounded-JSON-prefix"}`
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.status)
				_, _ = io.WriteString(w, prefix)
				_, _ = io.WriteString(w, strings.Repeat(" ", limit-len(prefix)))
				if test.oversized {
					_, _ = io.WriteString(w, "unread invalid suffix")
				}
			}))
			defer server.Close()
			client := wire.Client{BaseURL: server.URL, HTTP: server.Client(), Bearer: "test-only"}
			_, err := client.SubmitTransfer(t.Context(), nil, "original-key", wire.SubmitTransfer{QuoteRef: "original-quote"})
			var response *wire.ResponseError
			if test.oversized {
				if !errors.Is(err, wire.ErrResponseTooLarge) || !errors.As(err, &response) || !response.Truncated || response.Status != test.status || len(response.Body) != limit {
					t.Fatalf("oversized evidence must be marked and never parsed as success: %v", err)
				}
			} else if test.status < 300 {
				if err != nil {
					t.Fatalf("a complete response exactly at the limit was rejected: %v", err)
				}
			} else if !errors.As(err, &response) || response.Truncated || response.Status != test.status || len(response.Body) != limit {
				t.Fatalf("complete error evidence was incorrectly marked truncated: %v", err)
			}
		})
	}
}
