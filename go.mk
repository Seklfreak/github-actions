.PHONY: go-lint
go-lint:
	go vet ./...

.PHONY: go-test
go-test:
	go test ./...