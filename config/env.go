package config

import "github.com/avila-r/env"

func init() {
	env.Load(RootPath)
}
