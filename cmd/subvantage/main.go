package main

import (
	"github.com/Cod2rDude/subvantage/internal/tool"
	"github.com/Cod2rDude/subvantage/internal/types"
	"github.com/Cod2rDude/subvantage/internal/ui"
)

func main() {
	ui.Startup()

	var opts types.Options
	opts.Domain = "google.com"
	opts.Mode = types.ModeCombined
	opts.OutputFile = "out/test.txt"

	tool.Search(opts)
}
