package logger

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	base = "bitclient/"
)

// caller returns a string with the relative file path and line number
// of the function that invoked it, useful for logging or error tracking.
// It extracts the information two levels up the call stack.
//
// The function performs the following steps:
//  1. It uses `runtime.Caller(2)` to retrieve the caller information,
//     where `2` indicates that the function goes two levels up the call stack.
//  2. If successful, it extracts the full file path and line number.
//  3. The path is truncated to show the project-relative path (starting from the `bitclient/` directory).
//  4. The output is formatted as `[project-relative-path:line-number]`.
//
// If the caller information is unavailable, the function returns "[unknown:0]".
func caller() string {
	// Retrieve the file and line number of the caller two levels up the stack
	_, fullPath, line, ok := runtime.Caller(2)
	if !ok {
		// If caller info is not available, return a fallback value
		return "[unknown:0]"
	}

	// Truncate the full path to show only the part relative to the project base
	index := strings.Index(fullPath, base)
	path := fullPath
	if index != -1 {
		path = fullPath[index:] // Extract the path starting from the project base
	}

	// Return the formatted string with file path and line number
	return fmt.Sprintf("[%s:%d]", path, line)
}
