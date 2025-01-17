package rpc

import (
	"encoding/json"
	"strings"

	"github.com/avila-r/bitclient/failure"
	"github.com/avila-r/bitclient/logger"
)

// Json is a type alias for a map with string keys and interface{} values,
// used to represent JSON objects in the response.
type Json map[string]interface{}

// Array is a type alias for a slice of Json objects, representing a list of JSON objects.
type Array []Json

// PrintResult prints the result of a response in a readable format.
// It first tries to unmarshal the result as a JSON object, then as an array,
// and if both attempts fail, it prints the raw result.
func (r *Response) PrintResult() {
	// Try to unmarshal the result into a Json object.
	json, err := r.UnmarshalResult()
	if err != nil {
		// If unmarshaling as Json fails, try unmarshaling as an Array.
		arr, err := r.UnmarshalArray()
		if err != nil {
			// If both fail, print the raw result.
			logger.Print(r.Result)
			return
		}
		// If unmarshaling as Array succeeds, print the array.
		logger.Print(arr.ToString())
		return
	}
	// If unmarshaling as Json succeeds, print the Json object.
	logger.Print(json.ToString())
}

// UnmarshalResult unmarshals the response's result into a Json object.
// Returns an error if the unmarshaling fails.
func (r *Response) UnmarshalResult() (*Json, error) {
	result := Json{}
	if err := json.Unmarshal(r.Result, &result); err != nil {
		logger.Debugf("Error processing result: %v", err)
		return nil, failure.Of("failed to process result: %v", err.Error())
	}

	return &result, nil
}

// UnmarshalArray unmarshals the response's result into an Array of Json objects.
// Returns an error if the unmarshaling fails.
func (r *Response) UnmarshalArray() (*Array, error) {
	result := Array{}
	if err := json.Unmarshal(r.Result, &result); err != nil {
		logger.Debugf("Error processing result: %v", err)
		return nil, failure.Of("failed to process result: %v", err.Error())
	}

	return &result, nil
}

// ToString serializes a Json object into a formatted string.
// Returns a string representation of the Json object.
func (j Json) ToString() string {
	data, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		logger.Debugf("Failed to serialize json as string: %v", err)
	}
	return string(data)
}

// Print prints the Json in a readable format.
func (j Json) Print() {
	// If unmarshaling as Json succeeds, print the Json object.
	logger.Print(j.ToString())
}

// ToString serializes an Array of Json objects into a formatted string.
// Returns a string representation of the Array.
func (a Array) ToString() string {
	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		logger.Debugf("Failed to serialize array as string: %v", err)
	}
	return string(data)
}

// JsonResult unmarshals a response into a Json object and handles errors if any.
// Returns a Json object or an error.
func JsonResult(r *Response, err error, warning ...string) (*Json, error) {
	if r != nil {
		// If the response is not nil, try unmarshaling the result into a Json object.
		return r.UnmarshalResult()
	}
	// If the response is nil, handle the error.
	return handle[Json](err, warning...)
}

// ArrayResult unmarshals a response into an Array of Json objects and handles errors if any.
// Returns an Array of Json objects or an error.
func ArrayResult(r *Response, err error, warning ...string) (*Array, error) {
	if r != nil {
		// If the response is not nil, try unmarshaling the result into an Array.
		return r.UnmarshalArray()
	}
	// If the response is nil, handle the error.
	return handle[Array](err, warning...)
}

// handle is a generic function to handle errors and extract messages from them.
// It checks the error message and returns a custom error with the extracted message.
func handle[T any](err error, warning ...string) (*T, error) {
	if !strings.HasPrefix(err.Error(), "map") {
		// If the error is not a map-related error, return the original error.
		return nil, err
	}

	// stringToMap converts the string error message into a map of key-value pairs.
	stringToMap := func(s string) map[string]string {
		trimmed := strings.TrimPrefix(strings.TrimSuffix(s, "]"), "map[")
		result := make(map[string]string)
		var currentKey string
		var currentValue strings.Builder
		inValue := false

		// Iterate through the string character by character to build the map.
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

	// Check if the error contains a "message" field.
	if message, exists := stringToMap(err.Error())["message"]; exists {
		var err error
		if len(warning) > 0 {
			err = failure.Of("%s (%s)", strings.ToLower(message), warning[0])
		} else {
			err = failure.Of("%v", strings.ToLower(message))
		}
		return nil, err
	}

	// If no message is found, return the original error.
	return nil, err
}
