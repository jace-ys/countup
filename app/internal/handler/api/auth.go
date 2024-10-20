package api

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/markbates/goth/providers/google"
	"goa.design/goa/v3/security"

	apiv1 "github.com/jace-ys/countup/api/v1"
	"github.com/jace-ys/countup/api/v1/gen/api"
	"github.com/jace-ys/countup/internal/idgen"
	"github.com/jace-ys/countup/internal/service/user"
)

const (
	jwtClaimsIssuer   = "countup.api.AuthToken"
	jwtClaimsAudience = "countup.api.JWTAuth"
)

type AuthTokenClaims struct {
	jwt.RegisteredClaims

	Scopes jwt.ClaimStrings `json:"scopes"`
}

func (h *Handler) AuthToken(ctx context.Context, req *api.AuthTokenPayload) (*api.AuthTokenResult, error) {
	authedUser, err := h.authn.FetchUser(&google.Session{
		AccessToken: req.AccessToken,
	})
	if err != nil {
		return nil, fmt.Errorf("auth provider %s: fetch user context: %w", req.Provider, err)
	}

	if authedUser.Email == "" {
		return nil, api.MakeUnauthenticated(errors.New("missing field in fetched user context: email"))
	}

	user, err := h.users.CreateUserIfNotExists(ctx, authedUser.Email)
	if err != nil {
		return nil, fmt.Errorf("create new user: %w", err)
	}

	now := time.Now()

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    jwtClaimsIssuer,
			Audience:  jwt.ClaimStrings{jwtClaimsAudience},
			Subject:   user.ID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
		},
		Scopes: jwt.ClaimStrings{apiv1.AuthScopeAPIUser},
	})

	tok, err := claims.SignedString(h.jwtSigningSecret)
	if err != nil {
		return nil, fmt.Errorf("sign JWT token: %w", err)
	}

	return &api.AuthTokenResult{
		Token: tok,
	}, nil
}

var _ api.Auther = (*Handler)(nil)

func (h *Handler) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuer(jwtClaimsIssuer),
		jwt.WithAudience(jwtClaimsAudience),
		jwt.WithIssuedAt(),
		jwt.WithExpirationRequired(),
	)

	var claims AuthTokenClaims
	tok, err := parser.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return h.jwtSigningSecret, nil
	})
	if err != nil {
		return nil, api.MakeAccessDenied(fmt.Errorf("parse JWT auth token: %w", err))
	}

	if !tok.Valid {
		return nil, api.MakeAccessDenied(errors.New("invalid JWT auth token"))
	}

	var missingScopes []string
	for _, scope := range scheme.RequiredScopes {
		if !slices.Contains(claims.Scopes, scope) {
			missingScopes = append(missingScopes, scope)
		}
	}

	if len(missingScopes) > 0 {
		return nil, api.MakeAccessDenied(fmt.Errorf("missing authorization scopes: %v", missingScopes))
	}

	userID, err := idgen.FromString[idgen.User](claims.Subject)
	if err != nil {
		return nil, api.MakeAccessDenied(fmt.Errorf("parse user ID from claims: %w", err))
	}

	usr, err := h.users.GetUser(ctx, userID)
	if err != nil {
		return nil, api.MakeAccessDenied(fmt.Errorf("get user %s: %w", userID, err))
	}

	return user.ContextWithUser(ctx, usr), nil
}
