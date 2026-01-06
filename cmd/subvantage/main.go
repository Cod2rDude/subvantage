package main

import (
	"fmt"

	"github.com/Cod2rDude/subvantage/internal/tool"
	"github.com/Cod2rDude/subvantage/internal/types"
	"github.com/Cod2rDude/subvantage/internal/ui"
)

func printOptions(opts types.Options) {
	ui.Log("info", "Starting to finding subdomains in domain '"+opts.Domain+"'")

	switch opts.Mode {
	case types.ModeBruteForce:
		ui.Log("info", "Using bruteforce method.")
	case types.ModeCombined:
		ui.Log("info", "Using bruteforce + search method.")
	case types.ModeSearch:
		ui.Log("info", "Using search method.")
	default:
		ui.Log("info", "Using unknown method.")
	}

	if tool.FileExists(opts.OutputFile) {
		ui.Log("info", "Output will be saved in '"+opts.OutputFile+"'")
	}

	fmt.Println("")
}

// Main
func main() {
	ui.Startup()

	var opts types.Options
	opts.Domain = "google.com"
	opts.Mode = types.ModeSearch
	opts.OutputFile = "blabla/test.txt"

	printOptions(opts)

	tool.Search(opts)
}
