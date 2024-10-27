package web

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/jace-ys/countup/api/v1/gen/api"
	"github.com/jace-ys/countup/api/v1/gen/web"
	"github.com/jace-ys/countup/internal/transport/middleware/idgen"
)

const sessionName = "countup.session"

type authSession struct {
	State   string
	Token   string
	Session string
}

func (h *Handler) LoginGoogle(ctx context.Context) (*web.LoginGoogleResult, error) {
	state := idgen.RequestIDFromContext(ctx)

	session, err := h.authn.BeginAuth(state)
	if err != nil {
		return nil, err
	}

	redirectURL, err := session.GetAuthURL()
	if err != nil {
		return nil, err
	}

	sessionCookie, err := h.cookies.Encode(sessionName, &authSession{
		State:   state,
		Session: session.Marshal(),
	})
	if err != nil {
		return nil, err
	}

	return &web.LoginGoogleResult{
		RedirectURL:   redirectURL,
		SessionCookie: sessionCookie,
	}, nil
}

func (h *Handler) LoginGoogleCallback(ctx context.Context, req *web.LoginGoogleCallbackPayload) (*web.LoginGoogleCallbackResult, error) {
	var auth *authSession
	if err := h.cookies.Decode(sessionName, req.SessionCookie, &auth); err != nil {
		return nil, err
	}

	if req.State != auth.State {
		return nil, web.MakeUnauthorized(errors.New("invalid state"))
	}

	session, err := h.authn.UnmarshalSession(auth.Session)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Set("code", req.Code)

	accessToken, err := session.Authorize(h.authn, params)
	if err != nil {
		return nil, err
	}

	res, err := h.api.AuthToken(ctx, &api.AuthTokenPayload{
		Provider:    "google",
		AccessToken: accessToken,
	})
	if err != nil {
		return nil, err
	}

	sessionCookie, err := h.cookies.Encode(sessionName, &authSession{
		Token: res.Token,
	})
	if err != nil {
		return nil, err
	}

	return &web.LoginGoogleCallbackResult{
		RedirectURL:   "/",
		SessionCookie: sessionCookie,
	}, nil
}

func (h *Handler) Logout(ctx context.Context, req *web.LogoutPayload) (*web.LogoutResult, error) {
	return &web.LogoutResult{
		RedirectURL:   "/",
		SessionCookie: "",
	}, nil
}

func (h *Handler) SessionToken(ctx context.Context, req *web.SessionTokenPayload) (*web.SessionTokenResult, error) {
	var auth *authSession
	if err := h.cookies.Decode(sessionName, req.SessionCookie, &auth); err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", auth)

	return &web.SessionTokenResult{
		Token: auth.Token,
	}, nil
}
