// Code generated by goa v3.19.1, DO NOT EDIT.
//
// teapot HTTP server types
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package server

import (
	teapot "github.com/jace-ys/countup/api/v1/gen/teapot"
	goa "goa.design/goa/v3/pkg"
)

// EchoRequestBody is the type of the "teapot" service "Echo" endpoint HTTP
// request body.
type EchoRequestBody struct {
	Text *string `form:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
}

// EchoResponseBody is the type of the "teapot" service "Echo" endpoint HTTP
// response body.
type EchoResponseBody struct {
	Text string `form:"text" json:"text" xml:"text"`
}

// NewEchoResponseBody builds the HTTP response body from the result of the
// "Echo" endpoint of the "teapot" service.
func NewEchoResponseBody(res *teapot.EchoResult) *EchoResponseBody {
	body := &EchoResponseBody{
		Text: res.Text,
	}
	return body
}

// NewEchoPayload builds a teapot service Echo endpoint payload.
func NewEchoPayload(body *EchoRequestBody) *teapot.EchoPayload {
	v := &teapot.EchoPayload{
		Text: *body.Text,
	}

	return v
}

// ValidateEchoRequestBody runs the validations defined on EchoRequestBody
func ValidateEchoRequestBody(body *EchoRequestBody) (err error) {
	if body.Text == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("text", "body"))
	}
	return
}
