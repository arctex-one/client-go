# Arctex One Go client

A shared Go client for [Arctex One](https://github.com/arctex-one) digital asset processing. Sendico and MCRM use this module; application-specific gateways retain customer authorization, durable evidence, pricing and ledger posting.

```sh
go get github.com/arctex-one/client-go@v0.1.1
```

Requires Go 1.27.1 or later. The module uses only the Go standard library.

This first release binds the **draft public API contract 0.5.0**, served under `/v1`. The SDK release version, contract document version and HTTP API version are independent. This release does not establish production availability, retention guarantees or a financial migration approval. Use the API origin and tenant account references approved for your installation.

## Packages

| Package | Purpose |
| --- | --- |
| `generated` | Typed HTTP operations and wire models, generated from the bundled OpenAPI contract |
| `publicapi` (module root) | Exact-body callback authentication and committed financial-effect interpretation guards |
| `transport` | Shared HTTP boundary for consumers that retain exact raw request/response evidence |

The [OpenAPI contract](contract/openapi.json) describes request fields, states, errors, ownership and pagination. [Integration guidance](docs/integration.md) covers durable submissions, callbacks, recovery and financial boundaries. [Provenance](PROVENANCE.md) records the extraction source.

## Read an account

```go
import (
    "context"
    "net/http"
    "time"

    wire "github.com/arctex-one/client-go/generated"
)

func readAccounts(ctx context.Context, origin, currentBearer string) (*wire.AccountList, error) {
    client := wire.Client{
        BaseURL: origin,
        Bearer: currentBearer,
        HTTP: &http.Client{Timeout: 15 * time.Second},
    }
    return client.ListAccounts(ctx, nil)
}
```

Resolve the current bearer through your application's approved secret provider before each operation. `Bearer` is an in-memory credential, not a Vault path. Do not persist or log it. Use HTTPS and the installation's trusted TLS transport in production; HTTP is supported for private development stands. Redirects are refused, including when the supplied HTTP client permits them. The supplied client is copied. An omitted client or zero timeout uses a 30-second deadline; contexts can impose shorter deadlines. Do not mutate a typed client's fields while it is in use.

## Create and submit a quote

```go
request := wire.QuoteRequest{
    SourceAccountRef: accountRef,
    SourceAsset: wire.Asset{Network: network, Currency: "TRX", Kind: "NATIVE"},
    Destination: wire.Destination{
        Asset: wire.Asset{Network: network, Currency: "TRX", Kind: "NATIVE"},
        Address: recipient,
    },
    Amount: "5.000000",
    SettlementMode: "FIX_SOURCE",
    FeeTreatment: "ADD_TO_SOURCE",
}
// Persist quoteKey, request and business correlation BEFORE this call.
quote, err := client.CreateQuote(ctx, nil, quoteKey, request)
if err != nil { return err }
// Validate executable state, original asset/instructions and complete economics
// using your gateway's admission policy before offering or accepting this quote.
// Persist transferKey and this exact accepted input BEFORE submission.
transfer, err := client.SubmitTransfer(ctx, nil, transferKey, wire.SubmitTransfer{
    QuoteRef: quote.QuoteRef,
})
```

The snippet belongs inside an application function; `client`, `ctx`, account/network/recipient and durable keys come from that application. Network/asset/route support is installation-specific. The SDK does not generate payment identities, check customer ownership, reserve money or persist a submission. `client_reference` is optional business correlation, not an idempotency key. One accepted quote must not be used to create a second payment.

## Public operations

All methods accept `context.Context` and `url.Values` query parameters. POST methods also require the original idempotency key and a typed body. Reference arguments are path-escaped.

| Surface | Methods |
| --- | --- |
| Accounts | `ListAccounts`, `GetBalance`, `ListStatementEntries`, `GetStatementEntry` |
| Catalog | `ListAssets`, `ListBridgeRoutes` |
| Receiving addresses | `CreateReceivingAddress`, `ListReceivingAddresses`, `GetReceivingAddress` |
| Quotes and transfers | `CreateQuote`, `SubmitTransfer`, `ListTransfers`, `GetTransfer` |
| Recovery | `ListDepositRevisions`, `GetDeposit`, `ListFinancialEffects`, `GetFinancialEffect` |

## Errors and retry safety

```go
var response *wire.ResponseError
if errors.As(err, &response) {
    // Status and Code are available without rendering the remote body.
    // Body is private bounded evidence. Persist it only in approved storage.
    if response.Truncated { /* quarantine incomplete evidence */ }
}
if errors.Is(err, wire.ErrResponseTooLarge) { /* reconcile original operation */ }
```

HTTP responses are limited to 4 MiB. Oversized responses, including successful statuses, return an explicit truncated error and are never decoded as success. Non-2xx responses retain their status and body; unknown error codes remain unknown. Public transport error messages omit credentials, URLs and response bodies. Causes remain available through `errors.Is`/`errors.As`; avoid logging underlying private request details.

There are no application-level retries. Go's HTTP transport may retry an eligible request on a reused connection; it preserves that request's identity. After a timeout, lost reply, cancellation, 503 or unknown outcome, retain the original operation, request and key. Explicitly recover or replay that same submission under your existing policy. Do not replace the key, regenerate a quote or infer failure from an execution status. A 410 recovery checkpoint requires reconciliation, not silently restarting a stream.

## Development

```sh
go generate ./...
go test -race ./...
go vet ./...
git diff --exit-code -- generated/api.gen.go
```

Generation is deterministic and requires no external generator. Update `contract/openapi.json` only from a reviewed producer contract revision; regenerate and run the consumer contract tests together. CI verifies regeneration, race tests and vet on GitHub-hosted runners for this public repository. See [CONTRIBUTING.md](CONTRIBUTING.md).
