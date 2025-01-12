package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Properties struct {
	Main struct {
		Command          string `toml:"command"`
		ShortDescription string `toml:"short"`
		LongDescription  string `toml:"long"`
	} `toml:"main"`

	Info struct {
		License    string `toml:"license"`
		Author     string `toml:"author"`
		Version    string `toml:"version"`
		Repository string `toml:"repository"`
	} `toml:"info"`

	Advanced struct {
		Debug bool `toml:"debug"`
	} `toml:"advanced"`
}

var (
	config = func() *Properties {
		cfg := Properties{}

		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("failed to get current working directory: %v", err)
		}

		path := filepath.Join(dir, "config.toml")
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
)

func Get() *Properties {
	return config
}

func (p *Properties) ToString() string {
	return fmt.Sprintf(`
  [Main]:
    Command: %s
    Short Description: %s
    Long Description: %s
	
  [Info]:
    License: %s
    Author: %s
    Version: %s
    Repository: %s

  [Advanced]:
    Debug: %v`,

		p.Main.Command,
		p.Main.ShortDescription,
		p.Main.LongDescription,

		p.Info.License,
		p.Info.Author,
		p.Info.Version,
		p.Info.Repository,

		p.Advanced.Debug,
	)
}

func (p *Properties) Log() {
	log.Printf("Loaded configuration:\n%s", p.ToString())
}
