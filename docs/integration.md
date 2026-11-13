# Integrating Arctex.One

[Documentation index](../README.md#start-here) · [Setup](getting-started.md) · [API reference](api.md) · [Examples](examples.md)

## Ownership boundary

Arctex owns crypto custody, execution and its committed processing effects. The client owns wire serialization, bounded HTTP calls and portable authentication/interpretation guards. Each integrating gateway owns its original installation/tenant/credential binding, customer authorization, account/address attribution, durable submission journal, inbox, business pricing and customer ledger.

Neither the client nor an execution status authorizes a customer credit, debit, fee or fiat-payment completion. Gateway code must join committed evidence to its immutable original business attribution. Keep internal operating/reserved/pending movements, provider fees and your own commercial fees separate. Do not infer financial origin from a transaction hash, amount or description.

## Credentials and transport

Use a secret-provider lookup on each operation. Applications decide how to authenticate to Vault or another approved provider; this library contains no secret-store integration, environment fallback, custody key handling or token caching. A long-lived typed `Client` with a fixed `Bearer` will not rotate itself. Create a client from the current bearer for each operation, or call `transport.Client.Call` with the current bearer directly.

`transport.New(origin, httpClient)` validates an origin with no user information, path prefix, query or fragment. `Call(ctx, bearer, method, path, key, rawBody)` accepts only GET and POST `/v1/` operations. It preserves raw body/response bytes and rejects a POST without an idempotency key. It returns successful bytes without validating their business meaning; validate JSON and resource scope before storing or acting on them. Do not use arbitrary customer-supplied origins, paths or credentials.

Redirects are refused for both raw and typed calls. TLS trust, proxies and connection behavior belong to the supplied HTTP transport. Use an approved HTTPS origin in production. HTTP support exists for development. The library copies the HTTP client, not its underlying transport; callers must keep shared transport settings stable. A default 30-second timeout bounds otherwise unbounded clients. Context cancellation propagates through wrapped errors. Raw clients expose `CloseIdleConnections` for gateway shutdown; it releases idle connections without interrupting active calls.

## Durable submissions and uncertain outcomes

Before any POST, atomically retain the original key, exact validated request, tenant/credential reference and application correlation. Preserve decimal spelling, absent optional fields and accepted catalog resolution. On a known persisted response, return that original response rather than constructing a new operation. On transport or response uncertainty, keep the submission pending for recovery under the same owner.

A quotation is executable only when the producer supplies a supported route, executable readiness and complete accepted economics. Absent economics are not zero. Submit the original quote reference. Quote expiry does not establish that a previously submitted payment was never admitted. Unknown states, codes or origins require retained evidence and review, not a success/default mapping.

The generated methods return an allocated response alongside an error. Never consume that response when `err != nil`. Wire models preserve unknown JSON fields in `AdditionalFields`, decimal strings, version strings and optional field presence. They are bindings, not a JSON Schema validator or posting authorization.

## Pagination and recovery

Use `url.Values` for the surface's documented filters and opaque cursor. For example:

```go
query := url.Values{}
if checkpoint != "" { query.Set("cursor", checkpoint) }
page, err := client.ListFinancialEffects(ctx, query)
if err != nil { return err }
// Durably handle every item, deduplicating by immutable effect identity.
// Only then save page.NextCursor, even if the page was empty.
```

`has_more=false` means caught up to that projection's current head. Continue polling from its returned cursor; it does not close the stream or prove every source Accounting commit has been projected. Keep filters and credential scope unchanged. A malformed/mismatched cursor is an error; 410 requires explicit reconciliation. A deposit can reappear at a later version; compare deposit reference and version. The same financial effect reference must retain identical financial content.

Balance version, financial-effect cursor and chain event time are different orderings. Their numeric values do not prove a common cutoff. Financial migration/reconciliation requires the producer's Accounting-backed snapshot and exact posting membership. This library provides no canonical migration importer, financial-opening writer or custody activation route.

## Callbacks

For an Ed25519 subscription, read the exact raw HTTP body with your receiver's
size limit. Use a generated client configured with the trusted Arctex origin
and tenant bearer credential.

```go
err := publicapi.VerifyCallbackWithPublicKey(request.Context(), rawBody, request.Header, client, now)
```

The verifier requires the Arctex detached JWS profile, fetches the public JWK by
its protected `kid`, and checks its RFC 7638 thumbprint, retirement expiry, the
five-minute signing window and the original bytes. A lookup failure or an
invalid signature prevents acknowledgement. See the
[Arctex webhook guide](https://docs.arctex.one/guide.md#webhooks) for wire details.

For a migration, deploy public-key support before changing the subscription.
Persisted legacy deliveries continue to use their original HMAC key and
`VerifyCallback(rawBody, headers, key, now, window)`. Keep only the verification
modes allowed by the installation's migration policy, and never retry a failed
signature with another algorithm.

After authentication, decode `generated.FinancialEvent`, `DepositEvent` or `TransferEvent`. Persist the authenticated raw envelope and stable event/effect identities in a durable inbox before acknowledgement. Retries can use a new signing time and delivery ID, so timestamp windows do not replace durable event/effect deduplication. Keep key outages fail-closed. Do not log keys, raw callback bodies or sensitive payment details.

## Committed effects

`publicapi.ReadEffect(raw)` decodes a financial effect and checks known asset/line/origin representations, required identifiers, decimal spelling and commit timestamp. On an unknown financial value it returns the decoded record **and** `ErrUnrecognizedFinancialValue`, retaining evidence for quarantine. Other malformed/incomplete evidence also returns an error. Check the error before any dependent decision.

This guard does not prove account ownership, complete source history, posting balance or customer entitlement. It does not validate every field against JSON Schema. Those checks remain at producer and gateway boundaries. A committed adjustment/carry-in effect must not become a second customer deposit. Compensation requires its original committed linkage; signs and ordering are insufficient.

## Source-only builds and migration

Require `github.com/arctex-one/client-go v0.3.0`. Replace the module with a repository-ignored `.generated/arctex-client-go` build directory (adjust the relative path for nested modules). Run `go run github.com/arctex-one/client-go/tools/prepare@v0.3.0 <build-directory>` before every application build, test, analysis and container compile. The explicit pinned version is required; do not use `@latest`. Put generation in the actual build entry points, including CI and Docker, not in a manual prerequisite. `go build` itself does not execute Go generators.

The prepare command obtains its source from the verified public Go module, copies it to a temporary directory and runs the same source generator. It publishes only after generation succeeds, refuses an unowned output directory, preserves the old directory on failure and never mutates the shared Go module cache. Generated Go is ignored, never committed. The marker file identifies owned generated output; do not create one in a user directory. Parallel preparation of the same directory fails explicitly under its generation lock.

In the client checkout, use `make build`, `make test` or `make check`; each always regenerates first. `go generate ./...` remains available for development. The schema URL/version/checksum lock, generator, handwritten transport/guards, tests and documentation are tracked. Neither schema snapshots nor generated Go files are committed.

Replace old `github.com/lazysteff/crypto-processing/publicapi` or `github.com/arctex-one/crypto-processing/publicapi` imports with `github.com/arctex-one/client-go`, keeping `/generated` for wire imports. Remove the copied client source. Keep application-specific handlers, evidence repositories and ledger rules. The generated API/model names and draft 0.5.0 wire shape remain compatible. Both typed and raw clients use one transport. Callers supply a nonempty bearer and original POST key; zero timeouts become 30 seconds. Redirect refusal and explicit oversized-response errors are retained. Sendico's wrapper translates client errors into its existing project error model and preserves its legacy error-evidence shape.
