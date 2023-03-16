package missingtypeguard_test

import (
	"testing"

	"missingtypeguard"

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
}
