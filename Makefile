.PHONY: run lint

run:
	@go run main.go

lint:
	@golangci-lint run
