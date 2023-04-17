// This must be package main
package main

import (
	"github.com/JiaheEatingOut/expectfirst/pkg/analyzer"

	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{analyzer.Analyzer}
}

var AnalyzerPlugin analyzerPlugin
