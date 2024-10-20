package worker

import (
	"context"
	"encoding/json"

	"go.opentelemetry.io/otel/propagation"
	"goa.design/clue/log"

	"github.com/jace-ys/countup/internal/instrument"
	"github.com/jace-ys/countup/internal/transport/middleware/reqid"
)

var _ propagation.TextMapCarrier = (*JobMetadata)(nil)

type JobMetadata = propagation.MapCarrier

func withContextMetadata(ctx context.Context) EnqueueOption {
	md := make(JobMetadata)

	md[log.RequestIDKey] = reqid.RequestIDFromContext(ctx).String()
	instrument.OTel.Propagators().Inject(ctx, md)

	return WithMetadata(md)
}

func parseMetadata(metadata []byte) (JobMetadata, error) {
	md := make(JobMetadata)
	err := json.Unmarshal(metadata, &md)
	return md, err //nolint:wrapcheck
}
