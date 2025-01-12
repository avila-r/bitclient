package handlers

import (
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/logger"
	"github.com/avila-r/bitclient/rpc"
)

type blockchain Handler

// There's no handler for
// default 'blockchain' command
var Blockchain blockchain = nil

func (h *blockchain) Info(cmd *cobra.Command, args []string) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  rpc.MethodGetBlockchainInfo,
		Params:  rpc.NoParams,
	}

	response, err := rpc.DefaultClient.Do(request)
	if err != nil {
		logger.Errorf("failed to get blockchain info: %v", err.Error())
		return
	}

	logger.Print(response.ToString())
}
