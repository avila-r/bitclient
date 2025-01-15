package rpc_test

import (
	"fmt"
	"testing"

	"github.com/avila-r/env"

	"github.com/avila-r/bitclient/logger"
	"github.com/avila-r/bitclient/rpc"
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

func Test_GetMemoryInfo(t *testing.T) {
	if _, err := rpc.GetMemoryInfo(); err != nil {
		t.Errorf("Failed to get memory info: %v", err)
	}
}

func Test_GetInfo(t *testing.T) {
	if _, err := rpc.GetInfo(); err != nil {
		t.Errorf("Failed to get memory info: %v", err)
	}
}

func Test_GetHelp(t *testing.T) {
	cases := []struct{ Command []string }{
		{Command: nil},
		{Command: []string{"getblockchaininfo"}},
		{Command: []string{"getnetworkinfo"}},
		{Command: []string{"getwalletinfo"}},
		{Command: []string{"help"}},
	}

	for i, test := range cases {
		name := fmt.Sprintf("case %v", i)
		t.Run(name, func(t *testing.T) {
			if _, err := rpc.Help(test.Command...); err != nil {
				t.Errorf("Failed to get command help: %v", err)
			}
		})
	}
}

func Test_Logging(t *testing.T) {
	if _, err := rpc.GetLogging(); err != nil {
		t.Errorf("Failed to manage rpc logging: %v", err)
	}
}
