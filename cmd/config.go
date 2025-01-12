package cmd

import (
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/config"
)

var Config = &cobra.Command{
	Use:   "config",
	Short: "Print configuration details",
	Long:  "Displays the current configuration details loaded from the config.toml file.",
	Run: func(_ *cobra.Command, _ []string) {
		config.Get().Log()
	},
}

func init() {
	Root.AddCommand(Config)
}
