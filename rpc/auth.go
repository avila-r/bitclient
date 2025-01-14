package rpc

import (
	"net/http"
	"strings"

	"github.com/avila-r/bitclient/errs"
)

// Authentication represents the authentication details used for HTTP requests.
type Authentication struct {
	Type  AuthenticationType // The type of authentication (API key or credentials)
	Label string             // The authentication label (e.g., API key or username:password)
}

// AuthenticationType defines the type of authentication used.
type AuthenticationType string

// Constants representing different authentication types.
const (
	// AuthenticationTypeKey represents an API key authentication type.
	AuthenticationTypeKey AuthenticationType = "api-key"
	// AuthenticationTypeCredentials represents a username and password authentication type.
	AuthenticationTypeCredentials AuthenticationType = "user:password"
)

// Validate checks whether the authentication type and label are valid.
func (a *Authentication) Validate() error {
	// Check if the authentication type is valid (either API key or credentials).
	if a.Type != AuthenticationTypeCredentials && a.Type != AuthenticationTypeKey {
		return errs.Of("invalid authentication type")
	}

	// Ensure that the label is not empty.
	if a.Label == "" {
		return errs.Of("authentication label cannot be empty")
	}

	// If the type is credentials, ensure the label is in the correct format (username:password).
	if a.Type == AuthenticationTypeCredentials {
		if !strings.Contains(a.Label, ":") {
			return errs.Of("credentials must be in format 'username:password'")
		}

		parts := strings.SplitN(a.Label, ":", 2)
		if len(parts) != 2 {
			return errs.Of("credentials must contain exactly a ':' between 'username' and 'password'")
		}

		// Ensure that both username and password are not empty.
		if parts[0] == "" || parts[1] == "" {
			return errs.Of("username and password cannot be empty")
		}
	}

	return nil
}

// GetCredentials returns the username and password if the authentication type is "user:password".
func (a *Authentication) GetCredentials() (string, string) {
	// If validation fails, return empty strings.
	if err := a.Validate(); err != nil {
		return "", ""
	}

	// Split the label into username and password.
	parts := strings.SplitN(a.Label, ":", 2)

	return parts[0], parts[1]
}

// Setup prepares the HTTP request with the necessary authentication headers.
func (a *Authentication) Setup(req *http.Request) error {
	// Validate authentication details before setting up the request.
	if err := a.Validate(); err != nil {
		return err
	}

	// Based on the authentication type, set the appropriate headers for the request.
	switch a.Type {
	case AuthenticationTypeKey:
		req.Header.Set(AuthorizationHeaderLabel, Bearer(a.Label)) // Set the API key in the Authorization header.
	case AuthenticationTypeCredentials:
		// Get the username and password for basic authentication.
		username, password := a.GetCredentials()
		req.SetBasicAuth(username, password) // Set the basic auth credentials.
	default:
		return errs.Of("unsupported authentication type")
	}

	return nil
}
