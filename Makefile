BINDIR:=./bin
BINARIES:=$(BINDIR)/arcanumhue

.PHONY: init
init:
	@go mod tidy

.PHONY: lint
lint:
	@golangci-lint run  --config=.golangci.yaml --fix

.PHONY: run
run: init lint
	@go run main.go

.PHONY: run-%
run-%: init lint
	@go run main.go ${@:run-%=%}

.PHONY: build
build: $(BINARIES)

$(BINARIES): init
	@go build -o $@ main.go
