package config

// Imports
import (
	"encoding/json"
	"os"
)

// Public Structs
type Config struct {
	Version string `json:"version"`
}

// Public Functions
func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config

	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, err
}
