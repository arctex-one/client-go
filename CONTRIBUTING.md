# Contributing

Keep this module portable and independent of processing implementation, databases,
Vault bootstrap and application ledgers. The public contract is the authority for
wire types. Do not add duplicate financial models or gateway business logic.

Review changes against both Sendico and MCRM consumers. Preserve exact decimal
strings, optional-field presence, unknown response evidence and original
idempotency identities. Never add automatic payment replacement or inferred
financial commitment. Keep credential values and real customer data out of
fixtures, documentation, logs and issues.

Run `go generate ./...`, `go test -race ./...` and `go vet ./...`; inspect the full
diff and verify generation leaves no unexpected changes. API/error/transport
changes require behavior tests and a migration note. Update provenance when the
producer contract changes. Version SDK releases separately from contract versions.
