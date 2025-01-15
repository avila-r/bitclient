package blocks

import (
	"math/big"

	"github.com/avila-r/bitclient/rpc"
)

// GetBestBlockHash retrieves the hash of the best (tip) block in the most-work fully-validated chain.
//
// This function sends a JSON-RPC request to the Bitcoin client using the "getbestblockhash" procedure call.
// The response contains the hash of the best block in the blockchain, encoded as a hexadecimal string.
//
// Returns:
// - *rpc.Response: The JSON-RPC response containing the block hash as a hex-encoded string.
// - error: An error if the request fails or if there is an issue with the response.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient blockchain bestblockhash
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getbestblockhash
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getbestblockhash", "params": []}' \
//     -H 'content-type: text/plain;' {url}
//
// Note:
// Ensure the RPC client is properly configured and connected to the Bitcoin node
// before calling this function. The node must have synchronized with the blockchain
// to return a valid best block hash.
func GetBestBlockHash() (*rpc.Response, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBestBlockHash,
		Params:  rpc.NoParams,
	}

	return rpc.Client.Do(request)
}

// GetBlockchainInfo retrieves detailed state information regarding blockchain processing.
//
// This function sends a JSON-RPC request to the Bitcoin client using the "getblockchaininfo" procedure call.
// The response is an object containing various blockchain processing details, such as chain, blocks, headers,
// difficulty, median time, verification progress, and more.
//
// Returns:
// - *rpc.Json: The JSON-RPC response containing blockchain state information in a structured format.
// - error: An error if the request fails or if there is an issue with the response.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient blockchain info
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getblockchaininfo
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockchaininfo", "params": []}' \
//     -H 'content-type: text/plain;' {url}
//
// Note:
// Ensure the RPC client is properly configured and connected to the Bitcoin node before calling this function.
// The node must be running and synchronized to return accurate blockchain state information.
func GetBlockchainInfo() (*rpc.Json, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBlockchainInfo,
		Params:  rpc.NoParams,
	}

	return rpc.JsonResult(rpc.Client.Do(request))
}

// GetBlockCount retrieves the height of the most-work fully-validated chain.
//
// This function sends a JSON-RPC request to the Bitcoin client using the "getblockcount" procedure call.
// The response provides the current block count, where the genesis block has a height of 0.
//
// Returns:
// - *rpc.Response: The JSON-RPC response containing the block count as a numeric value.
// - error: An error if the request fails or if there is an issue with the response.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient blockchain blockcount
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getblockcount
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockcount", "params": []}' \
//     -H 'content-type: text/plain;' {url}
//
// Note:
// Ensure the RPC client is properly configured and connected to the Bitcoin node before calling this function.
// The node must be synchronized to the blockchain for the block count to be accurate.
func GetBlockCount() (*rpc.Response, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBlockCount,
		Params:  rpc.NoParams,
	}

	return rpc.Client.Do(request)
}

// GetChainTips retrieves information about all known tips in the block tree, including the main chain
// and orphaned branches.
//
// This function sends a JSON-RPC request to the Bitcoin client using the "getchaintips" procedure call.
// The response provides details for each chain tip, including its height, hash, branch length, and status.
// The status can indicate whether the chain tip is part of the active main chain or an orphaned branch.
//
// Returns:
//   - *rpc.Array: The JSON-RPC response containing an array of chain tip information, with each element represented
//     as a JSON object containing:
//   - "height" (numeric): Height of the chain tip.
//   - "hash" (string): Block hash of the tip, hex-encoded.
//   - "branchlen" (numeric): Zero for the main chain, otherwise the length of the branch connecting the tip to the main chain.
//   - "status" (string): Status of the chain tip, with possible values:
//   - "invalid": The branch contains at least one invalid block.
//   - "headers-only": Headers are valid, but not all blocks are available for this branch.
//   - "valid-headers": All blocks are available for this branch, but they were never fully validated.
//   - "valid-fork": The branch is fully validated but is not part of the active chain.
//   - "active": This is the tip of the active main chain, which is certainly valid.
//   - error: An error if the request fails or if there is an issue with the response.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient blockchain tips
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getchaintips
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getchaintips", "params": []}' \
//     -H 'content-type: text/plain;' {url}
//
// Note:
// Ensure the RPC client is properly configured and connected to the Bitcoin node before calling this function.
// The node must be synchronized to provide accurate information about chain tips.
func GetChainTips() (*rpc.Array, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetChainTips,
		Params:  rpc.NoParams,
	}

	return rpc.ArrayResult(rpc.Client.Do(request))
}

// GetChainTxStats retrieves the transaction statistics for a given chain of blocks.
//
// This function sends a JSON-RPC request to the Bitcoin client using the "getchaintxstats" procedure call.
// The response provides transaction statistics for the chain, including the number of transactions, the time span, and other metrics.
//
// Returns:
// - *rpc.Json: The JSON-RPC response containing transaction statistics for the chain in a structured format.
// - error: An error if the request fails or if there is an issue with the response.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient blockchain txstats
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getchaintxstats
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getchaintxstats", "params": [nblocks, blockhash]}' \
//     -H 'content-type: text/plain;' {url}
//
// Note:
// Ensure the RPC client is properly configured and connected to the Bitcoin node before calling this function.
// The node must be synchronized for accurate transaction statistics.
func GetChainTxStats(nblocks int, blockhash ...string) (*rpc.Json, error) {
	params := rpc.Params{}
	if nblocks > 0 {
		params = append(params, nblocks)
	}
	if len(blockhash) > 0 {
		params = append(params, blockhash)
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetChainTxStats,
		Params:  params,
	}

	return rpc.JsonResult(rpc.Client.Do(request))
}

// GetDifficulty retrieves the current mining difficulty of the Bitcoin network.
//
// This function sends a JSON-RPC request to the Bitcoin client using the "getdifficulty" procedure call.
// The response provides the current difficulty, which is a measure of how difficult it is to find a new block.
//
// Returns:
// - *big.Float: The current mining difficulty as a floating-point number.
// - error: An error if the request fails or if there is an issue with the response.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient blockchain difficulty
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getdifficulty
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getdifficulty", "params": []}' \
//     -H 'content-type: text/plain;' {url}
//
// Note:
// Ensure the RPC client is properly configured and connected to the Bitcoin node before calling this function.
// The node must be synchronized to return an accurate difficulty value.
func GetDifficulty() (*big.Float, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetDifficulty,
		Params:  rpc.NoParams,
	}

	response, err := rpc.Client.Do(request)
	if response == nil || err != nil {
		return nil, err
	}

	var difficulty big.Float
	difficulty.SetString(string(response.Result))
	return &difficulty, nil
}
