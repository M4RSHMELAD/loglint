package main

import (
	"github.com/M4RSHMELAD/loglint/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

// analyzerPlugin is the plugin type for golangci-lint
type analyzerPlugin struct{}

// GetAnalyzers returns the list of analyzers provided by this plugin
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{analyzer.Analyzer}
}

// AnalyzerPlugin is the exported variable that golangci-lint will look for
var AnalyzerPlugin analyzerPlugin

// main is a placeholder - this is meant to be built as a plugin
func main() {
	// Plugin entry point - used when building with -buildmode=plugin
}
