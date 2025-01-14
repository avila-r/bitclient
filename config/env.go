package config

import "github.com/avila-r/env"

// init is called automatically when the package is imported.
// It loads the environment variables from the specified RootPath.
func init() {
	env.Load(RootPath) // Load environment variables from the RootPath directory
}
