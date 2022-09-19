
.PHONY: run
run:
	@go run ./cmd/s3gw

.PHONY: test
test:
	go test ./...

.PHONY: tidy
tidy:
	go mod tidy
