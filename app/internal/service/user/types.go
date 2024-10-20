package user

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/jace-ys/countup/internal/slog"
)

type User struct {
	ID    string
	Email string
}

type ctxKeyUserInfo struct{}

func ContextWithUser(ctx context.Context, user *User) context.Context {
	ctx = context.WithValue(ctx, ctxKeyUserInfo{}, user)

	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("user.id", user.ID))
	span.SetAttributes(attribute.String("user.email", user.Email))

	ctx = slog.With(ctx, slog.KV("user.id", user.ID), slog.KV("user.email", user.Email))

	return ctx
}

func UserFromContext(ctx context.Context) *User {
	user, ok := ctx.Value(ctxKeyUserInfo{}).(*User)
	if !ok {
		return nil
	}
	return user
}
