.PHONY: test

GO=CGO_ENABLED=0 GO111MODULE=on go

test:
	$(GO) test ./... -cover
	$(GO) vet ./...
