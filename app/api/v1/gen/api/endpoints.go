// Code generated by goa v3.19.1, DO NOT EDIT.
//
// api endpoints
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package api

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "api" service endpoints.
type Endpoints struct {
	AuthToken        goa.Endpoint
	CounterGet       goa.Endpoint
	CounterIncrement goa.Endpoint
}

// NewEndpoints wraps the methods of the "api" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		AuthToken:        NewAuthTokenEndpoint(s),
		CounterGet:       NewCounterGetEndpoint(s, a.JWTAuth),
		CounterIncrement: NewCounterIncrementEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "api" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.AuthToken = m(e.AuthToken)
	e.CounterGet = m(e.CounterGet)
	e.CounterIncrement = m(e.CounterIncrement)
}

// NewAuthTokenEndpoint returns an endpoint function that calls the method
// "AuthToken" of service "api".
func NewAuthTokenEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*AuthTokenPayload)
		return s.AuthToken(ctx, p)
	}
}

// NewCounterGetEndpoint returns an endpoint function that calls the method
// "CounterGet" of service "api".
func NewCounterGetEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CounterGetPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api"},
			RequiredScopes: []string{"api"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.CounterGet(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCounterInfo(res, "default")
		return vres, nil
	}
}

// NewCounterIncrementEndpoint returns an endpoint function that calls the
// method "CounterIncrement" of service "api".
func NewCounterIncrementEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CounterIncrementPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api"},
			RequiredScopes: []string{"api"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.CounterIncrement(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCounterInfo(res, "default")
		return vres, nil
	}
}