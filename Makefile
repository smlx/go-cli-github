# local development targets

.PHONY: test
test: mod-tidy generate
	go test -v ./...

.PHONY: mod-tidy
mod-tidy:
	go mod tidy

.PHONY: generate
generate: mod-tidy
	go generate ./...

.PHONY: build
build:
	goreleaser build --clean --single-target --snapshot

.PHONY: lint
lint:
	golangci-lint run

.PHONY: fuzz
fuzz: mod-tidy generate
	go test -fuzz='^Fuzz' -fuzztime=10s -v ./internal/server

.PHONY: cover
cover: mod-tidy generate
	go test -v -count=1 -covermode=atomic -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -html=cover.out

.PHONY: docs
docs:
	go run golang.org/x/pkgsite/cmd/pkgsite@latest -http=:8080
