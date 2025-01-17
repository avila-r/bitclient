package handler

import (
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/logger"
	"github.com/avila-r/bitclient/network"
)

var Ping = func(cmd *cobra.Command, args []string) {
	if err := network.Ping(); err != nil {
		logger.Errorf("error occurred: %s", err.Error())
	} else {
		logger.Print("pong!")
	}
}
