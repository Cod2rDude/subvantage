package ui

import (
	_ "embed"
	"fmt"

	"github.com/Cod2rDude/subvantage/internal/color"
	"github.com/Cod2rDude/subvantage/internal/config"
)

//go:embed assets/banner.txt
var banner string

var toolname = "subvantage"
var startUpString1 = "Running '" + toolname + "' " + config.Version

// Public Functions
func Log(option string, message string) {
	fmt.Print("[")

	switch option {
	case "info":
		fmt.Print(color.Paint(color.Blue, "INFO"))
	case "warning":
		fmt.Print(color.Paint(color.Orange, "WARNING"))
	case "error":
		fmt.Print(color.Paint(color.Red, "ERROR"))
	default:
		fmt.Print(color.Paint(color.Blue, "INFO"))
	}

	fmt.Println("] " + message)
}

func Startup() {
	fmt.Println(color.Paint(color.Green, banner))
	fmt.Println(color.Paint(color.Green, "Developed by: ") + color.Paint(color.Blue, "Cod2rDude"))
	fmt.Println("")
	Log("info", startUpString1)
	Log("warning", "Use with caution, you are responsible for your actions.")
	Log("warning", "Developer is not responsible for any malicious use.")
	fmt.Println("")
}
