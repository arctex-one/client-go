# API reference

[Documentation index](../README.md#start-here) · [Examples](examples.md) · [Published schema](https://docs.arctex.one/openapi.json)

This reference describes client release `v0.3.0`. Resource fields, supported filters, status values and installation-specific route constraints are defined by the external schema pinned in that release. No schema snapshot is maintained here.

## Typed HTTP client

Import `wire "github.com/arctex-one/client-go/generated"` after build-time preparation. Construct `wire.Client{BaseURL: origin, HTTP: httpClient, Bearer: currentBearer}`.

Every method's first argument is `context.Context`. Reference arguments are strings, in the order shown in the HTTP path. All methods then accept `url.Values` for query parameters; use `nil` when no filters are needed. POST methods additionally take the original idempotency key and typed request body. They return a pointer to the response type and an error. Ignore the response whenever the error is non-nil.

| Method | HTTP path | POST request | Response |
| --- | --- | --- | --- |
| `ListAccounts` | `GET /v1/accounts` | — | `AccountList` |
| `GetBalance` | `GET /v1/accounts/{account_ref}/balance` | — | `Balance` |
| `ListStatementEntries` | `GET /v1/accounts/{account_ref}/statement-entries` | — | `StatementList` |
| `GetStatementEntry` | `GET /v1/accounts/{account_ref}/statement-entries/{entry_ref}` | — | `StatementEntry` |
| `ListAssets` | `GET /v1/assets` | — | `AssetList` |
| `ListBridgeRoutes` | `GET /v1/bridge-routes` | — | `BridgeRouteList` |
| `CreateReceivingAddress` | `POST /v1/receiving-addresses` | `CreateReceivingAddress` | `ReceivingAddress` |
| `ListReceivingAddresses` | `GET /v1/receiving-addresses` | — | `ReceivingAddressList` |
| `GetReceivingAddress` | `GET /v1/receiving-addresses/{receiving_address_ref}` | — | `ReceivingAddress` |
| `CreateQuote` | `POST /v1/quotes` | `QuoteRequest` | `Quote` |
| `SubmitTransfer` | `POST /v1/transfers` | `SubmitTransfer` | `Transfer` |
| `ListTransfers` | `GET /v1/transfers` | — | `TransferList` |
| `GetTransfer` | `GET /v1/transfers/{transfer_ref}` | — | `Transfer` |
| `ListDepositRevisions` | `GET /v1/deposits` | — | `DepositList` |
| `GetDeposit` | `GET /v1/deposits/{deposit_ref}` | — | `Deposit` |
| `ListFinancialEffects` | `GET /v1/financial-effects` | — | `FinancialEffectList` |
| `GetFinancialEffect` | `GET /v1/financial-effects/{effect_ref}` | — | `FinancialEffect` |

For example, `GetBalance(ctx, accountRef, nil)` reads an account, and `SubmitTransfer(ctx, nil, originalKey, request)` submits an accepted quote. The client path-escapes references and URL-encodes query values. It does not authorize customer access to those references.

## Query parameters

Pass filters through `url.Values`; its values are strings even for numeric parameters. Use only parameters published for the selected surface.

| Method | Optional query parameters |
| --- | --- |
| `GetBalance` | `as_of` timestamp |
| `ListStatementEntries` | `limit`, `offset` |
| `ListReceivingAddresses` | `account_ref`, `cursor`, `limit` |
| `ListDepositRevisions` | `account_ref`, `receiving_address_ref`, `cursor`, `limit` |
| `ListTransfers` | `quote_ref`, `transfer_ref`, `state`, `cursor`, `limit` |
| `ListFinancialEffects` | `account_ref`, `deposit_ref`, `transfer_ref`, `journal_entry_ref`, `cursor`, `limit` |

In the pinned contract, `limit` ranges from 1 to 100 and defaults to 25; statement `offset` is nonnegative and defaults to 0. Other methods have no published query parameters, so pass `nil`. Keep a stream's filters and credential scope unchanged across its cursor. The durable checkpoint rules for deposit revisions and financial effects are described in the [integration guide](integration.md#pagination-and-recovery); a cursor is not an Accounting snapshot identifier.

## Wire values

Amounts and versions are strings. Preserve decimal spelling, complete network/asset/contract identity and optional field presence across a replay. Never convert versions or amounts through floating point. Optional economics use pointers; absence does not mean zero or an executable quote.

Models retain unknown response properties in `AdditionalFields`. Unknown states, financial classifications and origins require explicit handling; decoding a value does not make it safe to act on. These bindings do not perform complete JSON Schema or business validation.

## Raw transport

Import `github.com/arctex-one/client-go/transport`.

| API | Contract |
| --- | --- |
| `New(origin string, client http.Client) (*Client, error)` | Origin has no user information, path prefix, query or fragment; supports HTTP/HTTPS, refuses redirects, supplies a 30-second zero-timeout default |
| `(*Client).Call(ctx, bearer, method, path, key, body)` | One application-level GET/POST attempt under `/v1/`; POST requires a nonempty original key; body and successful response are `json.RawMessage` |
| `(*Client).Origin() string` | The retained origin, with a trailing slash removed |
| `(*Client).CloseIdleConnections()` | Releases idle pooled connections without interrupting active calls or preventing later calls |

`Call` takes the bearer per operation. It preserves successful bytes without re-encoding or validating JSON. The caller validates the response and its ownership before retaining or using it. The raw client is safe for concurrent calls after construction; callers must keep shared transport settings stable.

## Errors

| Error | Meaning and caller action |
| --- | --- |
| `transport.ErrInvalidArgument` | Invalid origin, method, path, bearer or POST key; fix input without inventing a replacement financial operation |
| `transport.ErrUnavailable` / `*transport.RequestError` | Transport or body-read failure; preserve the original submission and recover its outcome; wrapped causes support `errors.Is`/`errors.As` |
| `*wire.ResponseError` / `*transport.ResponseError` | Same error type: non-2xx HTTP response or oversized response, with `Status`, bounded private `Body`, `Code` and `Truncated` |
| `wire.ErrResponseTooLarge` / `transport.ErrResponseTooLarge` | Same sentinel: response exceeded 4 MiB, including a successful HTTP status; never parse truncated evidence as success |
| JSON decoding error | The successful response could not be decoded; retain the original operation for compatibility/recovery review |
| `publicapi.ErrInvalidCallback` | Callback authentication failed; do not decode, acknowledge or post from it |
| `publicapi.ErrUnrecognizedFinancialValue` | Financial interpretation encountered an unknown value; retain/quarantine the returned decoded effect and error |

Transport errors have safe public messages, but unwrapped causes and response bodies can contain private request/evidence details. Store them only in approved private recovery records. The client performs no application-level retry. Standard Go HTTP transport behavior may retry an eligible original request on a reused connection.

An HTTP status does not by itself prove whether a payment was admitted. In particular, a timeout or 503 cannot authorize a new payment identity. A 410 recovery checkpoint requires explicit reconciliation. See [Recovery](integration.md#pagination-and-recovery).

## Callback authentication

The root package is named `publicapi`:

```text
VerifyCallback(body []byte, headers http.Header, key []byte,
    now time.Time, window time.Duration) error
```

It authenticates HMAC-SHA256 over the exact timestamp header, a period, and the original body. Required headers are `X-Callback-Timestamp`, `X-Callback-Algorithm: hmac-sha256`, and the hex `X-Callback-Signature`. The supplied replay window must be positive, and the clock and key must be available. Authenticate before decoding; persist a validated inbox before acknowledgement. Delivery retries can have a different timestamp and delivery ID, so durable event/effect deduplication remains necessary.

## Financial interpretation

```text
ReadEffect(raw []byte) (*generated.FinancialEffect, error)
```

`publicapi.ReadEffect` checks required identifiers, known line/asset/origin representations, decimal spelling and commit time. It can return a decoded record alongside an error; that record is evidence for quarantine, not permission to post. The function does not prove account ownership, complete source history or customer entitlement. See [Committed effects](integration.md#committed-effects).
