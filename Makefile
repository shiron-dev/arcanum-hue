BINDIR:=./bin
BINARIES:=$(BINDIR)/arcanumhue

.PHONY: init
init:
	@go mod tidy

.PHONY: lint
lint:
	@golangci-lint run  --config=.golangci.yaml --fix

.PHONY: fmt
fmt:
	@go fmt ./...
	@gofumpt -l -w .

.PHONY: run
run: init
	@go run main.go

.PHONY: run-%
run-%: init
	@go run main.go ${@:run-%=%}

.PHONY: build
build: $(BINARIES)

$(BINARIES): init
	@go build -o $@ main.go

.PHONY: gen
gen:
	@go generate ./...
	@go run scripts/main.go

.PHONY: release
release:
	@goreleaser release --snapshot --clean

.PHONY: test
test:
	@go test -v ./...
