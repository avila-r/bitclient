package errs

import (
	"errors"
	"fmt"
	"strings"

	"github.com/avila-r/bitclient/logger"
)

// Of creates an error, optionally formatting the message with arguments.
func Of(msg string, v ...any) error {
	if len(v) == 0 && strings.Contains(msg, "%") {
		logger.Debugf("invalid format string or missing arguments: %v", msg)
		return errors.New("malformed error")
	}

	if len(v) > 0 {
		return fmt.Errorf(msg, v...)
	}

	return errors.New(msg)
}
