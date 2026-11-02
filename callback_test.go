package publicapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"testing"
	"time"
)

func TestCallbackVerificationPreservesExactBytesAndBoundsReplay(t *testing.T) {
	now := time.Date(2026, 10, 26, 12, 0, 0, 123, time.UTC)
	body := []byte(`{"event_id":"event", "delivery_id":"attempt","type":"financial_effect.committed","data":{"effect_ref":"effect"}}`)
	key := []byte("isolated-consumer-fixture-secret")
	signed := func(at time.Time) http.Header {
		timestamp := at.Format(time.RFC3339Nano)
		mac := hmac.New(sha256.New, key)
		_, _ = mac.Write(append([]byte(timestamp+"."), body...))
		h := make(http.Header)
		h.Set("X-Callback-Timestamp", timestamp)
		h.Set("X-Callback-Algorithm", "hmac-sha256")
		h.Set("X-Callback-Signature", hex.EncodeToString(mac.Sum(nil)))
		return h
	}
	valid := signed(now)
	if err := VerifyCallback(body, valid, key, now, time.Minute); err != nil {
		t.Fatal(err)
	}
	type testCase struct {
		name      string
		body, key []byte
		headers   http.Header
		clock     time.Time
		window    time.Duration
	}
	tests := []testCase{
		{"changed raw body", append(append([]byte(nil), body...), '\n'), key, valid, now, time.Minute},
		{"wrong key", body, []byte("other-key"), valid, now, time.Minute},
		{"empty key", body, nil, valid, now, time.Minute},
		{"expired", body, key, signed(now.Add(-time.Minute - time.Nanosecond)), now, time.Minute},
		{"future", body, key, signed(now.Add(time.Minute + time.Nanosecond)), now, time.Minute},
		{"no window", body, key, valid, now, 0},
		{"no clock", body, key, valid, time.Time{}, time.Minute},
	}
	for _, field := range []string{"X-Callback-Timestamp", "X-Callback-Algorithm", "X-Callback-Signature"} {
		h := valid.Clone()
		h.Add(field, h.Get(field))
		tests = append(tests, testCase{"duplicate " + field, body, key, h, now, time.Minute})
		h = valid.Clone()
		h.Del(field)
		tests = append(tests, testCase{"missing " + field, body, key, h, now, time.Minute})
	}
	type headerValue struct{ field, value string }
	for _, replacement := range []headerValue{
		{"X-Callback-Algorithm", "none"},
		{"X-Callback-Timestamp", "invalid-time"},
		{"X-Callback-Signature", "not-hex"},
		{"X-Callback-Signature", "00"},
	} {
		h := valid.Clone()
		h.Set(replacement.field, replacement.value)
		tests = append(tests, testCase{"invalid " + replacement.field + replacement.value, body, key, h, now, time.Minute})
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := VerifyCallback(test.body, test.headers, test.key, test.clock, test.window); !errors.Is(err, ErrInvalidCallback) {
				t.Fatal("unauthenticated callback accepted")
			}
		})
	}
}
