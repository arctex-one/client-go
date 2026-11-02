# Arctex.One Go client

The shared Go client for Arctex.One processing: accounts, receiving addresses, quotations, transfers and committed financial effects over the public HTTP API. Sendico and MCRM integration gateways use this module while retaining their customer authorization, durable recovery and accounting rules.

**Module:** `github.com/arctex-one/client-go` · **Client release:** `v0.3.0` · **API:** `/v1`, draft contract `0.5.0` · **Go:** `1.27.1+`

## Start here

| Task | Documentation |
| --- | --- |
| Add the client to an application and automate generation | [Getting started](docs/getting-started.md) |
| Find a method, request type or error | [API reference](docs/api.md) |
| Read accounts, create an address, quote, submit, recover and receive callbacks | [Usage examples](docs/examples.md) |
| Design durable submissions, attribution, callbacks and financial recovery | [Integration guide](docs/integration.md) |
| Resolve build, schema, credential or recovery failures | [Troubleshooting](docs/troubleshooting.md) |
| Change the client or its schema pin | [Contributing](CONTRIBUTING.md) · [Provenance](PROVENANCE.md) |

## Build the client

```sh
git clone https://github.com/arctex-one/client-go.git
cd client-go
make build
```

`make build`, `make test` and `make check` regenerate wire code before compilation. The generator fetches the [published OpenAPI schema](https://docs.arctex.one/openapi.json), then verifies the version and SHA-256 recorded in [schema.lock.json](schema.lock.json). Generated Go and schema snapshots are not tracked in the current source tree. The library uses only the Go standard library.

Consumer applications must also generate before builds, tests, lint and container compilation:

```sh
go run github.com/arctex-one/client-go/tools/prepare@v0.3.0 .generated/arctex-client-go
```

The command prepares the pinned public module in an ignored, owned build directory. The application's `go.mod` replaces the dependency with that generated directory. Follow the [complete setup recipe](docs/getting-started.md); bare `go build` and `go get` do not run generators.

## Packages

| Import | Purpose |
| --- | --- |
| `github.com/arctex-one/client-go/generated` | Typed HTTP methods and wire models, generated during the build |
| `github.com/arctex-one/client-go` (package `publicapi`) | Callback authentication and financial-effect interpretation guards |
| `github.com/arctex-one/client-go/transport` | Bounded HTTP transport for gateways retaining exact request and response bytes |

Use the API origin, tenant credentials, account references and supported assets issued for your installation. The client makes no application-level retries and stores no credentials or payment journal. Persist the original request and idempotency key before a POST; an uncertain response must retain that identity. Execution state alone cannot authorize customer accounting.

SDK versions, schema document versions and HTTP API versions are separate. The current binding targets the draft contract; a client release does not establish production API availability or a retention SLA. HTTP is the supported integration protocol. See the [integration guide](docs/integration.md) for ownership and compatibility rules.
