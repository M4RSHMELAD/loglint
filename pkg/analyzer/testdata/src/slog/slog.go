package slog

import "log/slog"

func SlogExamples() {
	// Different slog methods
	slog.Info("starting application") // want "log message should start with lowercase letter"
	slog.Error("failed to connect")   // want "log message should start with lowercase letter"
	slog.Debug("debug mode enabled")  // want "log message should start with lowercase letter"
	slog.Warn("warning message")      // want "log message should start with lowercase letter"

	// Correct usage
	slog.Info("application started")
	slog.Error("failed to connect")
	slog.Debug("debug mode enabled")
	slog.Warn("warning message")

	// With emoji (triggers both English and special char checks)
	slog.Info("server started") // want "log message should be in English only" "log message should not contain special characters or emoji"

	// Sensitive data
	slog.Debug("jwt token generated")        // want "log message may contain sensitive data"
	slog.Info("secret configuration loaded") // want "log message may contain sensitive data"
}
