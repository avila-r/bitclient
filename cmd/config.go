package cmd

import (
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/config"
)

var Config = &cobra.Command{
	Use:   config.Get().Commands.Config.Use,
	Short: config.Get().Commands.Config.ShortDescription,
	Long:  config.Get().Commands.Config.LongDescription,
	Run: func(_ *cobra.Command, _ []string) {
		config.Get().Log()
	},
}
