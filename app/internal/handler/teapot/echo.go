package teapot

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/riverqueue/river"

	"github.com/jace-ys/countup/api/v1/gen/teapot"
	"github.com/jace-ys/countup/internal/worker"
)

func (h *Handler) Echo(ctx context.Context, req *teapot.EchoPayload) (*teapot.EchoResult, error) {
	switch {
	case strings.HasPrefix(req.Text, "error: "):
		msg := strings.TrimPrefix(req.Text, "error: ")
		h.workers.Enqueue(ctx, &EchoJobArgs{Error: msg}, worker.WithMaxAttempts(3))
		return nil, teapot.MakeUnwell(errors.New(msg))

	case strings.HasPrefix(req.Text, "panic: "):
		msg := strings.TrimPrefix(req.Text, "panic: ")
		h.workers.Enqueue(ctx, &EchoJobArgs{Panic: msg}, worker.WithMaxAttempts(3))
		panic(msg)

	case strings.HasPrefix(req.Text, "cancel: "):
		msg := strings.TrimPrefix(req.Text, "cancel: ")
		h.workers.Enqueue(ctx, &EchoJobArgs{Cancel: msg}, worker.WithMaxAttempts(3))
		return nil, teapot.MakeUnwell(errors.New(msg))

	case strings.HasPrefix(req.Text, "sleep: "):
		msg := strings.TrimPrefix(req.Text, "sleep: ")
		duration, err := time.ParseDuration(msg)
		if err != nil {
			duration = time.Second
		}
		time.Sleep(duration)
		h.workers.Enqueue(ctx, &EchoJobArgs{Sleep: duration}, worker.WithMaxAttempts(3))
	}

	return &teapot.EchoResult{Text: req.Text}, nil
}

type EchoJobArgs struct {
	Error  string
	Panic  string
	Cancel string
	Sleep  time.Duration
}

func (EchoJobArgs) Kind() string {
	return "countup.Echo"
}

type EchoWorker struct {
	river.WorkerDefaults[EchoJobArgs]
}

func (w *EchoWorker) Work(ctx context.Context, job *river.Job[EchoJobArgs]) error {
	switch {
	case job.Args.Error != "":
		return errors.New(job.Args.Error)
	case job.Args.Panic != "":
		panic(job.Args.Panic)
	case job.Args.Cancel != "":
		return river.JobCancel(errors.New(job.Args.Cancel))
	case job.Args.Sleep != 0:
		time.Sleep(job.Args.Sleep)
	}

	return nil
}
