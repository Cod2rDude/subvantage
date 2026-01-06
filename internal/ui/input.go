package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Cod2rDude/subvantage/internal/types"
)

// Public Functions
func GetInteractiveInput() types.Options {
	reader := bufio.NewReader(os.Stdin)
	var opts types.Options

	fmt.Println("Enter a domain, ")
	fmt.Print(">> ")
	domain, _ := reader.ReadString('\n')
	opts.Domain = strings.TrimSpace(domain)

	fmt.Println("\nSelect mode,")
	fmt.Println("[1] Search engine")
	fmt.Println("[2] Brute Force")
	fmt.Println("[3] All Combined")
	fmt.Print(">> ")

	modeStr, _ := reader.ReadString('\n')
	modeStr = strings.TrimSpace(modeStr)

	switch modeStr {
	case "2":
		opts.Mode = types.ModeBruteForce
	case "3":
		opts.Mode = types.ModeCombined
	default:
		opts.Mode = types.ModeSearch
	}

	fmt.Println("\nOutput file, (empty to skip)")
	fmt.Print(">> ")
	output, _ := reader.ReadString('\n')
	opts.OutputFile = strings.TrimSpace(output)

	return opts
}
