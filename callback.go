package publicapi

import (
	"net/http"
	"time"

	"github.com/arctex-one/api-protocol-go/webhook"
)

// ErrInvalidCallback covers authentication failures without exposing key or body.
var ErrInvalidCallback = webhook.ErrInvalidCallback

// VerifyCallback authenticates legacy HMAC deliveries over the original, bounded HTTP body before decoding.
// The receiving application supplies its clock/replay window and Vault-resolved
// key. Persist/deduplicate event and effect identities separately: a valid retry
// has a new signing time and delivery identity. This does not authorize posting.
func VerifyCallback(body []byte, headers http.Header, key []byte, now time.Time, window time.Duration) error {
	return webhook.VerifyHMAC(body, headers, key, now, window)
}
