# loglint

A Go linter for checking log messages compliance with best practices.

## Features

`loglint` checks your log messages for the following rules:

1. **Lowercase Start**: Log messages should start with a lowercase letter
2. **English Only**: Log messages should be in English only (no Cyrillic or other non-Latin characters)
3. **No Special Characters**: Log messages should not contain special characters or emoji
4. **No Sensitive Data**: Log messages should not contain potentially sensitive data (passwords, tokens, API keys, etc.)

## Supported Loggers

- `log/slog` (standard library)
- `go.uber.org/zap` (planned)

## Installation

### As a standalone tool

```bash
go install github.com/M4RSHMELAD/loglint@latest
```

### As golangci-lint plugin

Add to your `.golangci.yml`:

```yaml
linters-settings:
  custom:
    loglint:
      path: ./plugin/loglint.so
      description: Checks log messages for common issues
```

## Usage

### Standalone

```bash
go vet -vettool=$(which loglint) ./...
```

### With golangci-lint

```bash
golangci-lint run
```

## Examples

### ❌ Bad

```go
import "log/slog"

slog.Info("Starting server")                    // Should start with lowercase
slog.Error("ошибка подключения")                 // Should be in English
slog.Info("server started!")                     // No special characters
slog.Debug("password: " + password)              // Contains sensitive data
```

### ✅ Good

```go
import "log/slog"

slog.Info("starting server")
slog.Error("connection failed")
slog.Info("server started")
slog.Debug("user authenticated successfully")
```

## Rules Details

### Rule 1: Lowercase Start

Log messages should start with a lowercase letter to maintain consistency.

**Why?** This is a common convention in many projects and makes logs more uniform.

```go
// ❌ Bad
slog.Info("Starting server")
slog.Error("Failed to connect")

// ✅ Good
slog.Info("starting server")
slog.Error("failed to connect")
```

### Rule 2: English Only

Log messages should use only English characters (basic Latin alphabet).

**Why?** English is the universal language for technical logs, making them accessible to international teams.

```go
// ❌ Bad
slog.Info("запуск сервера")
slog.Error("ошибка подключения")

// ✅ Good
slog.Info("starting server")
slog.Error("connection failed")
```

### Rule 3: No Special Characters or Emoji

Log messages should not contain special characters (!, ?, ...) or emoji.

**Why?** Special characters can interfere with log parsing and don't add value to structured logs.

```go
// ❌ Bad
slog.Info("server started!")
slog.Error("connection failed!!!")
slog.Info("deployment successful 🚀")

// ✅ Good
slog.Info("server started")
slog.Error("connection failed")
slog.Info("deployment successful")
```

**Note:** Basic punctuation like `.`, `,`, `:`, `-`, `_`, `/` are allowed.

### Rule 4: No Sensitive Data

Log messages should not contain potentially sensitive information.

**Why?** Logs are often stored in plain text and can be accessed by various people/systems.

**Detected keywords:**
- password, pwd, passwd
- token, jwt, bearer
- api_key, apikey, api-key
- secret, private_key, private-key
- credit_card, card_number
- ssn, social_security

```go
// ❌ Bad
slog.Info("user password: " + password)
slog.Debug("api_key=" + apiKey)
slog.Info("token: " + token)

// ✅ Good
slog.Info("user authenticated successfully")
slog.Debug("api request completed")
slog.Info("token validated")
```

## Development

### Prerequisites

- Go 1.25
- golangci-lint (optional)

### Building

```bash
# Build the analyzer
go build ./pkg/analyzer

# Build the plugin
go build -buildmode=plugin -o loglint.so ./plugin
```

### Testing

```bash
# Run tests
go test ./pkg/analyzer -v

# Run linter on the project itself
go vet -vettool=$(which golangci-lint) ./...
```

### Project Structure

```
loglint/
├── pkg/analyzer/       # Main analyzer implementation
│   ├── analyzer.go     # Analyzer entry point
│   ├── rules.go        # Rule implementations
│   ├── loggers.go      # Logger detection logic
│   ├── helpers.go      # Helper functions
│   ├── analyzer_test.go # Tests
│   └── testdata/       # Test files
├── plugin/             # golangci-lint plugin
│   └── loglint.go
├── cmd/loglint/        # CLI (optional)
│   └── main.go
├── go.mod
└── README.md
```


