package api

import (
	"context"
	"errors"

	goa "goa.design/goa/v3/pkg"

	"github.com/jace-ys/countup/api/v1/gen/api"
	"github.com/jace-ys/countup/internal/service/counter"
	"github.com/jace-ys/countup/internal/service/user"
)

func (h *Handler) CounterGet(ctx context.Context) (*api.CounterInfo, error) {
	info, err := h.counter.GetInfo(ctx)
	if err != nil {
		return nil, goa.Fault("get counter info: %s", err)
	}

	return &api.CounterInfo{
		Count:           info.Count,
		LastIncrementBy: info.LastIncrementBy,
		LastIncrementAt: info.LastIncrementAtTimestamp(),
		NextFinalizeAt:  info.NextFinalizeAtTimestamp(),
	}, nil
}

func (h *Handler) CounterIncrement(ctx context.Context, req *api.CounterIncrementPayload) (*api.CounterInfo, error) {
	usr := user.UserFromContext(ctx)

	if err := h.counter.RequestIncrement(ctx, usr.Email); err != nil {
		var multipleIncrementRequestErr *counter.MultipleIncrementRequestError
		switch {
		case errors.As(err, &multipleIncrementRequestErr):
			return nil, api.MakeExistingIncrementRequest(
				errors.New("user already made an increment request in the recent finalize window, please try again after the next finalize time"),
			)
		default:
			return nil, goa.Fault("request increment: %s", err)
		}
	}

	info, err := h.counter.GetInfo(ctx)
	if err != nil {
		return nil, goa.Fault("get counter info: %s", err)
	}

	return &api.CounterInfo{
		Count:           info.Count,
		LastIncrementBy: info.LastIncrementBy,
		LastIncrementAt: info.LastIncrementAtTimestamp(),
		NextFinalizeAt:  info.NextFinalizeAtTimestamp(),
	}, nil
}
