package handlers

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/blocks"
	"github.com/avila-r/bitclient/config"
	"github.com/avila-r/bitclient/logger"
)

type blocksHandler Handler

var Blocks blocksHandler = func(cmd *cobra.Command, args []string) {
	blockhash, err := cmd.Flags().GetString("blockhash")
	if err != nil || blockhash == "" {
		if err := cmd.Help(); err != nil {
			logger.Errorf("failed to show output for command %s: %v", cmd.Short, err.Error())
		}
		return
	}

	verbosity, err := cmd.Flags().GetInt("verbosity")
	if err != nil {
		logger.Errorf("failed to get verbosity param: %v", err.Error())
		verbosity = 1 // Default value
	}

	if command := subcommand(cmd, config.Get().Commands.Blocks.Get.Use); command != nil {
		command.Flags().Set("blockhash", blockhash)
		command.Flags().Set("verbosity", fmt.Sprintf("%d", verbosity))
		command.Run(cmd, args)
	}
}

func (b *blocksHandler) Get(cmd *cobra.Command, args []string) {
	blockhash := ""
	if len(args) <= 0 {
		flag, err := cmd.Flags().GetString("blockhash")
		if err != nil {
			logger.Errorf("blockhash must be provided")
			return
		}
		blockhash = flag
	} else {
		blockhash = args[0]
	}

	if blockhash == "" {
		if err := cmd.Help(); err != nil {
			logger.Errorf("failed to show output for command %s: %v", cmd.Short, err.Error())
		}
		return
	}

	verbosity, err := cmd.Flags().GetInt("verbosity")
	if err != nil {
		logger.Errorf("failed to get verbosity param: %v", err.Error())
		verbosity = 1 // Default value
	}

	logger.Debugf("getting block with blockhash %v and verbosity %v", blockhash, verbosity)

	response, err := blocks.GetBlock(blockhash, verbosity)
	if err != nil {
		logger.Errorf("failed to get block info: %v", err.Error())
		return
	}

	response.PrintResult()
}
