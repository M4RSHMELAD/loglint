package basic

import "log/slog"

func Examples() {
	// Rule 1: Lowercase start - FAIL
	slog.Info("Starting server") // want "log message should start with lowercase letter"

	// Rule 1: Lowercase start - PASS
	slog.Info("starting server")

	// Rule 2: English only - FAIL
	slog.Error("ошибка подключения") // want "log message should be in English only"

	// Rule 2: English only - PASS
	slog.Error("connection failed")

	// Rule 3: Special chars - FAIL
	slog.Info("server started!") // want "log message should not contain special characters or emoji"
	slog.Warn("warning!!!")      // want "log message should not contain special characters or emoji"

	// Rule 3: Special chars - PASS
	slog.Info("server started")

	// Rule 4: Sensitive data - FAIL
	slog.Debug("password: secret123") // want "log message may contain sensitive data"
	slog.Info("api_key received")     // want "log message may contain sensitive data"
	slog.Error("token expired")       // want "log message may contain sensitive data"

	// Rule 4: Sensitive data - PASS
	slog.Info("user authenticated")
	slog.Debug("request completed")

	// Multiple violations
	slog.Error("Password Invalid!") // want "log message should start with lowercase letter" "log message should not contain special characters or emoji" "log message may contain sensitive data"
}
