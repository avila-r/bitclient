package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

// Properties defines the structure for configuration settings in the TOML file
type Properties struct {
	Main command `toml:"main"` // Main configuration settings

	// Info contains metadata about the project
	Info struct {
		License    string `toml:"license"`    // License information
		Author     string `toml:"author"`     // Author of the project
		Version    string `toml:"version"`    // Version of the project
		Repository string `toml:"repository"` // Repository URL
	} `toml:"info"`

	// Advanced contains additional advanced settings for the configuration
	Advanced struct {
		Debug bool `toml:"debug"` // Debug mode setting
	} `toml:"advanced"`

	// Commands contains the definitions for various command configurations
	Commands struct {
		Config command `toml:"config"` // Configuration command settings

		// Blockchain contains blockchain-related command settings
		Blockchain struct {
			command         // General command settings for blockchain
			Info    command `toml:"info"`
		} `toml:"blockchain"`

		// Blocks contains block-related command settings
		Blocks struct {
			command         // General command settings for blocks
			Get     command `toml:"get"`
		} `toml:"blocks"`
	} `toml:"commands"`
}

// command defines the structure for each command in the configuration
type command struct {
	Use              string `toml:"use"`   // Command usage
	ShortDescription string `toml:"short"` // Short description of the command
	LongDescription  string `toml:"long"`  // Detailed description of the command
}

// config is a global variable that loads and parses the configuration file on package initialization
var config = func() *Properties {
	cfg := Properties{} // Initialize an empty Properties struct

	// Define the path to the config.toml file
	path := filepath.Join(RootPath, "config.toml")

	// Check if the config file exists
	if _, err := os.Stat(path); err != nil {
		log.Fatalf("config file not found at %s.", path)
	}

	// Read the content of the config file
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config.toml: %v", err)
	}

	// Parse the content of the file into the Properties struct
	if err := toml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("failed to parse config.toml: %v", err)
	}

	// Return the populated Properties struct
	return &cfg
}()

// Get returns the loaded configuration as a Properties struct
func Get() *Properties {
	return config
}

// ToString serializes the Properties struct into a string in TOML format
func (p *Properties) ToString() string {
	bytes, err := toml.Marshal(p) // Marshal the struct to TOML
	if err != nil {
		slog.Debug("failed to marshal properties into .toml file", "error", err)
	}
	return string(bytes) // Return the serialized string
}

// Log outputs the string representation of the Properties struct to the console
func (p *Properties) Log() {
	fmt.Print(p.ToString()) // Print the serialized TOML string
}
