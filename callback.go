package publicapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"time"
)

// ErrInvalidCallback covers authentication failures without exposing key or body.
var ErrInvalidCallback = errors.New("invalid callback authentication")

// VerifyCallback authenticates the original, bounded HTTP body before decoding.
// The receiving application supplies its clock/replay window and Vault-resolved
// key. Persist/deduplicate event and effect identities separately: a valid retry
// has a new signing time and delivery identity. This does not authorize posting.
func VerifyCallback(body []byte, headers http.Header, key []byte, now time.Time, window time.Duration) error {
	if len(key) == 0 || now.IsZero() || window <= 0 {
		return ErrInvalidCallback
	}
	single := func(name string) string {
		values := headers.Values(name)
		if len(values) != 1 {
			return ""
		}
		return values[0]
	}
	timestamp := single("X-Callback-Timestamp")
	if single("X-Callback-Algorithm") != "hmac-sha256" {
		return ErrInvalidCallback
	}
	signedAt, err := time.Parse(time.RFC3339Nano, timestamp)
	if err != nil || signedAt.Before(now.Add(-window)) || signedAt.After(now.Add(window)) {
		return ErrInvalidCallback
	}
	signature, err := hex.DecodeString(single("X-Callback-Signature"))
	if err != nil || len(signature) != sha256.Size {
		return ErrInvalidCallback
	}
	mac := hmac.New(sha256.New, key)
	_, _ = mac.Write([]byte(timestamp + "."))
	_, _ = mac.Write(body)
	if !hmac.Equal(signature, mac.Sum(nil)) {
		return ErrInvalidCallback
	}
	return nil
}
