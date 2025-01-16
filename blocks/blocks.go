package blocks

import (
	"encoding/json"
	"regexp"
	"strconv"

	"github.com/avila-r/bitclient/errs"
	"github.com/avila-r/bitclient/rpc"
)

// BlockInfoVerbosity defines the verbosity levels for retrieving block information.
// The verbosity levels determine the detail and format of the response returned
// by the "getblock" RPC call.
type BlockInfoVerbosity int

const (
	// VerbositySerializedHexData indicates that the response will be a serialized, hex-encoded string of the block data.
	VerbositySerializedHexData BlockInfoVerbosity = iota

	// VerbosityBasicBlockInfo indicates that the response will be a JSON object containing basic block information.
	VerbosityBasicBlockInfo

	// VerbosityDetailedBlockInfo indicates that the response will be a JSON object containing detailed block
	// information, including details for each transaction in the block.
	VerbosityDetailedBlockInfo

	// VerbosityFullBlockInfoWithPrevout indicates that the response will include full block details, as well as
	// information about the previous outpoints for all transactions in the block.
	VerbosityFullBlockInfoWithPrevout
)

// VerbosityFrom converts an integer to the corresponding BlockInfoVerbosity value.
//
// Parameters:
// - i (int): The verbosity level as an integer.
//
// Returns:
// - BlockInfoVerbosity: The corresponding verbosity level.
// - error: An error if the provided integer is outside the valid range (0–3).
//
// Example:
//
//	verbosity, err := VerbosityFrom(2)
//	if err != nil {
//	    // Handle error
//	}
//
// Notes:
//   - The valid range for verbosity levels is 0–3. Values outside this range
//     will result in an error.
func VerbosityFrom(i int) (BlockInfoVerbosity, error) {
	if i < int(VerbositySerializedHexData) || i > int(VerbosityFullBlockInfoWithPrevout) {
		return 0, errs.Of("invalid verbosity level (%d), valid range is 0-3", i)
	}
	return BlockInfoVerbosity(i), nil
}

// IsBlockHashInvalid validates a block hash string.
//
// Parameters:
// - blockhash (string): The block hash to validate.
//
// Returns:
// - bool: True if the blockhash is invalid, false otherwise.
//
// A valid blockhash:
// - Must be exactly 64 characters long.
// - Must consist only of hexadecimal characters (0-9, a-f, A-F).
//
// Example:
//
//	if IsBlockHashInvalid("00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09") {
//	    // Handle invalid block hash
//	}
func IsBlockHashInvalid(blockhash string) bool {
	return len(blockhash) != 64 || !regexp.MustCompile("^[0-9a-fA-F]{64}$").MatchString(blockhash)
}

// GetBlock retrieves information about a specific block by its hash and verbosity level.
//
// This function sends a JSON-RPC request to the Bitcoin client using the "getblock" procedure call.
// The returned data depends on the verbosity level provided:
//   - 0: Serialized, hex-encoded data for the block.
//   - 1: Basic JSON object with block information.
//   - 2: Detailed JSON object with block information and transaction details.
//
// Parameters:
//   - block (string or numeric, required): The block hash or height of the target block.
//     The function accepts either a block hash (64-character hex string) or a numeric block height.
//   - verbosity (int): The verbosity level for the response. Valid values are:
//     0: Serialized hex data.
//     1: Basic block information (default).
//     2: Detailed block and transaction information.
//     3: Full block info with detailed previous outpoints.
//
// Returns:
// - *rpc.Response: The JSON-RPC response containing the block data according to the verbosity level.
// - error: An error if the blockhash is invalid, the verbosity level is out of range, or the request fails.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient blocks get 00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09 --verbosity 2
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getblock "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09" 2
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblock", "params": ["{blockhash}", {verbosity}]}' \
//     -H 'content-type: text/plain;' {url}
//
// Notes:
// - The blockhash must be exactly 64 hexadecimal characters. Validation is enforced using `IsBlockHashInvalid`.
// - The verbosity level is validated using `VerbosityFrom` to ensure it is within the range 0–3.
// - Ensure the RPC client is properly configured and connected to the Bitcoin node before calling this function.
// - The node must be synchronized with the blockchain to provide accurate block information.
//
// Verbosity Levels:
// - VerbositySerializedHexData (0): Serialized, hex-encoded block data.
// - VerbosityBasicBlockInfo (1): JSON object with basic block data.
// - VerbosityDetailedBlockInfo (2): JSON object with block and transaction details.
// - VerbosityFullBlockInfoWithPrevout (3): Full block details, including previous outpoints.
func GetBlock(block string, verbosity int) (*rpc.Response, error) {
	if IsBlockHashInvalid(block) {
		height, _ := strconv.Atoi(block)
		hash, err := GetBlockHash(height)
		if err != nil {
			return nil, errs.Of("block must be a valid block hash or a numeric height")
		} else {
			block = hash
		}
	}

	_, err := VerbosityFrom(verbosity)
	if err != nil {
		return nil, err
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBlock,
		Params:  rpc.Params{block, verbosity},
	}

	return rpc.Client.Do(request)
}

