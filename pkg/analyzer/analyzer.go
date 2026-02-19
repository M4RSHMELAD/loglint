package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer is the main loglint analyzer
var Analyzer = &analysis.Analyzer{
	Name:     "loglint",
	Doc:      "checks log messages for common issues",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// Filter only function calls
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)

		// Check if this is a logging function call
		if !isLogCall(call) {
			return
		}

		// Get the log message (first string argument)
		msg, ok := getLogMessage(call)
		if !ok {
			return
		}

		// Apply all rules
		checkLowercaseStart(pass, call, msg)
		checkEnglishOnly(pass, call, msg)
		checkNoSpecialChars(pass, call, msg)
		checkSensitiveData(pass, call, msg)
	})

	return nil, nil
}
