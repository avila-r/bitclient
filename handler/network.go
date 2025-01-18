package handler

import (
	"net"
	"strings"

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
	target, ok := getTargetIP(cmd, args)
	if !ok {
		return
	}

	if _, _, err := net.ParseCIDR(target); err != nil && net.ParseIP(target) == nil {
		logger.Errorf("invalid format. the target must be either a valid IP address (e.g., '192.168.1.1') or a valid subnet mask (e.g., '192.168.1.0/24)")
		return
	}

	time, err := cmd.Flags().GetInt("time")
	if err != nil {
		logger.Errorf("failed to unwrap time flag: %v", err.Error())
	}

	absolute, err := cmd.Flags().GetBool("absolute")
	if err != nil {
		logger.Errorf("failed to unwrap absolute flag: %v", err.Error())
	}

	ban := network.Ban{
		Target:   target,
		Time:     time,
		Absolute: absolute,
	}

	if err := network.SetBan(ban); err != nil {
		logger.Errorf("failed to ban target: %s", err.Error())
	} else {
		logger.Infof("target %s was banned!", ban.Target)
	}
}

func (n *networkHandler) Unban(cmd *cobra.Command, args []string) {
	target, ok := getTargetIP(cmd, args)
	if !ok {
		return
	}

	if err := network.Unban(target); err != nil {
		logger.Errorf("failed to unban target: %s", err.Error())
	} else {
		logger.Infof("target %s was unbanned!", target)
	}
}

func (n *networkHandler) Blacklist(cmd *cobra.Command, args []string) {
	list, err := network.ListBanned()
	if err != nil {
		logger.Errorf("failed to get blacklist: %s", err.Error())
		return
	}

	list.Print()
}

var getTargetIP = func(cmd *cobra.Command, args []string) (string, bool) {
	target := ""
	if len(args) <= 0 {
		flag, _ := cmd.Flags().GetString("target")
		target = flag
	} else {
		target = args[0]
	}

	if target == "" {
		// If no target is provided, show the command help
		if err := cmd.Help(); err != nil {
			logger.Errorf("failed to show output for command %s: %v", cmd.Short, err.Error())
		}
		return "", false
	}

	if _, _, err := net.ParseCIDR(target); err != nil && net.ParseIP(target) == nil {
		logger.Errorf("invalid format. the target must be either a valid IP address (e.g., '192.168.1.1') or a valid subnet mask (e.g., '192.168.1.0/24)")
		return "", false
	}

	return strings.TrimSpace(target), true
}
