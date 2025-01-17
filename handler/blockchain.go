package handler

import (
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/blocks"
	"github.com/avila-r/bitclient/logger"
)

// blockchainHandler is a custom handler type based on the Handler function type.
type blockchainHandler Handler

// Blockchain is a variable representing the handler for the 'blockchain' command.
// This handler is currently set to nil, meaning there's no handler defined for this command by default.
var Blockchain blockchainHandler = nil

// Info is a method that handles the 'info' subcommand of the 'blockchain' command.
// It retrieves information about the blockchain by calling blocks.GetBlockchainInfo()
// and logs the result or any errors encountered.
func (h *blockchainHandler) Info(cmd *cobra.Command, args []string) {
	response, err := blocks.GetBlockchainInfo()
	if err != nil {
		logger.Errorf("failed to get blockchain info: %v", err.Error())
		return
	}

	logger.Print(response.ToString())
}
