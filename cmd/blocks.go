package cmd

import (
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/config"
	"github.com/avila-r/bitclient/handler"
)

// bitclient blocks
var (
	Blocks = &cobra.Command{
		Use:   config.Get().Commands.Blocks.Use,
		Short: config.Get().Commands.Blocks.ShortDescription,
		Long:  config.Get().Commands.Blocks.LongDescription,
		Run:   handler.Blocks,
	}
)

// bitclient blocks get
var (
	BlocksGet = &cobra.Command{
		Use:   config.Get().Commands.Blocks.Get.Use,
		Short: config.Get().Commands.Blocks.Get.ShortDescription,
		Long:  config.Get().Commands.Blocks.Get.LongDescription,
		Run:   handler.Blocks.Get,
	}

	BlocksFilter = &cobra.Command{
		Use:   config.Get().Commands.Blocks.Filter.Use,
		Short: config.Get().Commands.Blocks.Filter.ShortDescription,
		Long:  config.Get().Commands.Blocks.Filter.LongDescription,
		Run:   handler.Blocks.Filter,
	}

	BlocksHash = &cobra.Command{
		Use:   config.Get().Commands.Blocks.Hash.Use,
		Short: config.Get().Commands.Blocks.Hash.ShortDescription,
		Long:  config.Get().Commands.Blocks.Hash.LongDescription,
		Run:   handler.Blocks.Hash,
	}

	BlocksHeader = &cobra.Command{
		Use:   config.Get().Commands.Blocks.Header.Use,
		Short: config.Get().Commands.Blocks.Header.ShortDescription,
		Long:  config.Get().Commands.Blocks.Header.LongDescription,
		Run:   handler.Blocks.Header,
	}

	BlocksStats = &cobra.Command{
		Use:   config.Get().Commands.Blocks.Stats.Use,
		Short: config.Get().Commands.Blocks.Stats.ShortDescription,
		Long:  config.Get().Commands.Blocks.Stats.LongDescription,
		Run:   handler.Blocks.Stats,
	}
)

func init() {
	Root.AddCommand(Blocks) // bitclient blocks
	// Flags
	{
		Blocks.PersistentFlags().StringP("block", "b", "", "Specify the block if has a target block (optional)")
		Blocks.Flags().IntP("verbosity", "v", 1, "Set full response's verbosity level (0-3, default: 0)")
	}

	// Subcommands
	{
		Blocks.AddCommand(BlocksGet) // bitclient blocks get
		// Flags
		{
			BlocksGet.Flags().Bool("filter", false, "Get filter")
			BlocksGet.Flags().Bool("stats", false, "Get stats")
			BlocksGet.Flags().Bool("header", false, "Get header")
			BlocksGet.Flags().Bool("hash", false, "Get blockhash")
			BlocksGet.Flags().Bool("hex", false, "Set to return the block header in hexadecimal encoding")
			BlocksGet.Flags().IntP("verbosity", "v", 1, "Set full response's verbosity level (0-3, default: 0)")
		}

		Blocks.AddCommand(BlocksHeader) // bitclient blocks header
		{
			BlocksHeader.Flags().Bool("hex", false, "Set to return the block header in hexadecimal encoding")
		}

		Blocks.AddCommand(BlocksFilter) // bitclient blocks filter
		Blocks.AddCommand(BlocksHash)   // bitclient blocks hash

		Blocks.AddCommand(BlocksStats) // bitclient blocks stats
		{
			BlocksStats.Flags().StringSliceP("stat", "s", []string{}, "A specific statistic to retrieve.")
		}
	}
}
