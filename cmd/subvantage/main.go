package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Cod2rDude/subvantage/internal/types"
	"github.com/Cod2rDude/subvantage/internal/ui"
)

func main() {
	var (
		searchDomain string
		bruteDomain  string
		bothDomain   string
		outputFile   string
	)

	flag.StringVar(&searchDomain, "s", "", "Search mode: Scrape search engines/APIs for domain")
	flag.StringVar(&bruteDomain, "b", "", "Bruteforce mode: Dictionary attack on domain")
	flag.StringVar(&bothDomain, "c", "", "Combine mode: Run both search and bruteforce")
	flag.StringVar(&outputFile, "o", "", "Save results to file")

	flag.Parse()

	opts := types.Options{
		OutputFile: outputFile,
		Mode:       types.ModeNone,
	}

	if searchDomain != "" {
		opts.Domain = searchDomain
		opts.Mode = types.ModeSearch
	} else if bruteDomain != "" {
		opts.Domain = bruteDomain
		opts.Mode = types.ModeBruteForce
	} else if bothDomain != "" {
		opts.Domain = bothDomain
		opts.Mode = types.ModeCombined
	}

	ui.Startup()

	if opts.Mode == types.ModeNone {
		opts = ui.GetInteractiveInput()
	}

	if opts.Domain == "" {
		ui.Log("error", "No domain provided.")
		os.Exit(1)
	}

	runTool(opts)
}
func runTool(opts types.Options) {
	fmt.Println("")
	ui.Log("info", "Starting Scan on: "+opts.Domain)

	if opts.OutputFile != "" {
		ui.Log("info", "Output will be saved to: "+opts.OutputFile)
	}

	switch opts.Mode {
	case types.ModeSearch:
		ui.Log("info", "Starting search...")
	case types.ModeBruteForce:
		ui.Log("info", "Starting brute force...")
	case types.ModeCombined:
		ui.Log("info", "Starting combined...")
	}
}
