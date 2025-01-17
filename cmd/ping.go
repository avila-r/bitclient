package cmd

import (
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/config"
	"github.com/avila-r/bitclient/handler"
)

var Ping = &cobra.Command{
	Use:   config.Get().Commands.Ping.Use,
	Short: config.Get().Commands.Ping.ShortDescription,
	Long:  config.Get().Commands.Ping.LongDescription,
	Run:   handler.Ping,
}

func init() {
	Root.AddCommand(Ping)
	Network.AddCommand(Ping)
}
