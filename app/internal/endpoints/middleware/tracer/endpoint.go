package tracer

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	goa "goa.design/goa/v3/pkg"

	"github.com/jace-ys/countup/internal/instrument"
)

func Endpoint(e goa.Endpoint) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		service := "unknown"
		if s, ok := ctx.Value(goa.ServiceKey).(string); ok {
			service = s
		}

		method := "unknown"
		if m, ok := ctx.Value(goa.MethodKey).(string); ok {
			method = m
		}

		ctx, span := instrument.OTel.Tracer().Start(ctx, fmt.Sprintf("goa.endpoint/%s.%s", service, method))
		span.SetAttributes(attribute.String("endpoint.service", service), attribute.String("endpoint.method", method))
		defer span.End()

		return e(ctx, req)
	}
}
