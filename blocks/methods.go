package blocks

import "github.com/avila-r/bitclient/rpc"

const (
	MethodGetBestBlockHash  rpc.Method = "getbestblockhash"  // Method to get the best block hash
	MethodGetBlock          rpc.Method = "getblock"          // Method to get block data by hash
	MethodGetBlockchainInfo rpc.Method = "getblockchaininfo" // Method to get blockchain info
	MethodGetBlockCount     rpc.Method = "getblockcount"     // Method to get the block count
	MethodGetBlockFilter    rpc.Method = "getblockfilter"    // Method to get a block filter by block hash
	MethodGetBlockHash      rpc.Method = "getblockhash"      // Method to get block hash by height
	MethodGetBlockHeader    rpc.Method = "getblockheader"    // Method to get block header by hash
	MethodGetBlockStats     rpc.Method = "getblockstats"     // Method to get block stats by hash or height
	MethodGetChainTips      rpc.Method = "getchaintips"      // Method to get chain tips
	MethodGetChainTxStats   rpc.Method = "getchaintxstats"   // Method to get chain transaction stats
	MethodGetDifficulty     rpc.Method = "getdifficulty"     // Method to get the current mining difficulty
)
