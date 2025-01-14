package rpc

import (
	"encoding/json"
	"strings"

	"github.com/avila-r/bitclient/errs"
	"github.com/avila-r/bitclient/logger"
)

type Json map[string]interface{}

type Array []Json

func (r *Response) PrintResult() {
	json, err := r.UnmarshalResult()
	if err != nil {
		arr, err := r.UnmarshalArray()
		if err != nil {
			// Print raw result
			logger.Print(r.Result)
			return
		}
		logger.Print(arr.ToString())
		return
	}
	logger.Print(json.ToString())
}

func (r *Response) UnmarshalResult() (*Json, error) {
	result := Json{}
	if err := json.Unmarshal(r.Result, &result); err != nil {
		logger.Debugf("Error processing result: %v", err)
		return nil, errs.Of("failed to process result: %v", err.Error())
	}

	return &result, nil
}

func (r *Response) UnmarshalArray() (*Array, error) {
	result := Array{}
	if err := json.Unmarshal(r.Result, &result); err != nil {
		logger.Debugf("Error processing result: %v", err)
		return nil, errs.Of("failed to process result: %v", err.Error())
	}

	return &result, nil
}

func (j Json) ToString() string {
	data, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		logger.Debugf("Failed to serialize json as string: %v", err)
	}

	return string(data)
}

func (a Array) ToString() string {
	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		logger.Debugf("Failed to serialize array as string: %v", err)
	}

	return string(data)
}

func JsonResult(r *Response, err error, warning ...string) (*Json, error) {
	if r != nil {
		return r.UnmarshalResult()
	}

	return handle[Json](r, err, warning...)
}

func ArrayResult(r *Response, err error, warning ...string) (*Array, error) {
	if r != nil {
		return r.UnmarshalArray()
	}

	return handle[Array](r, err, warning...)
}

func handle[T any](r *Response, err error, warning ...string) (*T, error) {
	if !strings.HasPrefix(err.Error(), "map") {
		return nil, err
	}

	stringToMap := func(s string) map[string]string {
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

	if message, exists := stringToMap(err.Error())["message"]; exists {
		var err error
		if len(warning) > 0 {
			err = errs.Of("%s (%s)", strings.ToLower(message), warning[0])
		} else {
			err = errs.Of("%v", strings.ToLower(message))
		}

		return nil, err
	}

	return nil, err
}
