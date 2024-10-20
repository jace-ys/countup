// Code generated by goa v3.19.1, DO NOT EDIT.
//
// web HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	web "github.com/jace-ys/countup/api/v1/gen/web"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildIndexRequest instantiates a HTTP request object with method and path
// set to call the "web" service "Index" endpoint
func (c *Client) BuildIndexRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: IndexWebPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("web", "Index", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeIndexResponse returns a decoder for responses returned by the web
// Index endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeIndexResponse may return the following errors:
//   - "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//   - error: internal error
func DecodeIndexResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body []byte
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("web", "Index", err)
			}
			return body, nil
		case http.StatusUnauthorized:
			var (
				body IndexUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("web", "Index", err)
			}
			err = ValidateIndexUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("web", "Index", err)
			}
			return nil, NewIndexUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("web", "Index", resp.StatusCode, string(body))
		}
	}
}

// BuildAnotherRequest instantiates a HTTP request object with method and path
// set to call the "web" service "Another" endpoint
func (c *Client) BuildAnotherRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AnotherWebPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("web", "Another", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeAnotherResponse returns a decoder for responses returned by the web
// Another endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeAnotherResponse may return the following errors:
//   - "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//   - error: internal error
func DecodeAnotherResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body []byte
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("web", "Another", err)
			}
			return body, nil
		case http.StatusUnauthorized:
			var (
				body AnotherUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("web", "Another", err)
			}
			err = ValidateAnotherUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("web", "Another", err)
			}
			return nil, NewAnotherUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("web", "Another", resp.StatusCode, string(body))
		}
	}
}

// BuildLoginGoogleRequest instantiates a HTTP request object with method and
// path set to call the "web" service "LoginGoogle" endpoint
func (c *Client) BuildLoginGoogleRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: LoginGoogleWebPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("web", "LoginGoogle", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeLoginGoogleResponse returns a decoder for responses returned by the
// web LoginGoogle endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeLoginGoogleResponse may return the following errors:
//   - "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//   - error: internal error
func DecodeLoginGoogleResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusFound:
			var (
				redirectURL string
				err         error
			)
			redirectURLRaw := resp.Header.Get("Location")
			if redirectURLRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("redirect_url", "header"))
			}
			redirectURL = redirectURLRaw
			var (
				sessionCookie    string
				sessionCookieRaw string

				cookies = resp.Cookies()
			)
			for _, c := range cookies {
				switch c.Name {
				case "countup.session":
					sessionCookieRaw = c.Value
				}
			}
			if sessionCookieRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("session_cookie", "cookie"))
			}
			sessionCookie = sessionCookieRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("web", "LoginGoogle", err)
			}
			res := NewLoginGoogleResultFound(redirectURL, sessionCookie)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body LoginGoogleUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("web", "LoginGoogle", err)
			}
			err = ValidateLoginGoogleUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("web", "LoginGoogle", err)
			}
			return nil, NewLoginGoogleUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("web", "LoginGoogle", resp.StatusCode, string(body))
		}
	}
}

// BuildLoginGoogleCallbackRequest instantiates a HTTP request object with
// method and path set to call the "web" service "LoginGoogleCallback" endpoint
func (c *Client) BuildLoginGoogleCallbackRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: LoginGoogleCallbackWebPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("web", "LoginGoogleCallback", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeLoginGoogleCallbackRequest returns an encoder for requests sent to the
// web LoginGoogleCallback server.
func EncodeLoginGoogleCallbackRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*web.LoginGoogleCallbackPayload)
		if !ok {
			return goahttp.ErrInvalidType("web", "LoginGoogleCallback", "*web.LoginGoogleCallbackPayload", v)
		}
		{
			v := p.SessionCookie
			req.AddCookie(&http.Cookie{
				Name:  "countup.session",
				Value: v,
			})
		}
		values := req.URL.Query()
		values.Add("code", p.Code)
		values.Add("state", p.State)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeLoginGoogleCallbackResponse returns a decoder for responses returned
