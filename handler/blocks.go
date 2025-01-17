package handler

import (
	"fmt"
	"strconv"

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
	blockhash, err := cmd.Flags().GetString("block")
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
		command.Flags().Set("block", blockhash)
		command.Flags().Set("verbosity", fmt.Sprintf("%d", verbosity))
		command.Run(cmd, args)
	}
}

// Get is a method that handles the 'get' subcommand of the 'blocks' command.
// It retrieves information about a block given its blockhash and verbosity level.
// If no blockhash is provided, it shows the command help.
func (b *blocksHandler) Get(cmd *cobra.Command, args []string) {
	target, ok := getTargetBlock(cmd, args)
	if !ok {
		return
	}

	// Get for default
	this := config.Get().Commands.Blocks.Get.Use
	command := this
	for _, c := range []string{"filter", "stats", "header", "hash"} {
		if cmd.Flags().Changed(c) {
			command = c
			break
		}
	}
	if command != this {
		if command := subcommand(cmd.Parent(), command); command != nil {
			command.Run(cmd, args)
		}
		return
	}

	// Retrieve the verbosity level from flags
	verbosity, err := cmd.Flags().GetInt("verbosity")
	if err != nil {
		logger.Errorf("failed to get verbosity param: %v", err.Error())
		verbosity = 1 // Default verbosity value
	}

	logger.Debugf("getting block with blockhash %v and verbosity %v", target, verbosity)

	response, err := blocks.GetBlock(target, verbosity)
	if err != nil {
		logger.Errorf("failed to get block info: %v", err.Error())
		return
	}

	response.PrintResult()
}

func (b *blocksHandler) Filter(cmd *cobra.Command, args []string) {
	target, ok := getTargetBlock(cmd, args)
	if !ok {
		return
	}

	logger.Debugf("getting block filter with blockhash %v", target)

	response, err := blocks.GetBlockFilter(target)
	if err != nil {
		logger.Errorf("failed to get block filter: %v", err.Error())
		return
	}

	response.Print()
}

func (b *blocksHandler) Hash(cmd *cobra.Command, args []string) {
	target, ok := getTargetBlock(cmd, args)
	if !ok {
		return
	}

	height, err := strconv.Atoi(target)
	if err != nil {
		logger.Errorf("target should be a valid height (numeric)")
		return
	}

	hash, err := blocks.GetBlockHash(height)
	if err != nil {
		logger.Errorf("failed to get block hash: %v", err.Error())
		return
	}

	logger.Print(hash)
}

func (b *blocksHandler) Header(cmd *cobra.Command, args []string) {
	target, ok := getTargetBlock(cmd, args)
	if !ok {
		return
	}

	logger.Debugf("getting block header with blockhash %v", target)

	hex, err := cmd.Flags().GetBool("hex")
	if err != nil {
		logger.Errorf("failed to get hex param: %v", err.Error())
	}

	response, err := blocks.GetBlockHeader(target, !hex)
	if err != nil {
		logger.Errorf("failed to get block header: %v", err.Error())
		return
	}

	response.PrintResult()
}

func (b *blocksHandler) Stats(cmd *cobra.Command, args []string) {
	target, ok := getTargetBlock(cmd, args)
	if !ok {
		return
	}

	stats := []string{}
	if len(args) <= 1 {
		// Get the 'block' flag if no argument is provided
		flag, _ := cmd.Flags().GetStringSlice("stat")
		stats = append(stats, flag...)
	} else {
		stats = append(stats, args[1:]...)
	}

	logger.Debugf("getting block stats %v of target block %v", stats, target)

	logger.Info(stats)

	response, err := blocks.GetBlockStats(target, stats...)
	if err != nil {
		logger.Errorf("failed to get block stats: %v", err.Error())
		return
	}

	response.Print()
}

var getTargetBlock = func(cmd *cobra.Command, args []string) (string, bool) {
	target := ""
	if len(args) <= 0 {
		// Get the 'block' flag if no argument is provided
		flag, _ := cmd.Flags().GetString("block")
		target = flag
	} else {
		// Use the first argument as the block
		target = args[0]
	}

	if target == "" {
		// If no blockhash is provided, show the command help
		if err := cmd.Help(); err != nil {
			logger.Errorf("failed to show output for command %s: %v", cmd.Short, err.Error())
		}
		return "", false
	} else {
		return target, true
	}
}
