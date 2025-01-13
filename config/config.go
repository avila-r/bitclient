package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Properties struct {
	Main command `toml:"main"`

	Info struct {
		License    string `toml:"license"`
		Author     string `toml:"author"`
		Version    string `toml:"version"`
		Repository string `toml:"repository"`
	} `toml:"info"`

	Advanced struct {
		Debug bool `toml:"debug"`
	} `toml:"advanced"`

	Commands struct {
		Config command `toml:"config"`

		Blockchain struct {
			command
			Info command `toml:"info"`
		} `toml:"blockchain"`

		Blocks struct {
			command
			Get command `toml:"get"`
		} `toml:"blocks"`
	} `toml:"commands"`
}

type command struct {
	Use              string `toml:"use"`
	ShortDescription string `toml:"short"`
	LongDescription  string `toml:"long"`
}

var config = func() *Properties {
	cfg := Properties{}

	path := filepath.Join(RootPath, "config.toml")
	if _, err := os.Stat(path); err != nil {
		log.Fatalf("config file not found at %s.", path)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config.toml: %v", err)
	}

	if err := toml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("failed to parse config.toml: %v", err)
	}
	return &cfg
}()

func Get() *Properties {
	return config
}

func (p *Properties) ToString() string {
	bytes, err := toml.Marshal(p)
	if err != nil {
		slog.Debug("failed to marshal properties into .toml file", "error", err)
	}
	return string(bytes)
}

func (p *Properties) Log() {
	fmt.Print(p.ToString())
}
