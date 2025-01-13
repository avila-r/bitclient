package blocks

import "github.com/avila-r/bitclient/rpc"

const (
	MethodGetBlock          rpc.Method = "getblock"
	MethodGetBlockchainInfo rpc.Method = "getblockchaininfo"
	MethodGetBlockCount     rpc.Method = "getblockcount"
	MethodGetBestBlockHash  rpc.Method = "getbestblockhash"
)
