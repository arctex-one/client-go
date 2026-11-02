# Troubleshooting

[Documentation index](../README.md#start-here) · [Setup](getting-started.md) · [Recovery rules](integration.md#durable-submissions-and-uncertain-outcomes)

## Generation and builds

| Symptom | Check and action |
| --- | --- |
| Missing `generated` package or replacement directory | Run the normal build target, which must prepare the pinned client first. Check the relative replacement path from the consumer's `go.mod`. Running bare Go commands does not generate bindings. |
| Published schema checksum/version mismatch | Stop the build. Confirm the schema publication and review the producer change before deliberately updating the client lock/release. Do not bypass verification or check in a schema copy. |
| Schema fetch failed, redirected or returned an error | Check HTTPS reachability and the published schema service. Restore access and rerun the build; no stale local fallback is used. Schema download does not require an API bearer. |
| Destination must be named `arctex-client-go` | Use the documented ignored build directory with that basename, outside the client source directory. |
| Unowned output directory | Choose a dedicated empty output path. Never add an ownership marker to unrelated files to force replacement. |
| Generation lock exists | Wait for the owner to finish and serialize builds sharing that destination. After a crashed build, verify no preparer still uses the output before removing its leftover sibling `.lock` directory. |
| Generated source appears in a diff | Ignore the generated output directory and remove it from tracking. Keep source, generator and schema URL/version/checksum metadata under version control. |

## Requests and credentials

| Symptom | Check and action |
| --- | --- |
| Invalid origin | Supply only the approved HTTP/HTTPS origin, without user information, `/v1`, query or fragment. Enforce HTTPS for production in the integrating application. |
| Invalid POST input | Supply the original nonempty idempotency key and current nonempty bearer. Methods are GET/POST under `/v1/`; the raw boundary rejects unsafe path segments. |
| Authentication/authorization HTTP error | Resolve the current tenant bearer through the approved secret provider and check account/actor permissions. Do not substitute another tenant's credential to recover an original operation. |
| Timeout, cancellation, body-read failure or 503 | Keep the original request/key and pending evidence. Resolve the original result or explicitly replay the same input under the application's recovery policy. These outcomes do not prove non-admission. |
| Redirect response | Correct the approved API origin with the installation operator. Redirect refusal prevents resending payment input or credentials through another route. |
| Response exceeds 4 MiB | Treat `ErrResponseTooLarge`/`Truncated` as incomplete evidence, even for a 2xx status. Retain the original operation and reconcile; do not decode the retained prefix as success. |
| Unknown response state/code | Preserve the raw/typed evidence and pause dependent automated financial decisions. An unknown value cannot be mapped to success, zero or a replacement operation. |

Safe transport error messages omit private request details. Unwrapped causes and response bodies may contain them; avoid putting those values into logs or public tickets.

## Callbacks and recovery

If verification fails, check the exact unmodified body bytes, all three signing headers, the resolved signing key, the receiver clock and its approved replay window. Authenticate before JSON decoding. A key outage must not disable authentication. Retried deliveries can use new timestamps and delivery IDs; deduplicate stable event/effect identities in the inbox.

A failed inbox write must not produce a successful acknowledgement. A failed page handler must not advance its checkpoint. For a 410 or mismatched cursor, follow the producer/application reconciliation procedure instead of silently starting over.

`has_more=false`, a transfer's execution state, an account balance version and a chain transaction hash cannot substitute for committed financial evidence or a common Accounting cutoff. Unknown/incomplete effects require retained evidence and review. Customer Ledger posting remains the gateway's responsibility.

When reporting a failure, include the client release, schema lock version/hash, method, HTTP status/code, timing and safe correlation identifiers. Keep credentials, keys, raw payment bodies and private evidence out of public issues.
