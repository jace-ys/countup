// Code generated by goa v3.19.1, DO NOT EDIT.
//
// web service
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package web

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Service is the web service interface.
type Service interface {
	// Index implements Index.
	Index(context.Context) (res []byte, err error)
	// Another implements Another.
	Another(context.Context) (res []byte, err error)
	// LoginGoogle implements LoginGoogle.
	LoginGoogle(context.Context) (res *LoginGoogleResult, err error)
	// LoginGoogleCallback implements LoginGoogleCallback.
	LoginGoogleCallback(context.Context, *LoginGoogleCallbackPayload) (res *LoginGoogleCallbackResult, err error)
	// Logout implements Logout.
	Logout(context.Context, *LogoutPayload) (res *LogoutResult, err error)
	// SessionToken implements SessionToken.
	SessionToken(context.Context, *SessionTokenPayload) (res *SessionTokenResult, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "countup"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "1.0.0"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "web"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [6]string{"Index", "Another", "LoginGoogle", "LoginGoogleCallback", "Logout", "SessionToken"}

// LoginGoogleCallbackPayload is the payload type of the web service
// LoginGoogleCallback method.
type LoginGoogleCallbackPayload struct {
	Code          string
	State         string
	SessionCookie string
}

// LoginGoogleCallbackResult is the result type of the web service
// LoginGoogleCallback method.
type LoginGoogleCallbackResult struct {
	RedirectURL   string
	SessionCookie string
}

// LoginGoogleResult is the result type of the web service LoginGoogle method.
type LoginGoogleResult struct {
	RedirectURL   string
	SessionCookie string
}

// LogoutPayload is the payload type of the web service Logout method.
type LogoutPayload struct {
	SessionCookie string
}

// LogoutResult is the result type of the web service Logout method.
type LogoutResult struct {
	RedirectURL   string
	SessionCookie string
}

// SessionTokenPayload is the payload type of the web service SessionToken
// method.
type SessionTokenPayload struct {
	SessionCookie string
}

// SessionTokenResult is the result type of the web service SessionToken method.
type SessionTokenResult struct {
	Token string
}

// MakeUnauthenticated builds a goa.ServiceError from an error.
func MakeUnauthenticated(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "unauthenticated", false, false, false)
}
