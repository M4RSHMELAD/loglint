package main

import (
	"github.com/M4RSHMELAD/loglint/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

// AnalyzerPlugin is the plugin type for golangci-lint
type AnalyzerPlugin struct{}

// GetAnalyzers returns the list of analyzers provided by this plugin
func (*AnalyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{analyzer.Analyzer}
}

// New returns an instance of the plugin
func New() *AnalyzerPlugin {
	return &AnalyzerPlugin{}
}

// main is a placeholder - this is meant to be built as a plugin
func main() {
	// Plugin entry point - used when building with -buildmode=plugin
}
