package network_test

import (
	"testing"

	"github.com/avila-r/env"

	"github.com/avila-r/bitclient/logger"
	"github.com/avila-r/bitclient/network"
)

func init() {
	var envs = []string{
		"RPC_URL",
		"RPC_AUTH_TYPE",
		"RPC_AUTH_LABEL",
	}

	for _, key := range envs {
		if env.Get(key) == "" {
			logger.Fatalf("Client isn't available. %v variables must be provided", envs)
		}
	}
}

func Test_ConnectToNode(t *testing.T) {
	// TODO
}

func Test_AddNode(t *testing.T) {
	// TODO
}

func Test_RemoveNode(t *testing.T) {
	// TODO
}

func Test_ClearBanned(t *testing.T) {
	// TODO
}

func Test_DisconnectNode(t *testing.T) {
	// TODO
}

func Test_InspectAddedNodes(t *testing.T) {
	// TODO
}

func Test_GetConnectionCount(t *testing.T) {
	// TODO
}

func Test_InspectTraffic(t *testing.T) {
	// TODO
}

func Test_GetNetworkInfo(t *testing.T) {
	if _, err := network.GetNetworkInfo(); err != nil {
		t.Errorf("Failed to get network info: %v", err)
	}
}

func Test_FindAddresses(t *testing.T) {
	// TODO
}

func Test_GetPeers(t *testing.T) {
	// TODO
}

func Test_ListBanned(t *testing.T) {
	// TODO
}

func Test_Ping(t *testing.T) {
	if err := network.Ping(); err != nil {
		t.Errorf("Failed to ping rpc server: %v", err)
	}
}

func Test_Health(t *testing.T) {
	if ok := network.Health(); !ok {
		t.Errorf("RPC server isn't uptime")
	}
}

func Test_SetBan(t *testing.T) {
	// TODO
}

func Test_Unbah(t *testing.T) {
	// TODO
}

func Test_SetNetworkActive(t *testing.T) {
	// TODO
}
