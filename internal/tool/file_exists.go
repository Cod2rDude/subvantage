package tool

import (
	"os"
	"path/filepath"
	"strings"
)

// Public Functions
func FileExists(filename string) bool {
	if filename != "" {
		if !strings.HasSuffix(strings.ToLower(filename), ".txt") {
			//ui.Log("error", "Output file must be a .txt file")
			return false
		}
		dir := filepath.Dir(filename)
		info, err := os.Stat(dir)

		if os.IsNotExist(err) {
			//ui.Log("error", "The directory does not exist: "+dir)
			return false
		}

		if err == nil && !info.IsDir() {
			//ui.Log("error", "The path "+dir+" is a file, not a directory")
			return false
		}
	} else {
		return false
	}

	return true
}
