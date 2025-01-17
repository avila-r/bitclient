package cmd

import (
	"github.com/spf13/cobra"

	"github.com/avila-r/bitclient/config"
	"github.com/avila-r/bitclient/handler"
)

// bitclient network
var (
	Network = &cobra.Command{
		Use:   config.Get().Commands.Network.Use,
		Short: config.Get().Commands.Network.ShortDescription,
		Long:  config.Get().Commands.Network.LongDescription,
	}
)

var (
	// bitclient network activate
	NetworkActivate = &cobra.Command{
		Use:   config.Get().Commands.Network.Activate.Use,
		Short: config.Get().Commands.Network.Activate.ShortDescription,
		Long:  config.Get().Commands.Network.Activate.LongDescription,
		Run:   handler.Network.Activate,
	}

	// bitclient network deactivate
	NetworkDeactivate = &cobra.Command{
		Use:   config.Get().Commands.Network.Deactivate.Use,
		Short: config.Get().Commands.Network.Deactivate.ShortDescription,
		Long:  config.Get().Commands.Network.Deactivate.LongDescription,
		Run:   handler.Network.Deactivate,
	}

	// bitclient network connections
	NetworkConnections = &cobra.Command{
		Use:   config.Get().Commands.Network.Connections.Use,
		Short: config.Get().Commands.Network.Connections.ShortDescription,
		Long:  config.Get().Commands.Network.Connections.LongDescription,
		Run:   handler.Network.Connections,
	}

	// bitclient network traffic
	NetworkTraffic = &cobra.Command{
		Use:   config.Get().Commands.Network.Traffic.Use,
		Short: config.Get().Commands.Network.Traffic.ShortDescription,
		Long:  config.Get().Commands.Network.Traffic.LongDescription,
		Run:   handler.Network.Traffic,
	}

	// bitclient network info
	NetworkInfo = &cobra.Command{
		Use:   config.Get().Commands.Network.Info.Use,
		Short: config.Get().Commands.Network.Info.ShortDescription,
		Long:  config.Get().Commands.Network.Info.LongDescription,
		Run:   handler.Network.Info,
	}

	// bitclient network peers
	NetworkPeers = &cobra.Command{
		Use:   config.Get().Commands.Network.Peers.Use,
		Short: config.Get().Commands.Network.Peers.ShortDescription,
		Long:  config.Get().Commands.Network.Peers.LongDescription,
		Run:   handler.Network.Peers,
	}

	// bitclient network ban
	NetworkBan = &cobra.Command{
		Use:   config.Get().Commands.Network.Ban.Use,
		Short: config.Get().Commands.Network.Ban.ShortDescription,
		Long:  config.Get().Commands.Network.Ban.LongDescription,
		Run:   handler.Network.Ban,
	}

	// bitclient network unban
	NetworkUnban = &cobra.Command{
		Use:   config.Get().Commands.Network.Unban.Use,
		Short: config.Get().Commands.Network.Unban.ShortDescription,
		Long:  config.Get().Commands.Network.Unban.LongDescription,
		Run:   handler.Network.Unban,
	}

	// bitclient network blacklist
	NetworkBlacklist = &cobra.Command{
		Use:   config.Get().Commands.Network.Blacklist.Use,
		Short: config.Get().Commands.Network.Blacklist.ShortDescription,
		Long:  config.Get().Commands.Network.Blacklist.LongDescription,
		Run:   handler.Network.Blacklist,
	}
)

func init() {
	Root.AddCommand(Network) // bitclient network
	// Flags
	{
	}

	// Subcommands
	Network.AddCommand(
		NetworkActivate,
		NetworkDeactivate,
		NetworkConnections,
		NetworkTraffic,
		NetworkInfo,
		NetworkPeers,
		NetworkBan,
		NetworkUnban,
		NetworkBlacklist,
	)
}
