package blocks

import (
	"github.com/avila-r/bitclient/rpc"
)

func GetBlockchainInfo() (*rpc.Json, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  rpc.MethodGetBlockchainInfo,
		Params:  rpc.NoParams,
	}

	return Result(rpc.Client.Do(request))
}
