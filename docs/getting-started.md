# Getting started

[Documentation index](../README.md#start-here) · [Examples](examples.md) · [API reference](api.md)

Use Go 1.27.1 or later and an application-owned API origin, bearer resolver and account mapping. Building the client needs HTTPS access to the public Go module infrastructure and the schema URL in `schema.lock.json`. Schema generation requires no tenant bearer or custody key.

## Add the dependency

For an application whose `go.mod` is at the repository root, add:

```go.mod
require github.com/arctex-one/client-go v0.3.0

replace github.com/arctex-one/client-go => ./.generated/arctex-client-go
```

Add `/.generated/` to the application's `.gitignore`, then prepare the dependency before resolving packages:

```sh
go run github.com/arctex-one/client-go/tools/prepare@v0.3.0 .generated/arctex-client-go
go mod tidy
```

Keep the required release and the preparation command on the same version. The replacement is a generated build artifact obtained from the public release. It must not point at a private processing checkout or a copied, tracked SDK.

For nested modules, adjust the replacement relative to that module's `go.mod`. For example, a module at `api/gateway/crypto` uses `../../../.generated/arctex-client-go`. Run preparation from the repository root with the root build-directory path.

## Generate on every build

A root application can use this Makefile:

```makefile
.PHONY: prepare build test vet

prepare:
	go run github.com/arctex-one/client-go/tools/prepare@v0.3.0 .generated/arctex-client-go

build: prepare
	go build ./...

test: prepare
	go test ./...

vet: prepare
	go vet ./...
```

Use the corresponding preparation step in CI, lint commands, development reloaders and container builds. In a container, copy the consumer module files first, run preparation at the repository root, and compile with the same relative directory layout used by `go.mod`. Do not copy a developer's generated directory into an image.

Go does not execute generators automatically during `go build`, `go test`, `go vet`, `go mod tidy` or `go get`. The application's build entry points must run preparation first and stop immediately if it fails. A pre-existing generated directory is not a fallback after a failed preparation.

## What preparation does

1. Go resolves the explicitly versioned public preparation tool with normal module verification.
2. The tool copies that release's source into a temporary build directory without modifying the module cache.
3. Its generator fetches the external schema and verifies the locked SHA-256 and contract version.
4. It generates the typed client and models, then publishes the completed build directory.

The destination must be named `arctex-client-go` and be outside the client source directory. An existing destination must carry the tool's ownership marker. Unowned directories and symlinks are rejected. Preparation uses an exclusive directory lock; serialize builds that prepare the same output path. Generation failure preserves the previous owned directory, but the consuming build must still stop.

## Make a first read

The [account example](examples.md#read-accounts) is a complete Go function. Supply the approved origin and a bearer resolved for the current operation. For production, use HTTPS with the installation's trusted TLS transport. The client permits HTTP for private development; applications enforce their own origin policy.

A typed client has `BaseURL`, `HTTP` and `Bearer` fields. The bearer is an in-memory value, not a secret-store reference. Resolve it again for each operation to support rotation. Reuse an application-owned HTTP connection pool, and do not mutate a typed client or shared transport while calls are running.

All calls accept a context. A zero HTTP timeout becomes 30 seconds; a context deadline can shorten it. Redirects are refused even if the supplied HTTP client permits them.

## Verify a client checkout

```sh
make build check
test -z "$(git ls-files generated contract)"
```

The checks compile the generated client, run race-enabled tests and vet, and compare two generated outputs for determinism. They do not send payments. Schema updates require review of the external publication and a deliberate lock/version update; see [Contributing](../CONTRIBUTING.md).
