package main

import (
	"fmt"

	"github.com/Cod2rDude/subvantage/internal/tool"
	"github.com/Cod2rDude/subvantage/internal/types"
	"github.com/Cod2rDude/subvantage/internal/ui"
)

func printOptions(opts types.Options) {
	ui.Log("tool", "Starting to finding subdomains in domain '"+opts.Domain+"'")

	switch opts.Mode {
	case types.ModeBruteForce:
		ui.Log("tool", "Using bruteforce method.")
	case types.ModeCombined:
		ui.Log("tool", "Using bruteforce + search method.")
	case types.ModeSearch:
		ui.Log("tool", "Using search method.")
	default:
		ui.Log("tool", "Using unknown method.")
	}

	if tool.FileExists(opts.OutputFile) {
		ui.Log("tool", "Output will be saved in '"+opts.OutputFile+"'")
	}

	fmt.Println("")
}

// Main
func main() {
	ui.Startup()

	var opts types.Options
	opts.Domain = "google.com"
	opts.Mode = types.ModeSearch
	opts.OutputFile = "test.txt"

	printOptions(opts)

	tool.Search(opts)
}
