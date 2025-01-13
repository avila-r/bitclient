package logger

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	base = "bitclient/"
)

func caller() string {
	_, fullPath, line, ok := runtime.Caller(2) // 2 levels up the stack
	if !ok {
		// Fallback if caller
		// info isn't available
		return "[unknown:0]"
	}

	// Truncate path to show project-relative path
	index := strings.Index(fullPath, base)
	path := fullPath
	if index != -1 {
		path = fullPath[index:] // Extract the path starting from the project base
	}

	return fmt.Sprintf("[%s:%d]", path, line)
}
