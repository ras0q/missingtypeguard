package missingtypeguard_test

import (
	"os"
	"testing"

	"github.com/ras0q/missingtypeguard"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)

	childDirs, err := os.ReadDir(testdata + "/src")
	if err != nil {
		t.Fatal(err)
	}

	for _, childDir := range childDirs {
		if childDir.IsDir() {
			t.Run(childDir.Name(), func(t *testing.T) {
				analysistest.Run(t, testdata, missingtypeguard.Analyzer, childDir.Name()+"/...")
			})
		}
	}
}
