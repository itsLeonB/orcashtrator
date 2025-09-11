# Unit Tests for Internal Packages

This directory contains unit tests for all internal packages in the orcashtrator application.

## Structure

Each package has its own test directory following the pattern `{package_name}_test`:

- `appconstant_test/` - Tests for application constants
- `config_test/` - Tests for configuration loading
- `delivery_test/` - Tests for HTTP delivery layer
- `domain_test/` - Tests for domain entities and logic
- `mapper_test/` - Tests for data mapping functions
- `provider_test/` - Tests for dependency providers
- `service_test/` - Tests for business logic services
- `util_test/` - Tests for utility functions

## Running Tests

### Run all tests
```bash
./run_tests.sh
```

### Run specific package tests
```bash
go test -v ./internal/test/util_test/...
go test -v ./internal/test/service_test/...
```

### Run with coverage
```bash
go test -v -coverprofile=coverage.out ./internal/test/...
go tool cover -html=coverage.out -o coverage.html
```

## Test Guidelines

- Each test file corresponds to a source file (e.g., `uuid_util_test.go` tests `uuid_util.go`)
- Tests use testify's assert package for assertions
- Mocks are created using testify/mock for external dependencies
- Tests aim for high coverage without forcing 100% if not practical
- Tests focus on behavior verification rather than implementation details

## Dependencies

- `github.com/stretchr/testify/assert` - Assertions
- `github.com/stretchr/testify/mock` - Mocking framework

## Coverage Goals

The tests aim for high coverage while maintaining practical and meaningful test cases. Focus is on:

1. Happy path scenarios
2. Error handling
3. Edge cases
4. Business logic validation
5. Data transformation accuracy
