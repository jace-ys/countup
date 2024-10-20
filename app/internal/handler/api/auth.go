package api

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/markbates/goth/providers/google"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"

	apiv1 "github.com/jace-ys/countup/api/v1"
	"github.com/jace-ys/countup/api/v1/gen/api"
	"github.com/jace-ys/countup/internal/service/user"
)

const (
	jwtClaimsIssuer   = "countup.api.AuthToken"
	jwtClaimsAudience = "count.api.JWTAuth"
)

type AuthTokenClaims struct {
	jwt.RegisteredClaims

	Scopes jwt.ClaimStrings `json:"scopes"`
}

func (h *Handler) AuthToken(ctx context.Context, req *api.AuthTokenPayload) (*api.AuthTokenResult, error) {
	session := &google.Session{
		AccessToken: req.AccessToken,
	}

	fetchedUser, err := h.authn.FetchUser(session)
	if err != nil {
		return nil, goa.Fault("fetch user context from provider %s: %s", req.Provider, err)
	}

	if fetchedUser.Email == "" {
		return nil, api.MakeIncompleteAuthInfo(errors.New("missing field in fetched user context: email"))
	}

	user, err := h.users.CreateUserIfNotExists(ctx, fetchedUser.Email)
	if err != nil {
		return nil, goa.Fault("create new user: %s", err)
	}

	now := time.Now()

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    jwtClaimsIssuer,
			Audience:  jwt.ClaimStrings{jwtClaimsAudience},
			Subject:   user.ID,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
		},
		Scopes: jwt.ClaimStrings{apiv1.AuthScopeAPIUser},
	})

	tok, err := claims.SignedString(h.jwtSigningSecret)
	if err != nil {
		return nil, goa.Fault("sign JWT token: %s", err)
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
		return nil, api.MakeForbidden(fmt.Errorf("parse JWT auth token: %w", err))
	}

	if !tok.Valid {
		return nil, api.MakeForbidden(errors.New("invalid JWT auth token"))
	}

	var missingScopes []string
	for _, scope := range scheme.RequiredScopes {
		if !slices.Contains(claims.Scopes, scope) {
			missingScopes = append(missingScopes, scope)
		}
	}

	if len(missingScopes) > 0 {
		return nil, api.MakeForbidden(
			fmt.Errorf("missing scopes in JWT auth token: %v", missingScopes),
		)
	}

	usr, err := h.users.GetUser(ctx, claims.Subject)
	if err != nil {
		return nil, api.MakeForbidden(errors.New("access denied"))
	}

	return user.ContextWithUser(ctx, usr), nil
}
