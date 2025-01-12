package cmd

import "github.com/avila-r/bitclient/config"

func init() {
	Root.PersistentFlags().Bool("debug", config.Get().Advanced.Debug, "Enable debug mode")
}
