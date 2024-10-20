package web

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/jace-ys/countup/api/v1/gen/api"
	"github.com/jace-ys/countup/api/v1/gen/web"
	"github.com/jace-ys/countup/internal/transport/middleware/reqid"
)

const sessionName = "countup.session"

type authSession struct {
	State   string
	Token   string
	Session string
}

func (h *Handler) LoginGoogle(ctx context.Context) (*web.LoginGoogleResult, error) {
	requestID := reqid.RequestIDFromContext(ctx)

	session, err := h.authn.BeginAuth(requestID.String())
	if err != nil {
		return nil, fmt.Errorf("auth provider: begin session: %w", err)
	}

	redirectURL, err := session.GetAuthURL()
	if err != nil {
		return nil, fmt.Errorf("auth provider: get redirection URL: %w", err)
	}

	sessionCookie, err := h.cookies.Encode(sessionName, &authSession{
		State:   requestID.String(),
		Session: session.Marshal(),
	})
	if err != nil {
		return nil, fmt.Errorf("encode session cookie: %w", err)
	}

	return &web.LoginGoogleResult{
		RedirectURL:   redirectURL,
		SessionCookie: sessionCookie,
	}, nil
}

func (h *Handler) LoginGoogleCallback(ctx context.Context, req *web.LoginGoogleCallbackPayload) (*web.LoginGoogleCallbackResult, error) {
	var auth *authSession
	if err := h.cookies.Decode(sessionName, req.SessionCookie, &auth); err != nil {
		return nil, fmt.Errorf("decode session cookie: %w", err)
	}

	if req.State != auth.State {
		return nil, web.MakeUnauthenticated(errors.New("auth state mismatch"))
	}

	session, err := h.authn.UnmarshalSession(auth.Session)
	if err != nil {
		return nil, fmt.Errorf("unmarshal session data: %w", err)
	}

	params := url.Values{}
	params.Set("code", req.Code)

	accessToken, err := session.Authorize(h.authn, params)
	if err != nil {
		return nil, web.MakeUnauthenticated(fmt.Errorf("auth provider: get access token: %w", err))
	}

	res, err := h.api.AuthToken(ctx, &api.AuthTokenPayload{
		Provider:    "google",
		AccessToken: accessToken,
	})
	if err != nil {
		return nil, web.MakeUnauthenticated(fmt.Errorf("exchange auth token: %w", err))
	}

	sessionCookie, err := h.cookies.Encode(sessionName, &authSession{
		Token: res.Token,
	})
	if err != nil {
		return nil, fmt.Errorf("encode session cookie: %w", err)
	}

	return &web.LoginGoogleCallbackResult{
		RedirectURL:   "/",
		SessionCookie: sessionCookie,
	}, nil
}

func (h *Handler) Logout(ctx context.Context, req *web.LogoutPayload) (*web.LogoutResult, error) {
	sessionCookie, err := h.cookies.Encode(sessionName, &authSession{})
	if err != nil {
		return nil, fmt.Errorf("encode session cookie: %w", err)
	}

	return &web.LogoutResult{
		RedirectURL:   "/",
		SessionCookie: sessionCookie,
	}, nil
}

func (h *Handler) SessionToken(ctx context.Context, req *web.SessionTokenPayload) (*web.SessionTokenResult, error) {
	var auth *authSession
	if err := h.cookies.Decode(sessionName, req.SessionCookie, &auth); err != nil {
		return nil, fmt.Errorf("decode session cookie: %w", err)
	}

	return &web.SessionTokenResult{
		Token: auth.Token,
	}, nil
}