// GetBlockFilter retrieves a BIP 157 compact block filter for a specified block.
//
// This function sends a JSON-RPC request to the Bitcoin client using the "getblockfilter" procedure call.
// The response contains the filter and filter header, which are hex-encoded. Compact block filters
// are used for light client applications and can be extended using the "filtertype" parameter.
//
// Parameters:
//   - block (string or numeric, required): The block hash or height of the target block.
//     The function accepts either a block hash (64-character hex string) or a numeric block height.
//
// Returns:
// - *rpc.Json: The JSON-RPC response containing the compact block filter and header.
// - error: An error if the blockhash is invalid or if the request fails.
//
// Notes:
//   - Ensure that compact block filters are enabled on the Bitcoin node by starting it with the
//     `-blockfilterindex=basic` or `-blockfilterindex` flag.
//   - A warning message will be included in the response if the block filter index is not active.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient blocks filter 00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getblockfilter "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09" "basic"
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockfilter", "params": ["{blockhash}", "basic"]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "getblockfilter",
//	  "params": ["00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09", "basic"]
//	}
//
// Filter Types:
// - "basic" (default): Returns the basic compact block filter.
//
// JSON Response:
//
//	{
//	  "filter": "hex",    // The hex-encoded filter data.
//	  "header": "hex"     // The hex-encoded filter header.
//	}
//
// Warning:
// If the compact block filter index is not active, the node will return an error. Ensure the index is enabled
// by configuring the node with the appropriate flags (`-blockfilterindex=basic` or `-blockfilterindex`).
//
// Example Response:
//
//	{
//	  "filter": "0123456789abcdef",
//	  "header": "fedcba9876543210"
//	}
func GetBlockFilter(block string) (*rpc.Json, error) {
	if IsBlockHashInvalid(block) {
		height, _ := strconv.Atoi(block)
		hash, err := GetBlockHash(height)
		if err != nil {
			return nil, errs.Of("block must be a valid block hash or a numeric height")
		} else {
			block = hash
		}
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBlockFilter,
		Params:  rpc.Params{block, "extended"},
	}

	result, err := rpc.Client.Do(request)
	warning := "maybe it's needed to activate compact block filter starting bitcoind with the -blockfilterindex=basic/-blockfilterindex flag"
	return rpc.JsonResult(result, err, warning)
}

// GetBlockHash retrieves the hash of a block at the specified height in the best-block chain.
//
// This function sends a JSON-RPC request to the Bitcoin client using the "getblockhash" procedure call.
// The response contains the block hash, encoded as a hexadecimal string.
//
// Parameters:
// - height (numeric): The height index of the block in the best-block chain.
//
// Returns:
// - *rpc.Response: The JSON-RPC response containing the block hash as a hex-encoded string.
// - error: An error if the request fails or if the provided height is invalid.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient blocks hash 1000
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getblockhash 1000
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockhash", "params": [1000]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "getblockhash",
//	  "params": [1000]
//	}
//
// JSON Response Example:
//
//	{
//	  "result": "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09",
//	  "error": null,
//	  "id": "curltest"
//	}
//
// Notes:
// - The `height` parameter must be a non-negative integer representing the block's position in the blockchain.
// - The genesis block is at height 0.
//
// Result:
// - `hex` (string): The block hash at the specified height, encoded as a hex string.
//
// Example Response:
//
//	{
//	  "result": "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09",
//	  "error": null,
//	  "id": "curltest"
//	}
func GetBlockHash(height int) (string, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBlockHash,
		Params:  rpc.Params{height},
	}

	response, err := rpc.Client.Do(request)
	if response == nil || err != nil {
		return "", err
	}

	hash := ""
	if err := json.Unmarshal(response.Result, &hash); err != nil {
		return string(response.Result), err
	}

	return hash, nil
}

