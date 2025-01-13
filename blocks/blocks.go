package blocks

import (
	"regexp"

	"github.com/avila-r/bitclient/errs"
	"github.com/avila-r/bitclient/rpc"
)

type ResponseVerbosity int

const (
	VerbositySerializedHexData ResponseVerbosity = iota
	VerbosityBasicBlockInfo
	VerbosityDetailedBlockInfo
	VerbosityFullBlockInfoWithPrevout
)

func VerbosityFrom(i int) (ResponseVerbosity, error) {
	if i < int(VerbositySerializedHexData) || i > int(VerbosityFullBlockInfoWithPrevout) {
		return 0, errs.Of("invalid verbosity level (%d), valid range is 0-3", i)
	}
	return ResponseVerbosity(i), nil
}

func GetBlock(blockhash string, verbosity int) (*rpc.Response, error) {
	if len(blockhash) != 64 || !regexp.MustCompile("^[0-9a-fA-F]{64}$").MatchString(blockhash) {
		return nil, errs.Of("blockhash provided isn't in valid format")
	}

	_, err := VerbosityFrom(verbosity)
	if err != nil {
		return nil, err
	}

	request := rpc.Request{
		ID:      rpc.Identifier,
		Version: rpc.Version2,
		Method:  rpc.MethodGetBlock,
		Params:  rpc.Params{blockhash, verbosity},
	}

	return rpc.Client.Do(request)
}
