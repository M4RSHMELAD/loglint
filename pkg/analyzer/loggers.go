package analyzer

import (
	"go/ast"
	"strings"
)

// isLogCall checks if the call expression is a logging function
func isLogCall(call *ast.CallExpr) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	// Check method name (Info, Error, Debug, Warn, etc.)
	methodName := sel.Sel.Name
	logMethods := []string{"Info", "Error", "Debug", "Warn", "Fatal", "Panic", "Print", "Println", "Printf"}
	
	isLogMethod := false
	for _, method := range logMethods {
		if methodName == method || strings.HasSuffix(methodName, method) {
			isLogMethod = true
			break
		}
	}
	
	if !isLogMethod {
		return false
	}

	// Check if it's from slog or zap package
	return isSlogCall(sel) || isZapCall(sel)
}

// isSlogCall checks if this is a log/slog call
func isSlogCall(sel *ast.SelectorExpr) bool {
	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	// Check if package name is "slog" or "log"
	return ident.Name == "slog" || ident.Name == "log"
}

// isZapCall checks if this is a zap logger call
func isZapCall(sel *ast.SelectorExpr) bool {
	// For zap, it's usually: logger.Info() where logger is *zap.Logger
	// We need to check the type, but for simplicity we can check common patterns
	
	// Check if X is an identifier that might be a logger
	if ident, ok := sel.X.(*ast.Ident); ok {
		// Common logger variable names
		loggerNames := []string{"logger", "log", "zap", "l"}
		for _, name := range loggerNames {
			if ident.Name == name {
				return true
			}
		}
	}
	
	return false
}

// getLogMessage extracts the log message string from a call expression
func getLogMessage(call *ast.CallExpr) (string, bool) {
	if len(call.Args) == 0 {
		return "", false
	}

	// First argument should be the message
	arg := call.Args[0]

	// Handle basic string literal
	if lit, ok := arg.(*ast.BasicLit); ok {
		if lit.Kind == 9 { // token.STRING
			// Remove quotes
			msg := lit.Value
			if len(msg) >= 2 {
				msg = msg[1 : len(msg)-1] // Remove surrounding quotes
			}
			return msg, true
		}
	}

	return "", false
}
