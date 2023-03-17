package main

import (
	"github.com/ras0q/missingtypeguard"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(missingtypeguard.Analyzer) }
