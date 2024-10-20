package worker

import (
	"encoding/json"
	"maps"
	"time"

	"github.com/riverqueue/river"
)

type EnqueueOption func(*river.InsertOpts)

func WithMetadata(md JobMetadata) EnqueueOption {
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

func WithSchedule(schedule time.Time) EnqueueOption {
	return func(opts *river.InsertOpts) {
		opts.ScheduledAt = schedule
	}
}

func WithMaxAttempts(attempts int) EnqueueOption {
	return func(opts *river.InsertOpts) {
		opts.MaxAttempts = attempts
	}
}

func WithPriority(priority int) EnqueueOption {
	return func(opts *river.InsertOpts) {
		opts.Priority = priority
	}
}
