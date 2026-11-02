.PHONY: generate build test check

generate:
	go run ./tools/generate ./contract/openapi.json ./generated

build: generate
	go build ./...

test: generate
	go test -race ./...

check: generate
	go test -race ./...
	go vet ./...
	@output=$$(mktemp -d); trap 'rm -rf "$$output"' EXIT; \
		go run ./tools/generate ./contract/openapi.json "$$output"; \
		cmp generated/api.gen.go "$$output/api.gen.go"
