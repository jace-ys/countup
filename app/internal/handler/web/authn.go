package web

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	goa "goa.design/goa/v3/pkg"

	apiv1 "github.com/jace-ys/countup/api/v1"
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
	reqID := reqid.RequestIDFromContext(ctx)

	session, err := h.authn.BeginAuth(reqID)
	if err != nil {
		return nil, web.MakeUnauthorized(fmt.Errorf("decode session cookie: %w", err))
	}

	redirectURL, err := session.GetAuthURL()
	if err != nil {
		return nil, web.MakeUnauthorized(fmt.Errorf("get redirection URL: %w", err))
	}

	sessionCookie, err := h.cookies.Encode(sessionName, &authSession{
		State:   reqID,
		Session: session.Marshal(),
	})
	if err != nil {
		return nil, web.MakeUnauthorized(fmt.Errorf("encode session cookie: %w", err))
	}

	return &web.LoginGoogleResult{
		RedirectURL:   redirectURL,
		SessionCookie: sessionCookie,
	}, nil
}

func (h *Handler) LoginGoogleCallback(ctx context.Context, req *web.LoginGoogleCallbackPayload) (*web.LoginGoogleCallbackResult, error) {
	var auth *authSession
	if err := h.cookies.Decode(sessionName, req.SessionCookie, &auth); err != nil {
		return nil, web.MakeUnauthorized(fmt.Errorf("decode session cookie: %w", err))
	}

	if req.State != auth.State {
		return nil, web.MakeUnauthorized(errors.New("invalid state"))
	}

	session, err := h.authn.UnmarshalSession(auth.Session)
	if err != nil {
		return nil, web.MakeUnauthorized(fmt.Errorf("unmarshal session data: %w", err))
	}

	params := url.Values{}
	params.Set("code", req.Code)

	accessToken, err := session.Authorize(h.authn, params)
	if err != nil {
		return nil, web.MakeUnauthorized(fmt.Errorf("authorize session: %w", err))
	}

	res, err := h.api.AuthToken(ctx, &api.AuthTokenPayload{
		Provider:    "google",
		AccessToken: accessToken,
	})
	if err != nil {
		var gerr *goa.ServiceError
		if errors.As(err, &gerr) {
			switch gerr.Name {
			case apiv1.ErrCodeIncompleteAuthInfo:
				return nil, web.MakeUnauthorized(errors.New("authentication failed"))
			}
		}
		return nil, web.MakeUnauthorized(errors.New("failed to exchange auth token"))
	}

	sessionCookie, err := h.cookies.Encode(sessionName, &authSession{
		Token: res.Token,
	})
	if err != nil {
		return nil, web.MakeUnauthorized(fmt.Errorf("encode session cookie: %w", err))
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
		return nil, web.MakeUnauthorized(fmt.Errorf("decode session cookie: %w", err))
	}

	return &web.SessionTokenResult{
		Token: auth.Token,
	}, nil
}
