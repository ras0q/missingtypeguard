package main

import (
	"missingtypeguard"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(missingtypeguard.Analyzer) }