// GetBlockHeader retrieves the header of a block specified by its block hash.
//
// This function sends a JSON-RPC request to the Bitcoin client using the "getblockheader" procedure call.
// The response contains either the hex-encoded block header data or a detailed JSON object with information about the block header,
// depending on the verbosity setting.
//
// Parameters:
//   - block (string or numeric, required): The block hash or height of the target block.
//     The function accepts either a block hash (64-character hex string) or a numeric block height.
//   - verbose (optional, bool): If true (default), returns a JSON object with detailed block header information.
//     If false, returns the block header as hex-encoded data.
//
// Returns:
// - *rpc.Response: The JSON-RPC response containing either the hex-encoded block header or a JSON object with block header details.
// - error: An error if the request fails or if the provided blockhash is invalid.
//
// Notes:
// - The `blockhash` parameter must be a valid 64-character hex string representing the block hash.
// - The verbosity parameter, if provided, determines the level of detail in the response:
//   - `true`: Returns a JSON object with block header details.
//   - `false`: Returns hex-encoded block header data (default is `true`).
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient blocks header 00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getblockheader "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09"
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockheader", "params": ["{blockhash}", true]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "getblockheader",
//	  "params": ["00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09", true]
//	}
//
// JSON Response Example (verbose=true):
//
//	{
//	  "hash": "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09",
//	  "confirmations": 100,
//	  "height": 123456,
//	  "version": 536870912,
//	  "versionHex": "20000000",
//	  "merkleroot": "abcd1234efgh5678",
//	  "time": 1622540800,
//	  "mediantime": 1622540600,
//	  "nonce": 123456789,
//	  "bits": "170f0f0f",
//	  "difficulty": 456.789,
//	  "chainwork": "0000000000000000000000000000000000000000000000000000000000000000",
//	  "nTx": 1000,
//	  "previousblockhash": "00000000abcd1234efgh5678",
//	  "nextblockhash": "00000000ijkl1234mnop5678"
//	}
//
// Result (for verbose=false):
//
//	{
//	  "hex": "0200000001abcd1234efgh5678..." // Serialized, hex-encoded block header data
//	}
func GetBlockHeader(block string, verbose ...bool) (*rpc.Response, error) {
	if IsBlockHashInvalid(block) {
		height, _ := strconv.Atoi(block)
		hash, err := GetBlockHash(height)
		if err != nil {
			return nil, errs.Of("block must be a valid block hash or a numeric height")
		} else {
			block = hash
		}
	}

	verbosity := true // Default value
	if len(verbose) > 0 {
		verbosity = verbose[0]
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBlockHeader,
		Params:  rpc.Params{block, verbosity},
	}

	return rpc.Client.Do(request)
}

// GetBlockStats retrieves statistical data for a given block specified by its hash or height.
//
// This function sends a JSON-RPC request to the Bitcoin client using the "getblockstats" procedure call.
// It computes various per-block statistics, with amounts in satoshis, for the block at the given height or block hash.
//
// Parameters:
//   - block (string or numeric, required): The block hash or height of the target block.
//     The function accepts either a block hash (64-character hex string) or a numeric block height.
//   - stats (optional, array of strings): A list of specific statistics to retrieve.
//     If no stats are provided, the function will return all available statistics. Example values:
//   - "height", "time", "avgfee", "minfeerate", "swtotal_size", etc.
//
// Returns:
// - *rpc.Json: A JSON object containing the requested block statistics.
// - error: An error if the request fails or if the provided block identifier is invalid.
//
// Notes:
// - The `hash_or_height` parameter can be either the block hash (64-character hex string) or a numeric block height.
// - The `stats` parameter is optional. If provided, it specifies which statistics to return. If not provided, all statistics will be returned.
// - The statistics are in satoshis, except for values like "height" and "time" which are in numeric format.
//
// Example Usage:
//
//   - Using Bitclient:
//     $ bitclient blocks stats 00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09 --only ["minfeerate", "avgfeerate"]
//
//   - Using the Bitcoin CLI:
//     $ bitcoin-cli getblockstats 1000 '["minfeerate", "avgfeerate"]'
//
//   - Using cURL:
//     $ curl --user {username} --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getblockstats", "params": ["{hash_or_height}", ["minfeerate", "avgfeerate"]]}' \
//     -H 'content-type: text/plain;' {url}
//
// RPC Request Example:
//
//	{
//	  "jsonrpc": "1.0",
//	  "id": "curltest",
//	  "method": "getblockstats",
//	  "params": ["00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09", ["minfeerate", "avgfeerate"]]
//	}
//
// JSON Response Example:
//
//	{
//	  "avgfee": 1000,
//	  "avgfeerate": 100,
//	  "avgtxsize": 250,
//	  "blockhash": "00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09",
//	  "feerate_percentiles": [10, 20, 30, 40, 50],
//	  "height": 1000,
//	  "ins": 500,
//	  "maxfee": 2000,
//	  "maxfeerate": 200,
//	  "maxtxsize": 300,
//	  "medianfee": 1500,
//	  "mediantime": 1622540600,
//	  "minfee": 500,
//	  "minfeerate": 50,
//	  "mintxsize": 200,
//	  "outs": 1000,
//	  "subsidy": 125000000,
//	  "swtotal_size": 1000,
//	  "swtotal_weight": 4000,
//	  "swtxs": 200,
//	  "time": 1622540800,
//	  "total_out": 10000000000,
//	  "total_size": 100000,
//	  "total_weight": 40000,
//	  "totalfee": 1000000,
//	  "txs": 1000,
//	  "utxo_increase": 50,
//	  "utxo_size_inc": 1000
//	}
func GetBlockStats(block string, stats ...string) (*rpc.Json, error) {
	if IsBlockHashInvalid(block) {
		height, _ := strconv.Atoi(block)
		hash, err := GetBlockHash(height)
		if err != nil {
			return nil, errs.Of("block must be a valid block hash or a numeric height")
		} else {
			block = hash
		}
	}

	params := rpc.Params{block}
	if len(stats) > 0 {
		params = append(params, stats)
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBlockStats,
		Params:  params,
	}

	return rpc.JsonResult(rpc.Client.Do(request))
}
