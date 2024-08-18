package worker

import (
	"context"
	"encoding/json"
	"maps"
	"time"

	"github.com/riverqueue/river"
)

func WithMetadata(ctx context.Context, md JobMetadata) EnqueueOpts {
	return func(o *river.InsertOpts) {
		existingMD, err := parseMetadata(o.Metadata)
		if err != nil {
			existingMD = make(JobMetadata)
		}

		maps.Copy(existingMD, md)
		metadata, err := json.Marshal(existingMD)
		if err != nil {
			return
		}

		o.Metadata = metadata
	}
}

func WithSchedule(schedule time.Time) EnqueueOpts {
	return func(opts *river.InsertOpts) {
		opts.ScheduledAt = schedule
	}
}

func WithMaxAttempts(attempts int) EnqueueOpts {
	return func(opts *river.InsertOpts) {
		opts.MaxAttempts = attempts
	}
}

func WithPriority(priority int) EnqueueOpts {
	return func(opts *river.InsertOpts) {
		opts.Priority = priority
	}
}
