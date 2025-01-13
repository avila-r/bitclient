package blocks

import (
	"regexp"
	"strconv"

	"github.com/avila-r/bitclient/errs"
	"github.com/avila-r/bitclient/rpc"
)

type BlockInfoVerbosity int

const (
	VerbositySerializedHexData BlockInfoVerbosity = iota
	VerbosityBasicBlockInfo
	VerbosityDetailedBlockInfo
	VerbosityFullBlockInfoWithPrevout
)

func VerbosityFrom(i int) (BlockInfoVerbosity, error) {
	if i < int(VerbositySerializedHexData) || i > int(VerbosityFullBlockInfoWithPrevout) {
		return 0, errs.Of("invalid verbosity level (%d), valid range is 0-3", i)
	}
	return BlockInfoVerbosity(i), nil
}

func IsBlockHashInvalid(blockhash string) bool {
	return len(blockhash) != 64 || !regexp.MustCompile("^[0-9a-fA-F]{64}$").MatchString(blockhash)
}

func GetBlock(blockhash string, verbosity int) (*rpc.Response, error) {
	if IsBlockHashInvalid(blockhash) {
		return nil, errs.Of("blockhash provided isn't in valid format")
	}

	_, err := VerbosityFrom(verbosity)
	if err != nil {
		return nil, err
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBlock,
		Params:  rpc.Params{blockhash, verbosity},
	}

	return rpc.Client.Do(request)
}

func GetBlockFilter(blockhash string) (*rpc.Json, error) {
	if IsBlockHashInvalid(blockhash) {
		return nil, errs.Of("blockhash provided isn't in valid format")
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBlockFilter,
		Params:  rpc.Params{blockhash, "extended"},
	}

	result, err := rpc.Client.Do(request)
	warning := "maybe it's needed to activate compact block filter starting bitcoind with the -blockfilterindex=basic/-blockfilterindex flag"
	return rpc.JsonResult(result, err, warning)
}

func GetBlockHash(height string) (*rpc.Response, error) {
	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBlockHash,
		Params:  rpc.Params{height},
	}

	return rpc.Client.Do(request)
}

func GetBlockHeader(blockhash string, verbose ...bool) (*rpc.Response, error) {
	verbosity := true // Default value
	if len(verbose) > 0 {
		verbosity = verbose[0]
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBlockHeader,
		Params:  rpc.Params{blockhash, verbosity},
	}

	return rpc.Client.Do(request)
}

func GetBlockStats(hash_or_height string, stats ...string) (*rpc.Json, error) {
	if invalid := IsBlockHashInvalid(hash_or_height); invalid {
		if _, err := strconv.Atoi(hash_or_height); err != nil {
			return nil, errs.Of("hash_or_height must be a valid block hash or a numeric height")
		}
	}

	params := rpc.Params{hash_or_height}
	for _, stat := range stats {
		params = append(params, stat)
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  MethodGetBlockStats,
		Params:  params,
	}

	return rpc.JsonResult(rpc.Client.Do(request))
}
