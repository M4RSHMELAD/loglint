# loglint
ATATT3xFfGF0N0VT38PksUEp4xhDV36eeh0As92royDroXmLExNytcL8iWfP2brS_azK_6pI0_4kGGxY7ITLaQ_c1t-51b176jEMJXmLK21WsNnPVzYWJhx21qQoWkvZz60n13QTjOdzxbUzd1w1p3TNOHrrne0QWEI80mG98u4NWSOBtICbGOU=24DA6FDA

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
go install github.com/yourusername/loglint@latest
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
# Check for issues
go vet -vettool=$(which loglint) ./...

# Check and apply automatic fixes
go vet -vettool=$(which loglint) -fix ./...
```

### With golangci-lint

```bash
golangci-lint run
```

## Auto-Fixes

`loglint` supports automatic fixes for the following rules:

1. **Lowercase Start**: Automatically converts the first letter to lowercase
   - Before: `slog.Info("Starting server")`
   - After: `slog.Info("starting server")`

2. **No Special Characters**: Removes special characters and emoji
   - Before: `slog.Info("server started!")`
   - After: `slog.Info("server started")`

To apply fixes automatically:

```bash
# Using go vet
go vet -vettool=$(which loglint) -fix ./...

# Preview changes without modifying files
go vet -vettool=$(which loglint) -fix -diff ./...
```

**Note:** Rules for "English only" and "Sensitive data" do not have auto-fixes as they require manual intervention.

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

## IDE Integration (GoLand / IntelliJ IDEA)

You can integrate `loglint` as an External Tool in GoLand for quick access.

### Setup External Tools

1. **Open Settings**: Press `Ctrl+Alt+S` (Windows/Linux) or `Cmd+,` (macOS)
2. Navigate to: **Tools → External Tools**
3. Click the **"+"** button to add a new tool

### Configuration Options

#### Option 1: Check Current Package (No Auto-Fix)

```
Name: loglint (check)
Description: Check log messages in current package
Program: C:\Users\YourUsername\GolandProjects\loglint\loglint.exe
Arguments: .
Working directory: $FileDir$
☑ Synchronize files after execution
☑ Open console for tool output
```

#### Option 2: Fix Current Package (Auto-Fix) ⭐ Recommended

```
Name: loglint (fix current)
Description: Auto-fix log messages in current package
Program: C:\Users\YourUsername\GolandProjects\loglint\loglint.exe
Arguments: -fix .
Working directory: $FileDir$
☑ Synchronize files after execution
☑ Open console for tool output
```

#### Option 3: Fix Entire Project (Auto-Fix All)

```
Name: loglint (fix all)
Description: Auto-fix log messages in entire project
Program: C:\Users\YourUsername\GolandProjects\loglint\loglint.exe
Arguments: -fix ./...
Working directory: $ProjectFileDir$
☑ Synchronize files after execution
☑ Open console for tool output
```

### Usage in IDE

After setup:

1. **Right-click** on any `.go` file or package folder
2. Select **External Tools** → **loglint (fix current)**
3. The tool will automatically fix capitalization issues!

**Tip:** You can also assign keyboard shortcuts to these tools:
- **Settings → Keymap → External Tools → loglint**
- Suggested: `Ctrl+Alt+L` for "loglint (fix current)"

### Platform-Specific Paths

**Windows:**
```
Program: C:\Users\YourUsername\GolandProjects\loglint\loglint.exe
```

**macOS/Linux:**
```
Program: /Users/yourusername/GolandProjects/loglint/loglint
```

**Cross-Platform (if loglint installed globally):**
```
Program: loglint
```

---

## Development

### Prerequisites

- Go 1.25
- golangci-lint (optional)

### Building

```bash
# Build the analyzer
go build -o loglint.exe ./cmd/loglint

# Build the plugin (Linux/macOS only)
go build -buildmode=plugin -o loglint.so ./plugin
```

### Testing

```bash
# Run tests
go test ./pkg/analyzer -v

# Test on current directory
./loglint.exe .

# Test with auto-fix
./loglint.exe -fix .
```

### Project Structure

```
loglint/
├── bin/
│   └── loglint.exe   
├── pkg/analyzer/       # Main analyzer implementation
│   ├── analyzer.go     # Analyzer entry point
│   ├── rules.go        # Rule implementations
│   ├── loggers.go      # Logger detection logic
│   ├── helpers.go      # Helper functions
│   ├── analyzer_test.go # Tests
│   └── testdata/       # Test files
├── plugin/             # golangci-lint plugin
│   └── loglint.go
├── cmd/loglint/        # CLI
│   └── main.go
├── go.mod
└── README.md
```

