// Code generated by goa v3.19.1, DO NOT EDIT.
//
// web HTTP server
//
// Command:
// $ goa gen github.com/jace-ys/countup/api/v1 -o api/v1

package server

import (
	"context"
	"net/http"
	"path"

	web "github.com/jace-ys/countup/api/v1/gen/web"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the web service endpoint HTTP handlers.
type Server struct {
	Mounts              []*MountPoint
	Index               http.Handler
	Another             http.Handler
	LoginGoogle         http.Handler
	LoginGoogleCallback http.Handler
	Logout              http.Handler
	SessionToken        http.Handler
	Static              http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the web service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *web.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	fileSystemStatic http.FileSystem,
) *Server {
	if fileSystemStatic == nil {
		fileSystemStatic = http.Dir(".")
	}
	fileSystemStatic = appendPrefix(fileSystemStatic, "/static")
	return &Server{
		Mounts: []*MountPoint{
			{"Index", "GET", "/"},
			{"Another", "GET", "/another"},
			{"LoginGoogle", "GET", "/login/google"},
			{"LoginGoogleCallback", "GET", "/login/google/callback"},
			{"Logout", "GET", "/logout"},
			{"SessionToken", "GET", "/session/token"},
			{"Serve static/", "GET", "/static/*"},
		},
		Index:               NewIndexHandler(e.Index, mux, decoder, encoder, errhandler, formatter),
		Another:             NewAnotherHandler(e.Another, mux, decoder, encoder, errhandler, formatter),
		LoginGoogle:         NewLoginGoogleHandler(e.LoginGoogle, mux, decoder, encoder, errhandler, formatter),
		LoginGoogleCallback: NewLoginGoogleCallbackHandler(e.LoginGoogleCallback, mux, decoder, encoder, errhandler, formatter),
		Logout:              NewLogoutHandler(e.Logout, mux, decoder, encoder, errhandler, formatter),
		SessionToken:        NewSessionTokenHandler(e.SessionToken, mux, decoder, encoder, errhandler, formatter),
		Static:              http.FileServer(fileSystemStatic),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "web" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Index = m(s.Index)
	s.Another = m(s.Another)
	s.LoginGoogle = m(s.LoginGoogle)
	s.LoginGoogleCallback = m(s.LoginGoogleCallback)
	s.Logout = m(s.Logout)
	s.SessionToken = m(s.SessionToken)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return web.MethodNames[:] }

// Mount configures the mux to serve the web endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountIndexHandler(mux, h.Index)
	MountAnotherHandler(mux, h.Another)
	MountLoginGoogleHandler(mux, h.LoginGoogle)
	MountLoginGoogleCallbackHandler(mux, h.LoginGoogleCallback)
	MountLogoutHandler(mux, h.Logout)
	MountSessionTokenHandler(mux, h.SessionToken)
	MountStatic(mux, http.StripPrefix("/static", h.Static))
}

// Mount configures the mux to serve the web endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountIndexHandler configures the mux to serve the "web" service "Index"
// endpoint.
func MountIndexHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/", f)
}

// NewIndexHandler creates a HTTP handler which loads the HTTP request and
// calls the "web" service "Index" endpoint.
func NewIndexHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeIndexResponse(encoder)
		encodeError    = EncodeIndexError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Index")
		ctx = context.WithValue(ctx, goa.ServiceKey, "web")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountAnotherHandler configures the mux to serve the "web" service "Another"
// endpoint.
func MountAnotherHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/another", f)
}

// NewAnotherHandler creates a HTTP handler which loads the HTTP request and
// calls the "web" service "Another" endpoint.
func NewAnotherHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeAnotherResponse(encoder)
		encodeError    = EncodeAnotherError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Another")
		ctx = context.WithValue(ctx, goa.ServiceKey, "web")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountLoginGoogleHandler configures the mux to serve the "web" service
// "LoginGoogle" endpoint.
func MountLoginGoogleHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/login/google", f)
}

// NewLoginGoogleHandler creates a HTTP handler which loads the HTTP request
// and calls the "web" service "LoginGoogle" endpoint.
func NewLoginGoogleHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeLoginGoogleResponse(encoder)
		encodeError    = EncodeLoginGoogleError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "LoginGoogle")
		ctx = context.WithValue(ctx, goa.ServiceKey, "web")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountLoginGoogleCallbackHandler configures the mux to serve the "web"
// service "LoginGoogleCallback" endpoint.
func MountLoginGoogleCallbackHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/login/google/callback", f)
}

// NewLoginGoogleCallbackHandler creates a HTTP handler which loads the HTTP
// request and calls the "web" service "LoginGoogleCallback" endpoint.
func NewLoginGoogleCallbackHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeLoginGoogleCallbackRequest(mux, decoder)
		encodeResponse = EncodeLoginGoogleCallbackResponse(encoder)
		encodeError    = EncodeLoginGoogleCallbackError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "LoginGoogleCallback")
		ctx = context.WithValue(ctx, goa.ServiceKey, "web")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountLogoutHandler configures the mux to serve the "web" service "Logout"
// endpoint.
func MountLogoutHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/logout", f)
}

// NewLogoutHandler creates a HTTP handler which loads the HTTP request and
// calls the "web" service "Logout" endpoint.
func NewLogoutHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeLogoutRequest(mux, decoder)
		encodeResponse = EncodeLogoutResponse(encoder)
		encodeError    = EncodeLogoutError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Logout")
		ctx = context.WithValue(ctx, goa.ServiceKey, "web")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountSessionTokenHandler configures the mux to serve the "web" service
// "SessionToken" endpoint.
func MountSessionTokenHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/session/token", f)
}

// NewSessionTokenHandler creates a HTTP handler which loads the HTTP request
// and calls the "web" service "SessionToken" endpoint.
func NewSessionTokenHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeSessionTokenRequest(mux, decoder)
		encodeResponse = EncodeSessionTokenResponse(encoder)
		encodeError    = EncodeSessionTokenError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "SessionToken")
		ctx = context.WithValue(ctx, goa.ServiceKey, "web")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// appendFS is a custom implementation of fs.FS that appends a specified prefix
// to the file paths before delegating the Open call to the underlying fs.FS.
type appendFS struct {
	prefix string
	fs     http.FileSystem
}

// Open opens the named file, appending the prefix to the file path before
// passing it to the underlying fs.FS.
func (s appendFS) Open(name string) (http.File, error) {
	switch name {
	case "/*":
		name = "/static"
	}
	return s.fs.Open(path.Join(s.prefix, name))
}

// appendPrefix returns a new fs.FS that appends the specified prefix to file paths
// before delegating to the provided embed.FS.
func appendPrefix(fsys http.FileSystem, prefix string) http.FileSystem {
	return appendFS{prefix: prefix, fs: fsys}
}

// MountStatic configures the mux to serve GET request made to "/static/*".
func MountStatic(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/static/*", h.ServeHTTP)
}
