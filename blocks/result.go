package blocks

import (
	"strings"

	"github.com/avila-r/bitclient/errs"
	"github.com/avila-r/bitclient/rpc"
)

func Result(r *rpc.Json, err error) (*rpc.Json, error) {
	if r != nil {
		return r, err
	}

	if !strings.HasPrefix(err.Error(), "map") {
		return nil, err
	}

	if message, exists := stringToMap(err.Error())["message"]; exists {
		return nil, errs.Of("%v", strings.ToLower(message))
	}

	return nil, err
}

func stringToMap(s string) map[string]string {
	trimmed := strings.TrimPrefix(strings.TrimSuffix(s, "]"), "map[")

	result := make(map[string]string)
	var currentKey string
	var currentValue strings.Builder
	inValue := false

	// Iterate through the string character by character
	for i := 0; i < len(trimmed); i++ {
		char := trimmed[i]

		if char == ':' && !inValue {
			currentKey = strings.TrimSpace(currentValue.String())
			currentValue.Reset()
			inValue = true
			continue
		}

		if char == ' ' && !inValue {
			if currentKey != "" && currentValue.Len() > 0 {
				result[currentKey] = currentValue.String()
				currentValue.Reset()
			}
			inValue = false
			continue
		}

		if inValue && char == ' ' {
			rest := trimmed[i+1:]
			if strings.Contains(rest, "code:") || strings.Contains(rest, "message:") {
				result[currentKey] = currentValue.String()
				currentValue.Reset()
				inValue = false
				continue
			}
		}

		currentValue.WriteByte(char)
	}

	if currentKey != "" && currentValue.Len() > 0 {
		result[currentKey] = currentValue.String()
	}

	return result
}
