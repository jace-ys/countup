package api

import (
	"context"
	"fmt"

	"github.com/markbates/goth/providers/google"
	"goa.design/goa/v3/security"

	"github.com/jace-ys/countup/api/v1/gen/api"
)

var _ api.Auther = (*Handler)(nil)

func (h *Handler) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	fmt.Println("token", token)
	return ctx, nil
}

func (h *Handler) AuthToken(ctx context.Context, req *api.AuthTokenPayload) (*api.AuthTokenResult, error) {
	session := &google.Session{
		AccessToken: req.AccessToken,
	}

	user, err := h.authn.FetchUser(session)
	if err != nil {
		return nil, err
	}

	fmt.Println(user.Email)

	return &api.AuthTokenResult{
		Token: req.AccessToken,
	}, nil
}
