// Code generated by goa v3.19.1, DO NOT EDIT.
//
// api HTTP server types
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package server

import (
	api "github.com/jace-ys/countup/api/v1/gen/api"
	apiviews "github.com/jace-ys/countup/api/v1/gen/api/views"
	goa "goa.design/goa/v3/pkg"
)

// AuthTokenRequestBody is the type of the "api" service "AuthToken" endpoint
// HTTP request body.
type AuthTokenRequestBody struct {
	Provider    *string `form:"provider,omitempty" json:"provider,omitempty" xml:"provider,omitempty"`
	AccessToken *string `form:"access_token,omitempty" json:"access_token,omitempty" xml:"access_token,omitempty"`
}

// AuthTokenResponseBody is the type of the "api" service "AuthToken" endpoint
// HTTP response body.
type AuthTokenResponseBody struct {
	Token string `form:"token" json:"token" xml:"token"`
}

// CounterGetResponseBody is the type of the "api" service "CounterGet"
// endpoint HTTP response body.
type CounterGetResponseBody struct {
	Count           int32  `form:"count" json:"count" xml:"count"`
	LastIncrementBy string `form:"last_increment_by" json:"last_increment_by" xml:"last_increment_by"`
	LastIncrementAt string `form:"last_increment_at" json:"last_increment_at" xml:"last_increment_at"`
	NextFinalizeAt  string `form:"next_finalize_at" json:"next_finalize_at" xml:"next_finalize_at"`
}

// CounterIncrementResponseBody is the type of the "api" service
// "CounterIncrement" endpoint HTTP response body.
type CounterIncrementResponseBody struct {
	Count           int32  `form:"count" json:"count" xml:"count"`
	LastIncrementBy string `form:"last_increment_by" json:"last_increment_by" xml:"last_increment_by"`
	LastIncrementAt string `form:"last_increment_at" json:"last_increment_at" xml:"last_increment_at"`
	NextFinalizeAt  string `form:"next_finalize_at" json:"next_finalize_at" xml:"next_finalize_at"`
}

// AuthTokenIncompleteAuthInfoResponseBody is the type of the "api" service
// "AuthToken" endpoint HTTP response body for the "incomplete_auth_info" error.
type AuthTokenIncompleteAuthInfoResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// AuthTokenUnauthorizedResponseBody is the type of the "api" service
// "AuthToken" endpoint HTTP response body for the "unauthorized" error.
type AuthTokenUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// AuthTokenForbiddenResponseBody is the type of the "api" service "AuthToken"
// endpoint HTTP response body for the "forbidden" error.
type AuthTokenForbiddenResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CounterGetUnauthorizedResponseBody is the type of the "api" service
// "CounterGet" endpoint HTTP response body for the "unauthorized" error.
type CounterGetUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CounterGetForbiddenResponseBody is the type of the "api" service
// "CounterGet" endpoint HTTP response body for the "forbidden" error.
type CounterGetForbiddenResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CounterIncrementExistingIncrementRequestResponseBody is the type of the
// "api" service "CounterIncrement" endpoint HTTP response body for the
// "existing_increment_request" error.
type CounterIncrementExistingIncrementRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CounterIncrementUnauthorizedResponseBody is the type of the "api" service
// "CounterIncrement" endpoint HTTP response body for the "unauthorized" error.
type CounterIncrementUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CounterIncrementForbiddenResponseBody is the type of the "api" service
// "CounterIncrement" endpoint HTTP response body for the "forbidden" error.
type CounterIncrementForbiddenResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewAuthTokenResponseBody builds the HTTP response body from the result of
// the "AuthToken" endpoint of the "api" service.
func NewAuthTokenResponseBody(res *api.AuthTokenResult) *AuthTokenResponseBody {
	body := &AuthTokenResponseBody{
		Token: res.Token,
	}
	return body
}

// NewCounterGetResponseBody builds the HTTP response body from the result of
// the "CounterGet" endpoint of the "api" service.
func NewCounterGetResponseBody(res *apiviews.CounterInfoView) *CounterGetResponseBody {
	body := &CounterGetResponseBody{
		Count:           *res.Count,
		LastIncrementBy: *res.LastIncrementBy,
		LastIncrementAt: *res.LastIncrementAt,
		NextFinalizeAt:  *res.NextFinalizeAt,
	}
	return body
}

