package rpc

import (
	"github.com/avila-r/bitclient/config"
)

type (
	Version string
	Method  string
	ID      string
	Params  []any
	Header  string
)

const (
	Version2 Version = "2.0"

	ContentTypeHeaderLabel   = "Content-Type"
	AuthorizationHeaderLabel = "Authorization"

	ContentTypeApplicationJson Header = "application/json"
)

var (
	Identifier ID = ID(config.Get().Main.Use)

	NoParams Params = Params{}

	Bearer = func(token string) string {
		return "Bearer " + token
	}
)
