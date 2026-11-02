// Package publicapi provides callback authentication and financial interpretation
// guards for Arctex One's public API. Typed requests, responses and HTTP methods
// live in generated; transport supports gateways retaining exact raw evidence.
// Applications own durable idempotency, credential rotation and customer ledgers.
//
//go:generate go run ./tools/generate ./contract/openapi.json ./generated
package publicapi
