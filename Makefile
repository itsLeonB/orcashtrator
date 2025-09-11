TEST_DIR := ./internal/test

.PHONY: http http-hot lint test test-verbose test-coverage test-coverage-html test-clean

http:
	go run cmd/http/main.go

http-hot:
	@echo "🚀 Starting HTTP server with hot reload..."
	air --build.cmd "go build -o bin/http cmd/http/main.go" --build.bin "./bin/http"

lint:
	golangci-lint run ./...

test:
	@echo "Running all tests..."
	@if [ -d $(TEST_DIR) ]; then \
		go test $(TEST_DIR)/...; \
	else \
		echo "No tests found in $(TEST_DIR), skipping."; \
	fi

test-verbose:
	@echo "Running all tests with verbose output..."
	@if [ -d $(TEST_DIR) ]; then \
		go test -v $(TEST_DIR)/...; \
	else \
		echo "No tests found in $(TEST_DIR), skipping."; \
	fi

test-coverage:
	@echo "Running all tests with coverage report..."
	@if [ -d $(TEST_DIR) ]; then \
		go test -v -cover -coverprofile=coverage.out -coverpkg=./internal/... $(TEST_DIR)/...; \
	else \
		echo "No tests found in $(TEST_DIR), skipping."; \
	fi

test-coverage-html:
	@echo "Running all tests and generating HTML coverage report..."
	@if [ -d $(TEST_DIR) ]; then \
		go test -v -cover -coverprofile=coverage.out -coverpkg=./internal/... $(TEST_DIR)/... && \
		go tool cover -html=coverage.out -o coverage.html && \
		echo "Coverage report generated: coverage.html"; \
	else \
		echo "No tests found in $(TEST_DIR), skipping."; \
	fi

test-clean:
	@echo "Cleaning test cache and running tests..."
	@if [ -d $(TEST_DIR) ]; then \
		go clean -testcache && go test -v $(TEST_DIR)/...; \
	else \
		echo "No tests found in $(TEST_DIR), skipping."; \
	fi
