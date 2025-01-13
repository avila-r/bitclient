package cmd

import (
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/config"
	"github.com/avila-r/bitclient/logger"
	// "github.com/avila-r/bitclient/handlers"
)

var Root = &cobra.Command{
	Use:   config.Get().Main.Use,
	Short: config.Get().Main.ShortDescription,
	Long:  config.Get().Main.LongDescription,
	// Currently disable interactive mode
	// Run:   handlers.Root,
}

func Execute() {
	if err := Root.Execute(); err != nil {
		logger.Fatalf("failed to run bitclient cmd: %v", err.Error())
	}
}
