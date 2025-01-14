package handlers

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/blocks"
	"github.com/avila-r/bitclient/config"
	"github.com/avila-r/bitclient/logger"
)

// blocksHandler is a custom handler type based on the Handler function type.
type blocksHandler Handler

// Blocks is a variable representing the handler for the 'blocks' command.
var Blocks blocksHandler = func(cmd *cobra.Command, args []string) {
	// Retrieve the 'blockhash' flag from the command input
	blockhash, err := cmd.Flags().GetString("blockhash")
	if err != nil || blockhash == "" {
		// If blockhash is missing or an error occurs, display command help
		if err := cmd.Help(); err != nil {
			logger.Errorf("failed to show output for command %s: %v", cmd.Short, err.Error())
		}
		return
	}

	// Retrieve the 'verbosity' flag from the command input
	verbosity, err := cmd.Flags().GetInt("verbosity")
	if err != nil {
		// If verbosity is not set, log the error and set default verbosity
		logger.Errorf("failed to get verbosity param: %v", err.Error())
		verbosity = 1 // Default verbosity value
	}

	// Find the 'get' subcommand under the 'blocks' command and set the flags
	if command := subcommand(cmd, config.Get().Commands.Blocks.Get.Use); command != nil {
		command.Flags().Set("blockhash", blockhash)
		command.Flags().Set("verbosity", fmt.Sprintf("%d", verbosity))
		command.Run(cmd, args)
	}
}

// Get is a method that handles the 'get' subcommand of the 'blocks' command.
// It retrieves information about a block given its blockhash and verbosity level.
// If no blockhash is provided, it shows the command help.
func (b *blocksHandler) Get(cmd *cobra.Command, args []string) {
	blockhash := ""
	if len(args) <= 0 {
		// Get the 'blockhash' flag if no argument is provided
		flag, err := cmd.Flags().GetString("blockhash")
		if err != nil {
			logger.Errorf("blockhash must be provided")
			return
		}
		blockhash = flag
	} else {
		// Use the first argument as the blockhash
		blockhash = args[0]
	}

	// If no blockhash is provided, show the command help
	if blockhash == "" {
		if err := cmd.Help(); err != nil {
			logger.Errorf("failed to show output for command %s: %v", cmd.Short, err.Error())
		}
		return
	}

	// Retrieve the verbosity level from flags
	verbosity, err := cmd.Flags().GetInt("verbosity")
	if err != nil {
		logger.Errorf("failed to get verbosity param: %v", err.Error())
		verbosity = 1 // Default verbosity value
	}

	logger.Debugf("getting block with blockhash %v and verbosity %v", blockhash, verbosity)

	response, err := blocks.GetBlock(blockhash, verbosity)
	if err != nil {
		logger.Errorf("failed to get block info: %v", err.Error())
		return
	}

	response.PrintResult()
}
