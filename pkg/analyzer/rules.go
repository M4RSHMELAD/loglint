package analyzer

import (
	"go/ast"
	"go/token"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

// checkLowercaseStart checks if log message starts with lowercase letter
func checkLowercaseStart(pass *analysis.Pass, call *ast.CallExpr, msg string) {
	if !isLowercase(msg) {
		// Get the first argument (the log message string literal)
		if len(call.Args) == 0 {
			return
		}

		arg := call.Args[0]
		lit, ok := arg.(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return
		}

		// Create the fixed version with lowercase first letter
		fixed := msg
		if len(msg) > 0 {
			runes := []rune(msg)
			runes[0] = unicode.ToLower(runes[0])
			fixed = string(runes)
		}

		// Create the fixed string literal with quotes
		fixedLiteral := `"` + fixed + `"`

		pass.Report(analysis.Diagnostic{
			Pos:     call.Pos(),
			Message: "log message should start with lowercase letter",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "Change first letter to lowercase",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     lit.Pos(),
							End:     lit.End(),
							NewText: []byte(fixedLiteral),
						},
					},
				},
			},
		})
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
		// Get the first argument (the log message string literal)
		if len(call.Args) == 0 {
			return
		}

		arg := call.Args[0]
		lit, ok := arg.(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return
		}

		// Remove special characters and emoji
		fixed := removeSpecialChars(msg)

		// Create the fixed string literal with quotes
		fixedLiteral := `"` + fixed + `"`

		pass.Report(analysis.Diagnostic{
			Pos:     call.Pos(),
			Message: "log message should not contain special characters or emoji",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "Remove special characters and emoji",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     lit.Pos(),
							End:     lit.End(),
							NewText: []byte(fixedLiteral),
						},
					},
				},
			},
		})
	}
}

// checkSensitiveData checks if log message contains potentially sensitive data
func checkSensitiveData(pass *analysis.Pass, call *ast.CallExpr, msg string) {
	if containsSensitiveData(msg) {
		pass.Reportf(call.Pos(), "log message may contain sensitive data")
	}
}
