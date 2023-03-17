package missingtypeguard_test

import (
	"testing"

	"github.com/ras0q/missingtypeguard"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)

	t.Run("single", func(t *testing.T) {
		analysistest.Run(t, testdata, missingtypeguard.Analyzer, "a/...")
	})

	t.Run("multi", func(t *testing.T) {
		analysistest.Run(t, testdata, missingtypeguard.Analyzer, "multipackage/...")
	})

	t.Run("standard", func(t *testing.T) {
		analysistest.Run(t, testdata, missingtypeguard.Analyzer, "standardpackage/...")
	})

	t.Run("constructor", func(t *testing.T) {
		analysistest.Run(t, testdata, missingtypeguard.Analyzer, "constructor/...")
	})
}
