package transport

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"goa.design/clue/log"
	"goa.design/goa/v3/grpc"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"

	"github.com/jace-ys/countup/internal/endpoints"
	"github.com/jace-ys/countup/internal/slog"
)

type GoaGRPCAdapter[E endpoints.GoaEndpoints, S any] struct {
	newFunc GoaGRPCNewFunc[E, S]
}

type GoaGRPCNewFunc[E endpoints.GoaEndpoints, S any] func(e E, uh grpc.UnaryHandler) S

func GoaGRPC[E endpoints.GoaEndpoints, S any](newFunc GoaGRPCNewFunc[E, S]) *GoaGRPCAdapter[E, S] {
	return &GoaGRPCAdapter[E, S]{
		newFunc: newFunc,
	}
}

func (a *GoaGRPCAdapter[E, S]) Adapt(ctx context.Context, ep E) S {
	srv := a.newFunc(ep, nil)
	return srv
}

type GoaHTTPServer interface {
	MethodNames() []string
	Mount(mux goahttp.Muxer)
	Service() string
	Use(m func(http.Handler) http.Handler)
}

type GoaHTTPAdapter[E endpoints.GoaEndpoints, S GoaHTTPServer] struct {
	newFunc   GoaHTTPNewFunc[E, S]
	mountFunc GoaHTTPMountFunc[S]
}

type GoaHTTPNewFunc[E endpoints.GoaEndpoints, S GoaHTTPServer] func(
	e E,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	files http.FileSystem,
) S

type GoaHTTPMountFunc[S GoaHTTPServer] func(
	mux goahttp.Muxer,
	srv S,
)

func GoaHTTP[E endpoints.GoaEndpoints, S GoaHTTPServer](newFunc GoaHTTPNewFunc[E, S], mountFunc GoaHTTPMountFunc[S]) *GoaHTTPAdapter[E, S] {
	return &GoaHTTPAdapter[E, S]{
		newFunc:   newFunc,
		mountFunc: mountFunc,
	}
}

func (a *GoaHTTPAdapter[E, S]) Adapt(ctx context.Context, ep E, files fs.FS) goahttp.ResolverMuxer {
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	formatter := goahttp.NewErrorResponse

	eh := func(ctx context.Context, w http.ResponseWriter, err error) {
		slog.Error(ctx, "failed to encode response", err,
			slog.KV(log.GoaMethodKey, ctx.Value(goa.MethodKey)),
			slog.KV(log.GoaServiceKey, ctx.Value(goa.ServiceKey)),
		)

		gerr := goa.Fault("failed to encode response")

		span := trace.SpanFromContext(ctx)
		span.SetStatus(codes.Error, gerr.GoaErrorName())
		span.SetAttributes(attribute.String("error", fmt.Sprintf("failed to encode response: %v", err)))

		if err := goahttp.ErrorEncoder(enc, formatter)(ctx, w, gerr); err != nil {
			panic(err)
		}
	}

	mux := goahttp.NewMuxer()
	srv := a.newFunc(ep, mux, dec, enc, eh, formatter, http.FS(files))
	a.mountFunc(mux, srv)

	return mux
}
