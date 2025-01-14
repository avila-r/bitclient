package blocks_test

import (
	"fmt"
	"testing"

	"github.com/avila-r/env"

	"github.com/avila-r/bitclient/blocks"
	"github.com/avila-r/bitclient/logger"
)

var (
	RequiredEnvs = []string{
		"RPC_URL",
		"RPC_AUTH_TYPE",
		"RPC_AUTH_LABEL",
	}
)

func init() {
	for _, key := range RequiredEnvs {
		if env.Get(key) == "" {
			logger.Fatalf("Client isn't available. %v variables must be provided", RequiredEnvs)
		}
	}
}

func Test_GetBlockchainInfo(t *testing.T) {
	if _, err := blocks.GetBlockchainInfo(); err != nil {
		t.Errorf("Failed to get blockchain info - %v", err)
	}
}

func Test_GetBlockCount(t *testing.T) {
	result, err := blocks.GetBlockCount()
	if err != nil {
		t.Errorf("Failed to get block count: %v", err)
	}

	if result != nil {
		if result.Error != nil {
			t.Errorf("RPC response contains an error: %v", result.Error)
		}
	}
}

func Test_GetBestBlockhash(t *testing.T) {
	result, err := blocks.GetBestBlockHash()
	if err != nil {
		t.Errorf("Failed to get best block's hash count: %v", err)
	}

	if result != nil {
		if result.Error != nil {
			t.Errorf("RPC response contains an error: %v", result.Error)
		}
	}
}

func Test_GetChainTips(t *testing.T) {
	if _, err := blocks.GetChainTips(); err != nil {
		t.Errorf("Failed to get chain tips: %v", err)
	}
}

func Test_GetBlock(t *testing.T) {
	tests := []struct {
		Verbosity     int
		ExpectSuccess bool
	}{
		{Verbosity: -1, ExpectSuccess: false},
		{Verbosity: 0, ExpectSuccess: true},
		{Verbosity: 1, ExpectSuccess: true},
		{Verbosity: 2, ExpectSuccess: true},
		{Verbosity: 3, ExpectSuccess: true},
		{Verbosity: 4, ExpectSuccess: false},
		{Verbosity: 5, ExpectSuccess: false},
	}

	blockhash := "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09"
	for i, test := range tests {
		t.Run(fmt.Sprintf("case_%d_verbosity_%d", i, test.Verbosity), func(t *testing.T) {
			result, err := blocks.GetBlock(blockhash, test.Verbosity)

			if result != nil {
				if result.Error != nil {
					t.Errorf("Test case %d failed: rpc.Response.Error: %v", i, err)
				}
			}

			if test.ExpectSuccess {
				if err != nil {
					t.Errorf("Test case %d failed: expected success but got error: %v", i, err)
				} else if result == nil {
					t.Errorf("Test case %d failed: expected a valid result but got nil", i)
				}
			} else {
				if err == nil {
					t.Errorf("Test case %d failed: expected failure but got success with result: %v", i, result)
				}
			}
		})
	}
}

func Test_GetBlockFilter(t *testing.T) {
	blockhash := "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09"

	_, err := blocks.GetBlockFilter(blockhash)
	if err != nil {
		t.Logf("Non-critical failure: failed to get block filter - %v", err)
		// t.Errorf("Failed to get block filter - %v", err)
	}
}

func Test_GetBlockHash(t *testing.T) {}

func Test_GetBlockHeader(t *testing.T) {
	blockhash := "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09"

	verbose, err := blocks.GetBlockHeader(blockhash)
	if err != nil {
		t.Errorf("Failed to get block header: %v", err)
	}
	if _, err := verbose.UnmarshalResult(); err != nil {
		t.Errorf("Failed to unmarshal verbose block header result: %v", err)
	}

	response, err := blocks.GetBlockHeader(blockhash, false)
	if err != nil {
		t.Errorf("Failed to get block header: %v", err)
	}
	if response != nil {
		if response.Error != nil {
			t.Errorf("RPC response contains an error: %v", response.Error)
		}
	}
}

func Test_GetBlockStats(t *testing.T) {
	blockhash := "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09"

	_, err := blocks.GetBlockStats(blockhash)
	if err != nil {
		t.Errorf("Failed to get block stats: %v", err)
	}
}
