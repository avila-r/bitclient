package cmd

import (
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/config"
	"github.com/avila-r/bitclient/handler"
)

var (
	// bitclient blockchain
	Blockchain = &cobra.Command{
		Use:   config.Get().Commands.Blockchain.Use,
		Short: config.Get().Commands.Blockchain.ShortDescription,
		Long:  config.Get().Commands.Blockchain.LongDescription,
	}

	// bitclient blockchain info
	BlockchainInfo = &cobra.Command{
		Use:   config.Get().Commands.Blockchain.Info.Use,
		Short: config.Get().Commands.Blockchain.Info.ShortDescription,
		Long:  config.Get().Commands.Blockchain.Info.LongDescription,
		Run:   handler.Blockchain.Info,
	}
)

func init() {
	Root.AddCommand(Blockchain) // bitclient blockchain
	{
		// Subcommands
		Blockchain.AddCommand(BlockchainInfo) // bitclient blockchain info
	}
}
