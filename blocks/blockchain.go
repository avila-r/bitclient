package blocks

import (
	"github.com/avila-r/bitclient/rpc"
)

func GetBlockchainInfo() (*rpc.Json, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBlockchainInfo,
		Params:  rpc.NoParams,
	}

	return rpc.JsonResult(rpc.Client.Do(request))
}

func GetBlockCount() (*rpc.Response, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBlockCount,
		Params:  rpc.NoParams,
	}

	return rpc.Client.Do(request)
}

func GetBestBlockHash() (*rpc.Response, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBestBlockHash,
		Params:  rpc.NoParams,
	}

	return rpc.Client.Do(request)
}
