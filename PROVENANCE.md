# Extraction provenance

This module was extracted from the public integration kit consumed by Sendico's
Arctex gateway and MCRM's crypto gateway. Business ledgers, immutable application
evidence, customer attribution, credentials and custody implementations are not
part of the extraction.

The reviewed base kit was generated from `arctex-one/crypto-processing` commit
`61f4ca42898fcb0e05520b199c1d68f5b0918ad2`, draft contract 0.5.0. Its original export
manifest SHA-256 was
`d12e55ab86f7a5edc06dadb51c0602294ad2933ddbbcfd075017701e619c2614`.
All export file checksums and tracked source hashes were verified, and generated
source/export were independently reproduced before extraction.

The OpenAPI bytes and wire-model generator semantics are preserved. The public
module path changes to `github.com/arctex-one/client-go`. Generic HTTP transport
was extracted from Sendico's `api/gateway/arctex/adapter/http.go` at Sendico main
`3ca3cad21d7677d659da99de225dbe18e1dd5e78`, combined with the reviewed preview's
redirect refusal and explicit truncated-response error behavior, and shared by
the typed client and raw gateway boundary. Standalone validation, safe error
messages and regression tests accompany that extraction.

Future producer contract updates must be reviewed and regenerated here. Consumers
pin a public source-only SDK release. Every build prepares that verified source
in an ignored directory and regenerates the wire API; generated Go is not tracked.
No private processing checkout is required. This provenance identifies source code only; it
does not attest production deployment or financial/custody migration readiness.
