package user

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/jace-ys/countup/internal/ctxlog"
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

	ctx = ctxlog.With(ctx, ctxlog.KV("user.id", user.ID), ctxlog.KV("user.email", user.Email))

	return ctx
}

func UserFromContext(ctx context.Context) *User {
	user, ok := ctx.Value(ctxKeyUserInfo{}).(*User)
	if !ok {
		return nil
	}
	return user
}