// NewCounterIncrementResponseBody builds the HTTP response body from the
// result of the "CounterIncrement" endpoint of the "api" service.
func NewCounterIncrementResponseBody(res *apiviews.CounterInfoView) *CounterIncrementResponseBody {
	body := &CounterIncrementResponseBody{
		Count:           *res.Count,
		LastIncrementBy: *res.LastIncrementBy,
		LastIncrementAt: *res.LastIncrementAt,
		NextFinalizeAt:  *res.NextFinalizeAt,
	}
	return body
}

// NewAuthTokenIncompleteAuthInfoResponseBody builds the HTTP response body
// from the result of the "AuthToken" endpoint of the "api" service.
func NewAuthTokenIncompleteAuthInfoResponseBody(res *goa.ServiceError) *AuthTokenIncompleteAuthInfoResponseBody {
	body := &AuthTokenIncompleteAuthInfoResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewAuthTokenUnauthorizedResponseBody builds the HTTP response body from the
// result of the "AuthToken" endpoint of the "api" service.
func NewAuthTokenUnauthorizedResponseBody(res *goa.ServiceError) *AuthTokenUnauthorizedResponseBody {
	body := &AuthTokenUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewAuthTokenForbiddenResponseBody builds the HTTP response body from the
// result of the "AuthToken" endpoint of the "api" service.
func NewAuthTokenForbiddenResponseBody(res *goa.ServiceError) *AuthTokenForbiddenResponseBody {
	body := &AuthTokenForbiddenResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCounterGetUnauthorizedResponseBody builds the HTTP response body from the
// result of the "CounterGet" endpoint of the "api" service.
func NewCounterGetUnauthorizedResponseBody(res *goa.ServiceError) *CounterGetUnauthorizedResponseBody {
	body := &CounterGetUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCounterGetForbiddenResponseBody builds the HTTP response body from the
// result of the "CounterGet" endpoint of the "api" service.
func NewCounterGetForbiddenResponseBody(res *goa.ServiceError) *CounterGetForbiddenResponseBody {
	body := &CounterGetForbiddenResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCounterIncrementExistingIncrementRequestResponseBody builds the HTTP
// response body from the result of the "CounterIncrement" endpoint of the
// "api" service.
func NewCounterIncrementExistingIncrementRequestResponseBody(res *goa.ServiceError) *CounterIncrementExistingIncrementRequestResponseBody {
	body := &CounterIncrementExistingIncrementRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCounterIncrementUnauthorizedResponseBody builds the HTTP response body
// from the result of the "CounterIncrement" endpoint of the "api" service.
func NewCounterIncrementUnauthorizedResponseBody(res *goa.ServiceError) *CounterIncrementUnauthorizedResponseBody {
	body := &CounterIncrementUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCounterIncrementForbiddenResponseBody builds the HTTP response body from
// the result of the "CounterIncrement" endpoint of the "api" service.
func NewCounterIncrementForbiddenResponseBody(res *goa.ServiceError) *CounterIncrementForbiddenResponseBody {
	body := &CounterIncrementForbiddenResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewAuthTokenPayload builds a api service AuthToken endpoint payload.
func NewAuthTokenPayload(body *AuthTokenRequestBody) *api.AuthTokenPayload {
	v := &api.AuthTokenPayload{
		Provider:    *body.Provider,
		AccessToken: *body.AccessToken,
	}

	return v
}

// NewCounterIncrementPayload builds a api service CounterIncrement endpoint
// payload.
func NewCounterIncrementPayload(token *string) *api.CounterIncrementPayload {
	v := &api.CounterIncrementPayload{}
	v.Token = token

	return v
}

// ValidateAuthTokenRequestBody runs the validations defined on
// AuthTokenRequestBody
func ValidateAuthTokenRequestBody(body *AuthTokenRequestBody) (err error) {
	if body.Provider == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("provider", "body"))
	}
	if body.AccessToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("access_token", "body"))
	}
	if body.Provider != nil {
		if !(*body.Provider == "google") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.provider", *body.Provider, []any{"google"}))
		}
	}
	return
}
