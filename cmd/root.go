package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/handlers"
)

var Root = &cobra.Command{
	Use: "bitclient",
	Run: handlers.Root,
}

func Execute() {
	if err := Root.Execute(); err != nil {
		log.Fatalf("failed to run bitclient cmd: %v", err.Error())
	}
}
