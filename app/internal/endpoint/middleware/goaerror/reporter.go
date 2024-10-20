package goaerror

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	goa "goa.design/goa/v3/pkg"

	"github.com/jace-ys/countup/internal/ctxlog"
	"github.com/jace-ys/countup/internal/transport/middleware/reqid"
)

func Reporter(e goa.Endpoint) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		res, err := e(ctx, req)
		if err == nil {
			return res, nil
		}

		var gerr *goa.ServiceError
		if !errors.As(err, &gerr) {
			gerr = goa.Fault("an unexpected error occurred")
		}
		gerr.ID = reqid.RequestIDFromContext(ctx)

		span := trace.SpanFromContext(ctx)
		span.SetStatus(codes.Error, gerr.Name)
		span.SetAttributes(attribute.String("error", err.Error()))
		ctxlog.Error(ctx, "endpoint error", err, ctxlog.KV("err.name", gerr.Name))

		return res, gerr
	}
}
