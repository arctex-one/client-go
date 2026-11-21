package publicapi

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/arctex-one/api-protocol-go/webhook"
	wire "github.com/arctex-one/client-go/generated"
)

// CallbackVerificationWindow is the maximum age of a public-key delivery.
const CallbackVerificationWindow = webhook.Ed25519VerificationWindow

// CallbackKeyClient is satisfied by the generated client. Implementations must
// resolve keys from the receiver's trusted Arctex installation.
type CallbackKeyClient interface {
	GetWebhookVerificationKey(context.Context, string, url.Values) (*wire.WebhookVerificationKey, error)
}

// VerifyCallbackWithPublicKey verifies a detached Ed25519 JWS over the original,
// bounded request body. Persist event and effect identities separately to
// deduplicate authenticated retries.
func VerifyCallbackWithPublicKey(ctx context.Context, body []byte, headers http.Header, client CallbackKeyClient, now time.Time) error {
	if client == nil {
		return ErrInvalidCallback
	}
	delivery, err := webhook.ParseEd25519(headers, now)
	if err != nil {
		return err
	}
	kid := delivery.KeyID()
	jwk, err := client.GetWebhookVerificationKey(ctx, kid, nil)
	if err != nil {
		var response *wire.ResponseError
		if errors.As(err, &response) && response.Status == http.StatusNotFound && !response.Truncated {
			return ErrInvalidCallback
		}
		return fmt.Errorf("callback verification key lookup: %w", err)
	}
	if jwk == nil || jwk.Kid != kid || !CallbackKeyUsable(jwk, now) {
		return ErrInvalidCallback
	}
	return delivery.Verify(body, callbackPublicKey(jwk))
}

// ValidCallbackKeyID restricts the untrusted key selector before lookup.
func ValidCallbackKeyID(kid string) bool {
	return webhook.ValidCallbackKeyID(kid)
}

// CallbackKeyUsable validates the public key profile and its rotation cutoff.
// Check it on every use, including when the key came from a cache.
func CallbackKeyUsable(key *wire.WebhookVerificationKey, now time.Time) bool {
	if key == nil || !ValidCallbackKeyID(key.Kid) || key.Kid != CallbackKeyID(callbackPublicKey(key)) {
		return false
	}
	if key.ExpiresAt == nil {
		return true
	}
	expiry, err := time.Parse(time.RFC3339Nano, *key.ExpiresAt)
	return err == nil && now.Before(expiry)
}

// CallbackKeyID returns the RFC 7638 thumbprint of an Ed25519 public JWK.
func CallbackKeyID(key ed25519.PublicKey) string {
	return webhook.CallbackKeyID(key)
}

func callbackPublicKey(key *wire.WebhookVerificationKey) ed25519.PublicKey {
	if key.Kty != "OKP" || key.Crv != "Ed25519" || key.Alg != "EdDSA" || key.Use != "sig" ||
		len(key.X) != base64.RawURLEncoding.EncodedLen(ed25519.PublicKeySize) {
		return nil
	}
	value, err := base64.RawURLEncoding.Strict().DecodeString(key.X)
	if err != nil {
		return nil
	}
	return ed25519.PublicKey(value)
}
