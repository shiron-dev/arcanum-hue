BINDIR:=./bin
BINARIES:=$(BINDIR)/arcanumhue

.PHONY: init
init:
	@go mod tidy

.PHONY: run
run:
	@go run main.go

.PHONY: build
build: $(BINARIES)

$(BINARIES): init
	@go build -o $@ main.go
