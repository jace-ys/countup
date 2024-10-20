package apiv1

import (
	"embed"

	. "goa.design/goa/v3/dsl"
)

//go:embed gen/http/*.json gen/http/*.yaml
var OpenAPIFS embed.FS

const (
	AuthScopeAPIUser = "api.user"
)

const (
	ErrCodeUnauthenticated          = "unauthenticated"
	ErrCodeAccessDenied             = "access_denied"
	ErrCodeExistingIncrementRequest = "existing_increment_request"
)

const (
	CookieNameWebSession = "session_cookie:countup.session"
)

var CounterInfo = ResultType("application/vnd.countup.counter-info`", "CounterInfo", func() {
	Field(1, "count", Int32)
	Field(2, "last_increment_by", String)
	Field(3, "last_increment_at", String)
	Field(4, "next_finalize_at", String)
	Required("count", "last_increment_by", "last_increment_at", "next_finalize_at")
})
