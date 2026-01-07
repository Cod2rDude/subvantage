package tool

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Cod2rDude/subvantage/internal/color"
	"github.com/Cod2rDude/subvantage/internal/config"
	"github.com/Cod2rDude/subvantage/internal/types"
	"github.com/Cod2rDude/subvantage/internal/ui"
)

// Private Variables
var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64)...",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)...",
	"Mozilla/5.0 (X11; Linux x86_64)...",
}

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
					time.Sleep(50 * time.Millisecond)
				}
			}
		}
	}()

	url := fmt.Sprintf("https://crt.sh/?q=%%.%s&output=json", opts.Domain)

	clamp := func(n int, min int, max int) int {
		if n < min {
			return min
		} else if n > max {
			return max
		} else {
			return n
		}
	}

	uniqueSubs := make(map[string]bool)
	maxRetries := clamp(config.MaxRetries, 1, 5)
	found := false

	for i := 0; i < maxRetries; i++ {
		client := &http.Client{
			Timeout: 60 * time.Second,
		}

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", userAgents[i%len(userAgents)])
		resp, err := client.Do(req)

		if err != nil {
			fmt.Print("\r\033[2K")

			switch i {
			case maxRetries - 1:
				stop <- true
				ui.Log("error", fmt.Sprintf("Failed to fetch crt.sh after %d retries. Program failed.", maxRetries))
			default:
				ui.Log("error", fmt.Sprintf("Error fetching crt.sh, %v, trying again.", err))
			}

			continue
		} else {
			defer resp.Body.Close()

			fmt.Print("\r\033[2K")
			ui.Log("tool", "Bingo!")
			stop <- true

			var results []types.CRTResult
			json.NewDecoder(resp.Body).Decode(&results)

			for _, res := range results {
				for _, name := range strings.Split(res.NameValue, "\n") {
					clean := strings.ToLower(strings.TrimSpace(strings.Replace(name, "*.", "", 1)))
					uniqueSubs[clean] = true
				}
			}

			found = true

			break
		}
	}

	if !found {
		return
	}

	ui.Log("tool", fmt.Sprintf("Found %d subdomains.", len(uniqueSubs)))

	if FileExists(opts.OutputFile) {
		var sb strings.Builder
		for sub := range uniqueSubs {
			sb.WriteString(sub + "\n")
		}
		outputData := sb.String()

		err := os.WriteFile(opts.OutputFile, []byte(outputData), 0644)
		if err == nil {
			ui.Log("tool", fmt.Sprintf("Saved %d subdomains to %s", len(uniqueSubs), opts.OutputFile))
			return
		}
		ui.Log("error", "Could not write to file, printing to terminal instead.")
	}

	ui.Log("tool", "Writing to terminal.")
	fmt.Println("")
	var c int = 0
	for sub := range uniqueSubs {
		fmt.Println(color.Paint(color.Blue, fmt.Sprintf("[%d] %s", c, sub)))
		c += 1
	}
}
