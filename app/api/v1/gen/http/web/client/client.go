// Code generated by goa v3.19.1, DO NOT EDIT.
//
// web client HTTP transport
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the web service endpoint HTTP clients.
type Client struct {
	// Index Doer is the HTTP client used to make requests to the index endpoint.
	IndexDoer goahttp.Doer

	// Another Doer is the HTTP client used to make requests to the another
	// endpoint.
	AnotherDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the web service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		IndexDoer:           doer,
		AnotherDoer:         doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Index returns an endpoint that makes HTTP requests to the web service index
// server.
func (c *Client) Index() goa.Endpoint {
	var (
		decodeResponse = DecodeIndexResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildIndexRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.IndexDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("web", "index", err)
		}
		return decodeResponse(resp)
	}
}

// Another returns an endpoint that makes HTTP requests to the web service
// another server.
func (c *Client) Another() goa.Endpoint {
	var (
		decodeResponse = DecodeAnotherResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildAnotherRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AnotherDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("web", "another", err)
		}
		return decodeResponse(resp)
	}
}
