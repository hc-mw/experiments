test-wallet:
	@go test ./wallet -v
test-wallet-cover:
	@go test ./wallet -v -cover
test-iteration:
	@go test ./iteration -v
benchmark-iteration:
	@go test ./iteration -bench=. -v
test-di:
	@go test ./dependency_injection -v

test: test-wallet test-iteration

benchmark: benchmark-iteration