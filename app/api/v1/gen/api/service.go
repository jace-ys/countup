// Code generated by goa v3.19.1, DO NOT EDIT.
//
// api service
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package api

import (
	"context"

	apiviews "github.com/jace-ys/countup/api/v1/gen/api/views"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Service is the api service interface.
type Service interface {
	// AuthToken implements AuthToken.
	AuthToken(context.Context, *AuthTokenPayload) (res *AuthTokenResult, err error)
	// CounterGet implements CounterGet.
	CounterGet(context.Context) (res *CounterInfo, err error)
	// CounterIncrement implements CounterIncrement.
	CounterIncrement(context.Context, *CounterIncrementPayload) (res *CounterInfo, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "countup"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "1.0.0"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "api"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"AuthToken", "CounterGet", "CounterIncrement"}

// AuthTokenPayload is the payload type of the api service AuthToken method.
type AuthTokenPayload struct {
	Provider    string
	AccessToken string
}

// AuthTokenResult is the result type of the api service AuthToken method.
type AuthTokenResult struct {
	Token string
}

// CounterIncrementPayload is the payload type of the api service
// CounterIncrement method.
type CounterIncrementPayload struct {
	Token *string
}

// CounterInfo is the result type of the api service CounterGet method.
type CounterInfo struct {
	Count           int32
	LastIncrementBy string
	LastIncrementAt string
	NextFinalizeAt  string
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "unauthorized", false, false, false)
}

// MakeForbidden builds a goa.ServiceError from an error.
func MakeForbidden(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "forbidden", false, false, false)
}

// MakeIncompleteAuthInfo builds a goa.ServiceError from an error.
func MakeIncompleteAuthInfo(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "incomplete_auth_info", false, false, false)
}

// MakeExistingIncrementRequest builds a goa.ServiceError from an error.
func MakeExistingIncrementRequest(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "existing_increment_request", false, true, false)
}

// NewCounterInfo initializes result type CounterInfo from viewed result type
// CounterInfo.
func NewCounterInfo(vres *apiviews.CounterInfo) *CounterInfo {
	return newCounterInfo(vres.Projected)
}

// NewViewedCounterInfo initializes viewed result type CounterInfo from result
// type CounterInfo using the given view.
func NewViewedCounterInfo(res *CounterInfo, view string) *apiviews.CounterInfo {
	p := newCounterInfoView(res)
	return &apiviews.CounterInfo{Projected: p, View: "default"}
}

// newCounterInfo converts projected type CounterInfo to service type
// CounterInfo.
func newCounterInfo(vres *apiviews.CounterInfoView) *CounterInfo {
	res := &CounterInfo{}
	if vres.Count != nil {
		res.Count = *vres.Count
	}
	if vres.LastIncrementBy != nil {
		res.LastIncrementBy = *vres.LastIncrementBy
	}
	if vres.LastIncrementAt != nil {
		res.LastIncrementAt = *vres.LastIncrementAt
	}
	if vres.NextFinalizeAt != nil {
		res.NextFinalizeAt = *vres.NextFinalizeAt
	}
	return res
}

// newCounterInfoView projects result type CounterInfo to projected type
// CounterInfoView using the "default" view.
func newCounterInfoView(res *CounterInfo) *apiviews.CounterInfoView {
	vres := &apiviews.CounterInfoView{
		Count:           &res.Count,
		LastIncrementBy: &res.LastIncrementBy,
		LastIncrementAt: &res.LastIncrementAt,
		NextFinalizeAt:  &res.NextFinalizeAt,
	}
	return vres
}
