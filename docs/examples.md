# Usage examples

[Documentation index](../README.md#start-here) · [Setup](getting-started.md) · [API reference](api.md)

Each Go block below is a complete source file that can be compiled after [client preparation](getting-started.md#add-the-dependency). The functions are application examples, not additional SDK APIs. They do not run until called. Supply your approved origin, current credential, account/asset mapping and durable storage operations. No example contains credentials or a runnable payment destination.

## Read accounts

Resolve the bearer for the current operation. A supplied HTTP client can share the application's connection pool.

```go
package example

import (
    "context"
    "net/http"

    wire "github.com/arctex-one/client-go/generated"
)

func ReadAccounts(ctx context.Context, origin, currentBearer string, httpClient *http.Client) (*wire.AccountList, error) {
    client := wire.Client{BaseURL: origin, HTTP: httpClient, Bearer: currentBearer}
    return client.ListAccounts(ctx, nil)
}
```

`BaseURL` is an origin such as the HTTPS endpoint issued during onboarding, with no `/v1` suffix. Use `GetBalance(ctx, accountRef, nil)` for an assigned account's balance; it is a processing account, not a customer Ledger balance or physical-wallet inventory.

## Create a receiving address

Authorize the customer's account/asset first. The `persist` callback must retain the key, exact request and original application ownership before the remote call. After success, durably bind the returned address reference to that ownership before exposing it to a customer.

```go
package example

import (
    "context"

    wire "github.com/arctex-one/client-go/generated"
)

func CreateAddress(ctx context.Context, client *wire.Client, originalKey string,
    request wire.CreateReceivingAddress,
    persist func(context.Context, string, wire.CreateReceivingAddress) error,
) (*wire.ReceivingAddress, error) {
    if err := persist(ctx, originalKey, request); err != nil {
        return nil, err
    }
    return client.CreateReceivingAddress(ctx, nil, originalKey, request)
}
```

The request contains `AccountRef` and `Asset`. Retain historical address attribution through rotations and late deposits; a pooled account alone does not identify the original customer.

## Request a quote

This example constructs an ordinary native-asset quote. Network/route support comes from the installation's catalog. For a token, use the exact `TOKEN_CONTRACT` asset identity and contract from that catalog. Do not normalize the original request differently on replay.

```go
package example

import (
    "context"

    wire "github.com/arctex-one/client-go/generated"
)

func RequestNativeQuote(ctx context.Context, client *wire.Client, originalKey,
    accountRef, network, currency, recipient, amount string,
    persist func(context.Context, string, wire.QuoteRequest) error,
) (*wire.Quote, error) {
    asset := wire.Asset{Network: network, Currency: currency, Kind: "NATIVE"}
    request := wire.QuoteRequest{
        SourceAccountRef: accountRef,
        SourceAsset: asset,
        Destination: wire.Destination{Asset: asset, Address: recipient},
        Amount: amount,
        SettlementMode: "FIX_SOURCE",
        FeeTreatment: "ADD_TO_SOURCE",
    }
    if err := persist(ctx, originalKey, request); err != nil {
        return nil, err
    }
    return client.CreateQuote(ctx, nil, originalKey, request)
}
```

Before presenting or accepting a quote, your gateway checks its original instructions, asset/account scope, executable state/readiness and complete economics. Absent price fields are not zero. Apply your customer pricing and approval rules separately from Arctex processing fees.

## Submit an accepted quote

Call this only after the quote has passed the gateway's admission and customer-authorization rules. The optional `ClientReference` is business correlation, not the idempotency key. The persistence callback must converge retries onto the original retained record and reject conflicting input.

```go
package example

import (
    "context"

    wire "github.com/arctex-one/client-go/generated"
)

func SubmitAcceptedQuote(ctx context.Context, client *wire.Client, originalKey string,
    request wire.SubmitTransfer,
    persist func(context.Context, string, wire.SubmitTransfer) error,
) (*wire.Transfer, error) {
    if err := persist(ctx, originalKey, request); err != nil {
        return nil, err
    }
    return client.SubmitTransfer(ctx, nil, originalKey, request)
}
```

Set `request.QuoteRef` to the original accepted reference. If the reply is lost or invalid, retain the submission for recovery. Replay the same key and request under your recovery policy; do not create a replacement quote/payment because the response was uncertain. A returned transfer state is not a customer settlement receipt.

For multiple withdrawals, persist each independent operation, accepted quote and key separately, then submit each through the same method. There is no batch submission method or atomic all-or-nothing guarantee in this client. Concurrency does not change customer reservation, provider admission or per-operation recovery ownership. Apply application-level concurrency bounds, and resolve the current bearer per operation.

## Recover one page of financial effects

The handler must validate each effect's scope and financial interpretation and durably deduplicate/apply or quarantine it under the application's existing rules. Save the returned cursor only after every item is durably handled. If saving fails, re-reading the page must be safe.

```go
package example

import (
    "context"
    "net/url"

    wire "github.com/arctex-one/client-go/generated"
)

func RecoverEffectPage(ctx context.Context, client *wire.Client, checkpoint string,
    handle func(context.Context, wire.FinancialEffect) error,
    saveCheckpoint func(context.Context, string) error,
) (bool, error) {
    query := url.Values{}
    if checkpoint != "" {
        query.Set("cursor", checkpoint)
    }
    page, err := client.ListFinancialEffects(ctx, query)
    if err != nil {
        return false, err
    }
    for _, effect := range page.Items {
        if err := handle(ctx, effect); err != nil {
            return false, err
        }
    }
    if err := saveCheckpoint(ctx, page.NextCursor); err != nil {
        return false, err
    }
    return page.HasMore, nil
}
```

Keep polling from the saved cursor after `has_more=false` or an empty page. That value describes the current projection head, not completeness of source Accounting or a common cutoff with an independently read balance. Do not silently reset an expired/mismatched cursor. Use `publicapi.ReadEffect` when interpreting raw effects, checking its error before any financial decision.

## Authenticate and retain a financial callback

The application supplies its body-size limit, clock, replay window and current/approved rotation key. `retainInbox` must validate stable event identities and tenant scope, store the original authenticated bytes and deduplicate durably. An HTTP success follows successful retention. This receiver performs no customer posting.

```go
package example

import (
    "context"
    "encoding/json"
    "io"
    "net/http"
    "time"

    publicapi "github.com/arctex-one/client-go"
    wire "github.com/arctex-one/client-go/generated"
)

func FinancialCallback(w http.ResponseWriter, r *http.Request, key []byte,
    now time.Time, window time.Duration, bodyLimit int64,
    retainInbox func(context.Context, json.RawMessage, wire.FinancialEvent) error,
) {
    if bodyLimit <= 0 {
        http.Error(w, "receiver unavailable", http.StatusServiceUnavailable)
        return
    }
    raw, err := io.ReadAll(http.MaxBytesReader(w, r.Body, bodyLimit))
    if err != nil {
        http.Error(w, "invalid body", http.StatusBadRequest)
        return
    }
    if err := publicapi.VerifyCallback(raw, r.Header, key, now, window); err != nil {
        http.Error(w, "invalid authentication", http.StatusUnauthorized)
        return
    }
    var event wire.FinancialEvent
    if err := json.Unmarshal(raw, &event); err != nil {
        http.Error(w, "invalid event", http.StatusBadRequest)
        return
    }
    if err := retainInbox(r.Context(), raw, event); err != nil {
        http.Error(w, "event not retained", http.StatusServiceUnavailable)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
```

The inbox worker joins the event/effect to immutable application attribution and committed evidence before a business consequence. A delivery ID or timestamp alone is insufficient for deduplication. Do not log raw callback bodies or signing keys. The application owns HTTP-server timeouts and key-rotation policy.

## Retain exact HTTP response bytes

Use the raw boundary when your gateway needs the original successful response bytes before projection. Resolve the bearer on every call, retain original POST input first and validate the returned JSON/account scope in the gateway.

```go
package example

import (
    "context"
    "encoding/json"
    "net/http"
    "net/url"

    "github.com/arctex-one/client-go/transport"
)

func ReadRawBalance(ctx context.Context, boundary *transport.Client,
    currentBearer, accountRef string,
) (json.RawMessage, error) {
    // Public references must be encoded as individual path segments.
    return boundary.Call(ctx, currentBearer, http.MethodGet,
        "/v1/accounts/"+url.PathEscape(accountRef)+"/balance", "", nil)
}
```

Construct the boundary once with `transport.New(origin, httpClient)`. Its `CloseIdleConnections` method supports application shutdown. Inspect `ResponseError` through `errors.As`; keep its bounded body in approved private evidence, and reject truncated evidence even when the HTTP status is successful. See the [error reference](api.md#errors).