// by the web LoginGoogleCallback endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeLoginGoogleCallbackResponse may return the following errors:
//   - "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//   - error: internal error
func DecodeLoginGoogleCallbackResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusFound:
			var (
				redirectURL string
				err         error
			)
			redirectURLRaw := resp.Header.Get("Location")
			if redirectURLRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("redirect_url", "header"))
			}
			redirectURL = redirectURLRaw
			var (
				sessionCookie    string
				sessionCookieRaw string

				cookies = resp.Cookies()
			)
			for _, c := range cookies {
				switch c.Name {
				case "countup.session":
					sessionCookieRaw = c.Value
				}
			}
			if sessionCookieRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("session_cookie", "cookie"))
			}
			sessionCookie = sessionCookieRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("web", "LoginGoogleCallback", err)
			}
			res := NewLoginGoogleCallbackResultFound(redirectURL, sessionCookie)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body LoginGoogleCallbackUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("web", "LoginGoogleCallback", err)
			}
			err = ValidateLoginGoogleCallbackUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("web", "LoginGoogleCallback", err)
			}
			return nil, NewLoginGoogleCallbackUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("web", "LoginGoogleCallback", resp.StatusCode, string(body))
		}
	}
}

// BuildLogoutRequest instantiates a HTTP request object with method and path
// set to call the "web" service "Logout" endpoint
func (c *Client) BuildLogoutRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: LogoutWebPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("web", "Logout", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeLogoutRequest returns an encoder for requests sent to the web Logout
// server.
func EncodeLogoutRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*web.LogoutPayload)
		if !ok {
			return goahttp.ErrInvalidType("web", "Logout", "*web.LogoutPayload", v)
		}
		{
			v := p.SessionCookie
			req.AddCookie(&http.Cookie{
				Name:  "countup.session",
				Value: v,
			})
		}
		return nil
	}
}

// DecodeLogoutResponse returns a decoder for responses returned by the web
// Logout endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeLogoutResponse may return the following errors:
//   - "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//   - error: internal error
func DecodeLogoutResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusFound:
			var (
				redirectURL string
				err         error
			)
			redirectURLRaw := resp.Header.Get("Location")
			if redirectURLRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("redirect_url", "header"))
			}
			redirectURL = redirectURLRaw
			var (
				sessionCookie    string
				sessionCookieRaw string

				cookies = resp.Cookies()
			)
			for _, c := range cookies {
				switch c.Name {
				case "countup.session":
					sessionCookieRaw = c.Value
				}
			}
			if sessionCookieRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("session_cookie", "cookie"))
			}
			sessionCookie = sessionCookieRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("web", "Logout", err)
			}
			res := NewLogoutResultFound(redirectURL, sessionCookie)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body LogoutUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("web", "Logout", err)
			}
			err = ValidateLogoutUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("web", "Logout", err)
			}
			return nil, NewLogoutUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("web", "Logout", resp.StatusCode, string(body))
		}
	}
}

// BuildSessionTokenRequest instantiates a HTTP request object with method and
// path set to call the "web" service "SessionToken" endpoint
func (c *Client) BuildSessionTokenRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SessionTokenWebPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("web", "SessionToken", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSessionTokenRequest returns an encoder for requests sent to the web
// SessionToken server.
func EncodeSessionTokenRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*web.SessionTokenPayload)
		if !ok {
			return goahttp.ErrInvalidType("web", "SessionToken", "*web.SessionTokenPayload", v)
		}
		{
			v := p.SessionCookie
			req.AddCookie(&http.Cookie{
				Name:  "countup.session",
				Value: v,
			})
		}
		return nil
	}
}

// DecodeSessionTokenResponse returns a decoder for responses returned by the
// web SessionToken endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeSessionTokenResponse may return the following errors:
//   - "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//   - error: internal error
func DecodeSessionTokenResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SessionTokenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("web", "SessionToken", err)
			}
			err = ValidateSessionTokenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("web", "SessionToken", err)
			}
			res := NewSessionTokenResultOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body SessionTokenUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("web", "SessionToken", err)
			}
			err = ValidateSessionTokenUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("web", "SessionToken", err)
			}
			return nil, NewSessionTokenUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("web", "SessionToken", resp.StatusCode, string(body))
		}
	}
}
