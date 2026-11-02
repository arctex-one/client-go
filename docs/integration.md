# Integrating Arctex One

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

Read the exact raw HTTP body with your receiver's size limit. Resolve the current or approved rotation key. Authenticate **before decoding**:

```go
err := publicapi.VerifyCallback(rawBody, request.Header, signingKey, now, replayWindow)
if err != nil { /* reject; do not acknowledge or post */ }
```

The signature is HMAC-SHA256 over the exact `X-Callback-Timestamp`, a period, and the original body. The algorithm header must be `hmac-sha256`. Missing/ambiguous headers, wrong signatures, changed bytes and timestamps outside the supplied window fail closed. Your application supplies the clock/window policy; the client invents no SLA.

After authentication, decode `generated.FinancialEvent`, `DepositEvent` or `TransferEvent`. Persist the authenticated raw envelope and stable event/effect identities in a durable inbox before acknowledgement. Retries can use a new signing time and delivery ID, so timestamp windows do not replace durable event/effect deduplication. Keep key outages fail-closed. Do not log keys, raw callback bodies or sensitive payment details.

## Committed effects

`publicapi.ReadEffect(raw)` decodes a financial effect and checks known asset/line/origin representations, required identifiers, decimal spelling and commit timestamp. On an unknown financial value it returns the decoded record **and** `ErrUnrecognizedFinancialValue`, retaining evidence for quarantine. Other malformed/incomplete evidence also returns an error. Check the error before any dependent decision.

This guard does not prove account ownership, complete source history, posting balance or customer entitlement. It does not validate every field against JSON Schema. Those checks remain at producer and gateway boundaries. A committed adjustment/carry-in effect must not become a second customer deposit. Compensation requires its original committed linkage; signs and ordering are insufficient.

## Migration from bundled preview clients

Replace old `github.com/lazysteff/crypto-processing/publicapi` or `github.com/arctex-one/crypto-processing/publicapi` imports with `github.com/arctex-one/client-go`, keeping `/generated` for wire imports. Require the published version in `go.mod`, remove the old local replacement and remove the duplicate client source. Keep application-specific handlers, evidence repositories and ledger rules.

The generated API/model names and draft 0.5.0 wire shape remain compatible. Both typed and raw clients now use one transport. Callers must supply a nonempty bearer and original POST key; zero timeouts become 30 seconds. Redirect refusal and explicit oversized-response errors are retained from the reviewed organization preview. Sendico's wrapper translates client errors into its existing project error model and preserves its legacy error-evidence shape.
