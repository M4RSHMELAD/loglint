.PHONY: help build test clean install lint

# Default target
help:
	@echo "Available targets:"
	@echo "  make build    - Build the analyzer"
	@echo "  make test     - Run tests"
	@echo "  make lint     - Run linters"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make install  - Install the linter"
	@echo "  make plugin   - Build golangci-lint plugin"

# Build the analyzer
build:
	@echo "Building loglint..."
	go build -o bin/loglint ./cmd/loglint

# Run tests
test:
	@echo "Running tests..."
	go test -v ./pkg/analyzer

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./pkg/analyzer
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Build golangci-lint plugin
plugin:
	@echo "Building golangci-lint plugin..."
	go build -buildmode=plugin -o loglint.so ./plugin

# Install the linter
install:
	@echo "Installing loglint..."
	go install ./cmd/loglint

# Run linters on the project itself
lint:
	@echo "Running linters..."
	go vet ./...
	go fmt ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f bin/loglint
	rm -f loglint.so
	rm -f coverage.out coverage.html

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

# Run example
example:
	@echo "Running analyzer on testdata..."
	go vet -vettool=$$(go build -o /tmp/loglint ./cmd/loglint && echo /tmp/loglint) ./pkg/analyzer/testdata/src/...
