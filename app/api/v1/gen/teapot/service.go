// Code generated by goa v3.21.0, DO NOT EDIT.
//
// teapot service
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package teapot

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Service is the teapot service interface.
type Service interface {
	// Echo implements Echo.
	Echo(context.Context, *EchoPayload) (res *EchoResult, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "countup"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "1.0.0"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "teapot"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"Echo"}

// EchoPayload is the payload type of the teapot service Echo method.
type EchoPayload struct {
	Text string
}

// EchoResult is the result type of the teapot service Echo method.
type EchoResult struct {
	Text string
}

// MakeUnwell builds a goa.ServiceError from an error.
func MakeUnwell(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "unwell", false, false, false)
}
