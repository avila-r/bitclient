package blocks

import "github.com/avila-r/bitclient/rpc"

const (
	MethodGetBestBlockHash  rpc.Method = "getbestblockhash"
	MethodGetBlock          rpc.Method = "getblock"
	MethodGetBlockchainInfo rpc.Method = "getblockchaininfo"
	MethodGetBlockCount     rpc.Method = "getblockcount"
	MethodGetBlockFilter    rpc.Method = "getblockfilter"
	MethodGetBlockHash      rpc.Method = "getblockhash"
	MethodGetBlockHeader    rpc.Method = "getblockheader"
	MethodGetBlockStats     rpc.Method = "getblockstats"
)
