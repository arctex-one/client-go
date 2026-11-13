package publicapi

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	wire "github.com/arctex-one/client-go/generated"
)

func TestVerifyCallbackWithPublicKey(t *testing.T) {
	key, private, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatal(err)
	}
	kid := CallbackKeyID(key)
	now := time.Date(2026, 11, 13, 12, 0, 0, 0, time.UTC)
	body := []byte("{\n  \"event_id\": \"event\"\n}")
	jwk := wire.WebhookVerificationKey{Kty: "OKP", Crv: "Ed25519", Alg: "EdDSA", Use: "sig", Kid: kid, X: base64.RawURLEncoding.EncodeToString(key)}
	headers := func(alg string, at time.Time, suffix string) http.Header {
		ts := at.Format(time.RFC3339Nano)
		raw := `{"alg":"` + alg + `","kid":"` + kid + `","b64":false,"crit":["b64"],"ts":"` + ts + `"` + suffix + "}"
		protected := base64.RawURLEncoding.EncodeToString([]byte(raw))
		signed := ed25519.Sign(private, append([]byte(protected+"."), body...))
		h := make(http.Header)
		h.Set("X-Callback-Algorithm", "jws+eddsa-ed25519")
		h.Set("X-Callback-JWS", protected+".."+base64.RawURLEncoding.EncodeToString(signed))
		h.Set("X-Callback-Key-Id", kid)
		h.Set("X-Callback-Timestamp", ts)
		return h
	}
	t.Run("lookup and exact wire bytes", func(t *testing.T) {
		calls := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			calls++
			if r.Method != http.MethodGet || r.URL.Path != "/v1/webhook-verification-keys/"+kid || r.Header.Get("Authorization") != "Bearer fixture" {
				t.Error("unexpected authenticated lookup")
			}
			_ = json.NewEncoder(w).Encode(jwk)
		}))
		defer server.Close()
		client := &wire.Client{BaseURL: server.URL, Bearer: "fixture"}
		if err := VerifyCallbackWithPublicKey(t.Context(), body, headers("EdDSA", now, ""), client, now); err != nil {
			t.Fatal(err)
		}
		tampered := []byte(strings.ReplaceAll(string(body), "  ", " "))
		if err := VerifyCallbackWithPublicKey(t.Context(), tampered, headers("EdDSA", now, ""), client, now); !errors.Is(err, ErrInvalidCallback) {
			t.Fatalf("modified bytes accepted: %v", err)
		}
		if calls != 2 {
			t.Fatalf("lookups = %d", calls)
		}
	})

	for _, tc := range []struct {
		name string
		head http.Header
	}{
		{"stale", headers("EdDSA", now.Add(-CallbackVerificationWindow-time.Nanosecond), "")},
		{"future", headers("EdDSA", now.Add(time.Nanosecond), "")},
		{"wrong algorithm", headers("HS256", now, "")},
		{"duplicate field", headers("EdDSA", now, `,"alg":"EdDSA"`)},
		{"key URL", headers("EdDSA", now, `,"jku":"https://untrusted.invalid/key"`)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { t.Error("invalid header caused lookup") }))
			defer server.Close()
			client := &wire.Client{BaseURL: server.URL, Bearer: "fixture"}
			if err := VerifyCallbackWithPublicKey(t.Context(), body, tc.head, client, now); !errors.Is(err, ErrInvalidCallback) {
				t.Fatalf("got %v", err)
			}
		})
	}

	for _, tc := range []struct {
		name string
		edit func(*wire.WebhookVerificationKey)
	}{
		{"expired", func(key *wire.WebhookVerificationKey) {
			expiry := now.Format(time.RFC3339Nano)
			key.ExpiresAt = &expiry
		}},
		{"different public material", func(key *wire.WebhookVerificationKey) {
			key.X = base64.RawURLEncoding.EncodeToString(make([]byte, ed25519.PublicKeySize))
		}},
		{"wrong use", func(key *wire.WebhookVerificationKey) { key.Use = "enc" }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			reply := jwk
			tc.edit(&reply)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { _ = json.NewEncoder(w).Encode(reply) }))
			defer server.Close()
			if err := VerifyCallbackWithPublicKey(t.Context(), body, headers("EdDSA", now, ""), &wire.Client{BaseURL: server.URL, Bearer: "fixture"}, now); !errors.Is(err, ErrInvalidCallback) {
				t.Fatalf("got %v", err)
			}
		})
	}
}

func TestWebhookVerificationKeyLookupRefusesRedirect(t *testing.T) {
	destination := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Error("lookup followed redirect")
	}))
	defer destination.Close()
	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, destination.URL, http.StatusTemporaryRedirect)
	}))
	defer origin.Close()
	client := &wire.Client{BaseURL: origin.URL, Bearer: "fixture", HTTP: origin.Client()}
	_, err := client.GetWebhookVerificationKey(context.Background(), "fixture", nil)
	var response *wire.ResponseError
	if !errors.As(err, &response) || response.Status != http.StatusTemporaryRedirect {
		t.Fatalf("expected HTTP 307, got %v", err)
	}
}
