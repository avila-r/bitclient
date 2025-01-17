package handler

import (
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/logger"
	"github.com/avila-r/bitclient/network"
)

type networkHandler Handler

var Network networkHandler = nil

func (n *networkHandler) Activate(cmd *cobra.Command, args []string) {
	if err := network.SetNetworkActive(true); err != nil {
		logger.Errorf("failed to activate network activity: %s", err.Error())
	} else {
		logger.Info("network activity enabled!")
	}
}

func (n *networkHandler) Deactivate(cmd *cobra.Command, args []string) {
	if err := network.SetNetworkActive(false); err != nil {
		logger.Errorf("failed to deactivate network activity: %s", err.Error())
	} else {
		logger.Info("network activity disabled!")
	}
}

func (n *networkHandler) Connections(cmd *cobra.Command, args []string) {
	response, err := network.GetConnectionCount()
	if err != nil {
		logger.Errorf("failed to get connection count: %s", err.Error())
		return
	}

	count := 0
	if err := response.Bind(&count); err != nil {
		logger.Errorf("failed to get rpc response's result: %s", err.Error())
		return
	}
	logger.Print(count)
}

func (n *networkHandler) Traffic(cmd *cobra.Command, args []string) {
	traffic, err := network.InspectTraffic()
	if err != nil {
		logger.Errorf("failed to inspect network traffic: %s", err.Error())
		return
	}

	traffic.Print()
}

func (n *networkHandler) Info(cmd *cobra.Command, args []string) {
	info, err := network.GetNetworkInfo()
	if err != nil {
		logger.Errorf("failed to inspect network info: %s", err.Error())
		return
	}

	info.Print()
}

func (n *networkHandler) Peers(cmd *cobra.Command, args []string) {
	peers, err := network.GetPeers()
	if err != nil {
		logger.Errorf("failed to get peers: %s", err.Error())
		return
	}

	peers.Print()
}

func (n *networkHandler) Ban(cmd *cobra.Command, args []string) {
}

func (n *networkHandler) Unban(cmd *cobra.Command, args []string) {
}

func (n *networkHandler) Blacklist(cmd *cobra.Command, args []string) {
}
