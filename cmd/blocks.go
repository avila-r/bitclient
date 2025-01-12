package cmd

import (
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/config"
	"github.com/avila-r/bitclient/handlers"
)

// bitclient blocks
var (
	Blocks = &cobra.Command{
		Use:   config.Get().Commands.Blocks.Use,
		Short: config.Get().Commands.Blocks.ShortDescription,
		Long:  config.Get().Commands.Blocks.LongDescription,
		Run:   handlers.Blocks,
	}
)

// bitclient blocks get
var (
	BlocksGet = &cobra.Command{
		Use:   config.Get().Commands.Blocks.Get.Use,
		Short: config.Get().Commands.Blocks.Get.ShortDescription,
		Long:  config.Get().Commands.Blocks.Get.LongDescription,
		Run:   handlers.Blocks.Get,
	}

	blockhash string
	verbosity int
)

func init() {
	// Flags
	Blocks.Flags().StringVarP(&blockhash, "blockhash", "b", "", "Specify the blockhash (required)")
	Blocks.Flags().IntVarP(&verbosity, "verbosity", "v", 1, "Set verbosity level (0-3, default: 0)")

	Root.AddCommand(Blocks) // bitclient blocks
	{
		BlocksGet.Flags().StringVarP(&blockhash, "blockhash", "b", "", "Specify the blockhash (required)")
		BlocksGet.Flags().IntVarP(&verbosity, "verbosity", "v", 1, "Set verbosity level (0-3, default: 0)")

		// Subcommands
		Blocks.AddCommand(BlocksGet) // bitclient blocks get
	}
}
