package handlers

import (
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/blocks"
	"github.com/avila-r/bitclient/logger"
)

type blockchainHandler Handler

// There's no handler for
// default 'blockchain' command
var Blockchain blockchainHandler = nil

func (h *blockchainHandler) Info(cmd *cobra.Command, args []string) {
	response, err := blocks.GetBlockchainInfo()
	if err != nil {
		logger.Errorf("failed to get blockchain info: %v", err.Error())
		return
	}

	logger.Print(response.ToString())
}
