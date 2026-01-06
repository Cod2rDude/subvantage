package tool

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Cod2rDude/subvantage/internal/color"
	"github.com/Cod2rDude/subvantage/internal/types"
	"github.com/Cod2rDude/subvantage/internal/ui"
)

// Public Functions
func Search(opts types.Options) {
	stop := make(chan bool)
	go func() {
		chars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		for {
			for _, char := range chars {
				select {
				case <-stop:
					return
				default:
					fmt.Printf("\r| %s Finding subdomains for %s...", char, opts.Domain)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()

	url := fmt.Sprintf("https://crt.sh/?q=%%.%s&output=json", opts.Domain)

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	resp, err := client.Do(req)

	stop <- true
	fmt.Println("")

	if err != nil {
		ui.Log("error", "Error fetching crt.sh, "+opts.Domain)
		return
	}
	defer resp.Body.Close()

	var results []types.CRTResult
	json.NewDecoder(resp.Body).Decode(&results)

	uniqueSubs := make(map[string]bool)
	for _, res := range results {
		for _, name := range strings.Split(res.NameValue, "\n") {
			clean := strings.ToLower(strings.TrimSpace(strings.Replace(name, "*.", "", 1)))
			uniqueSubs[clean] = true
		}
	}

	//TODO: Retry
	ui.Log("info", "Bingo!")
	ui.Log("info", fmt.Sprintf("Found %d subdomains.", len(uniqueSubs)))

	if FileExists(opts.OutputFile) {
		var sb strings.Builder
		for sub := range uniqueSubs {
			sb.WriteString(sub + "\n")
		}
		outputData := sb.String()

		err := os.WriteFile(opts.OutputFile, []byte(outputData), 0644)
		if err == nil {
			ui.Log("info", fmt.Sprintf("Saved %d subdomains to %s", len(uniqueSubs), opts.OutputFile))
			return
		}
		ui.Log("error", "Could not write to file, printing to terminal instead.")
	}

	ui.Log("info", "Writing to terminal.")
	fmt.Println("")
	var c int = 0
	for sub := range uniqueSubs {
		fmt.Println(color.Paint(color.Blue, fmt.Sprintf("[%d] %s", c, sub)))
		c += 1
	}
}
