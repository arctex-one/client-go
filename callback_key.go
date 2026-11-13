package publicapi

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	wire "github.com/arctex-one/client-go/generated"
)

// CallbackVerificationWindow is the maximum age of a public-key delivery.
const CallbackVerificationWindow = 5 * time.Minute

// CallbackKeyClient is satisfied by the generated client. Implementations must
// resolve keys from the receiver's trusted Arctex installation.
type CallbackKeyClient interface {
	GetWebhookVerificationKey(context.Context, string, url.Values) (*wire.WebhookVerificationKey, error)
}

// VerifyCallbackWithPublicKey verifies a detached Ed25519 JWS over the original,
// bounded request body. Persist event and effect identities separately to
// deduplicate authenticated retries.
func VerifyCallbackWithPublicKey(ctx context.Context, body []byte, headers http.Header, client CallbackKeyClient, now time.Time) error {
	if client == nil || now.IsZero() || callbackHeader(headers, "X-Callback-Algorithm") != "jws+eddsa-ed25519" {
		return ErrInvalidCallback
	}
	compact := callbackHeader(headers, "X-Callback-JWS")
	protected, detached, ok := strings.Cut(compact, ".")
	payload, encodedSignature, separated := strings.Cut(detached, ".")
	if !ok || !separated || payload != "" || len(protected) > 1024 || len(encodedSignature) != base64.RawURLEncoding.EncodedLen(ed25519.SignatureSize) {
		return ErrInvalidCallback
	}
	encoding := base64.RawURLEncoding.Strict()
	raw, err := encoding.DecodeString(protected)
	if err != nil {
		return ErrInvalidCallback
	}
	alg, kid, timestamp, err := callbackProtectedFields(raw)
	if err != nil || alg != "EdDSA" || !ValidCallbackKeyID(kid) ||
		kid != callbackHeader(headers, "X-Callback-Key-Id") || timestamp != callbackHeader(headers, "X-Callback-Timestamp") {
		return ErrInvalidCallback
	}
	signedAt, err := time.Parse(time.RFC3339Nano, timestamp)
	if err != nil || signedAt.Before(now.Add(-CallbackVerificationWindow)) || signedAt.After(now) {
		return ErrInvalidCallback
	}
	signature, err := encoding.DecodeString(encodedSignature)
	if err != nil {
		return ErrInvalidCallback
	}
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
	key := callbackPublicKey(jwk)
	if key == nil || !ed25519.Verify(key, append([]byte(protected+"."), body...), signature) {
		return ErrInvalidCallback
	}
	return nil
}

// ValidCallbackKeyID restricts the untrusted key selector before lookup.
func ValidCallbackKeyID(kid string) bool {
	if len(kid) != base64.RawURLEncoding.EncodedLen(sha256.Size) {
		return false
	}
	_, err := base64.RawURLEncoding.Strict().DecodeString(kid)
	return err == nil
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
	if len(key) != ed25519.PublicKeySize {
		return ""
	}
	x := base64.RawURLEncoding.EncodeToString(key)
	digest := sha256.Sum256([]byte(`{"crv":"Ed25519","kty":"OKP","x":"` + x + `"}`))
	return base64.RawURLEncoding.EncodeToString(digest[:])
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

// Reject duplicate and unknown protected fields before resolving a key. This
// profile accepts only the protected parameters emitted by the Arctex signer.
func callbackProtectedFields(raw []byte) (alg, kid, timestamp string, err error) {
	decoder := json.NewDecoder(bytes.NewReader(raw))
	if token, err := decoder.Token(); err != nil || token != json.Delim('{') {
		return "", "", "", ErrInvalidCallback
	}
	fields := make(map[string]json.RawMessage, 5)
	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			return "", "", "", ErrInvalidCallback
		}
		name, ok := token.(string)
		if !ok || fields[name] != nil {
			return "", "", "", ErrInvalidCallback
		}
		var value json.RawMessage
		if err := decoder.Decode(&value); err != nil {
			return "", "", "", ErrInvalidCallback
		}
		fields[name] = value
	}
	if token, err := decoder.Token(); err != nil || token != json.Delim('}') || len(bytes.TrimSpace(raw[decoder.InputOffset():])) != 0 {
		return "", "", "", ErrInvalidCallback
	}
	var crit []string
	if len(fields) != 5 || string(fields["b64"]) != "false" ||
		json.Unmarshal(fields["crit"], &crit) != nil || len(crit) != 1 || crit[0] != "b64" ||
		json.Unmarshal(fields["alg"], &alg) != nil || json.Unmarshal(fields["kid"], &kid) != nil ||
		json.Unmarshal(fields["ts"], &timestamp) != nil {
		return "", "", "", ErrInvalidCallback
	}
	return alg, kid, timestamp, nil
}
