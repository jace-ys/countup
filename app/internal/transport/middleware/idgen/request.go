package idgen

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"goa.design/clue/log"
	"google.golang.org/grpc"
)

type ctxRequestIDKey struct{}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, next grpc.UnaryHandler) (_ any, err error) {
		ctx = newRequestID(ctx)
		return next(ctx, req)
	}
}

func HTTP() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := newRequestID(r.Context())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func newRequestID(ctx context.Context) context.Context {
	requestID := shortIDGen()
	ctx = context.WithValue(ctx, ctxRequestIDKey{}, requestID)

	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String(log.RequestIDKey, requestID))

	return ctx
}

func shortIDGen() string {
	b := make([]byte, 6)
	io.ReadFull(rand.Reader, b) //nolint:errcheck
	return base64.RawURLEncoding.EncodeToString(b)
}

func RequestIDFromContext(ctx context.Context) string {
	requestID, ok := ctx.Value(ctxRequestIDKey{}).(string)
	if !ok {
		return "null"
	}
	return requestID
}
