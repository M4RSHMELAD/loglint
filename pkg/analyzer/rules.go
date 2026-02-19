package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// checkLowercaseStart checks if log message starts with lowercase letter
func checkLowercaseStart(pass *analysis.Pass, call *ast.CallExpr, msg string) {
	if !isLowercase(msg) {
		pass.Reportf(call.Pos(), "log message should start with lowercase letter")
	}
}

// checkEnglishOnly checks if log message contains only English characters
func checkEnglishOnly(pass *analysis.Pass, call *ast.CallExpr, msg string) {
	if containsNonEnglish(msg) {
		pass.Reportf(call.Pos(), "log message should be in English only")
	}
}

// checkNoSpecialChars checks if log message contains special characters or emoji
func checkNoSpecialChars(pass *analysis.Pass, call *ast.CallExpr, msg string) {
	if containsSpecialChars(msg) {
		pass.Reportf(call.Pos(), "log message should not contain special characters or emoji")
	}
}

// checkSensitiveData checks if log message contains potentially sensitive data
func checkSensitiveData(pass *analysis.Pass, call *ast.CallExpr, msg string) {
	if containsSensitiveData(msg) {
		pass.Reportf(call.Pos(), "log message may contain sensitive data")
	}
}
