package rpc

import (
	"net/http"
	"strings"

	"github.com/avila-r/bitclient/errs"
)

type (
	Authentication struct {
		Type  AuthenticationType
		Label string
	}

	AuthenticationType string
)

const (
	AuthenticationTypeKey         AuthenticationType = "api-key"
	AuthenticationTypeCredentials AuthenticationType = "user:password"
)

func (a *Authentication) Validate() error {
	if a.Type != AuthenticationTypeCredentials && a.Type != AuthenticationTypeKey {
		return errs.Of("invalid authentication type")
	}

	if a.Label == "" {
		return errs.Of("authentication label cannot be empty")
	}

	if a.Type == AuthenticationTypeCredentials {
		if !strings.Contains(a.Label, ":") {
			return errs.Of("credentials must be in format 'username:password'")
		}

		parts := strings.SplitN(a.Label, ":", 2)
		if len(parts) != 2 {
			return errs.Of("credentials must contain exactly a ':' between 'username' and 'password'")
		}

		if parts[0] == "" || parts[1] == "" {
			return errs.Of("username and password cannot be empty")
		}
	}

	return nil
}

func (a *Authentication) GetCredentials() (string, string) {
	if err := a.Validate(); err != nil {
		return "", ""
	}

	parts := strings.SplitN(a.Label, ":", 2)

	return parts[0], parts[1]
}

func (a *Authentication) Setup(req *http.Request) error {
	if err := a.Validate(); err != nil {
		return err
	}

	switch a.Type {
	case AuthenticationTypeKey:
		req.Header.Set(AuthorizationHeaderLabel, Bearer(a.Label))
	case AuthenticationTypeCredentials:
		username, password := a.GetCredentials()
		req.SetBasicAuth(username, password)
	default:
		return errs.Of("unsupported authentication type")
	}

	return nil
}
