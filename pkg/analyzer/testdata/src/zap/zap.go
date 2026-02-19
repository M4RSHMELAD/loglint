package zap

import "go.uber.org/zap"

func ZapExamples() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Different zap methods - FAIL
	logger.Info("Starting server")             // want "log message should start with lowercase letter"
	logger.Error("Database connection failed") // want "log message should start with lowercase letter"
	logger.Debug("Debug information")          // want "log message should start with lowercase letter"
	logger.Warn("Memory usage high")           // want "log message should start with lowercase letter"

	// Correct usage - PASS
	logger.Info("starting server")
	logger.Error("database connection failed")
	logger.Debug("debug information")
	logger.Warn("memory usage high")

	// Special characters - FAIL
	logger.Info("connection established!!!") // want "log message should not contain special characters or emoji"
	logger.Error("error occurred?")          // want "log message should not contain special characters or emoji"

	// Sensitive data - FAIL
	logger.Debug("api-key verification")      // want "log message may contain sensitive data"
	logger.Info("password validation passed") // want "log message may contain sensitive data"
	logger.Error("private_key not found")     // want "log message may contain sensitive data"
}
