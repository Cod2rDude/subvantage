package tool

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Cod2rDude/subvantage/internal/types"
	"github.com/Cod2rDude/subvantage/internal/ui"
)

// Public Functions
func Search(opts types.Options) {
	url := fmt.Sprintf("https://crt.sh/?q=%%.%s&output=json", opts.Domain)

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		ui.Log("error", "Error fetching crt.sh, "+opts.Domain)
		return
	}
	defer resp.Body.Close()

	var results []types.CRTResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		ui.Log("error", "Error decoding json.")
	}

	uniqueSubs := make(map[string]bool)
	for _, res := range results {
		for _, name := range strings.Split(res.NameValue, "\n") {
			clean := strings.ToLower(strings.TrimSpace(strings.Replace(name, "*.", "", 1)))
			uniqueSubs[clean] = true
		}
	}

	var sb strings.Builder
	for sub := range uniqueSubs {
		sb.WriteString(sub + "\n")
	}
	outputData := sb.String()

	if opts.OutputFile != "" {
		err := os.WriteFile(opts.OutputFile, []byte(outputData), 0644)
		if err == nil {
			ui.Log("info", fmt.Sprintf("Saved %d subdomains to %s", len(uniqueSubs), opts.OutputFile))
			return
		}
		ui.Log("error", "Could not write to file, printing to terminal instead.")
	}

	fmt.Print(outputData)
}
