package apiv1

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("countup", func() {
	Title("Count Up")
	Description("A production-ready Go service deployed on Kubernetes")
	Version("1.0.0")
	Server("countup", func() {
		Services("api", "web", "teapot")
		Host("local-http", func() {
			URI("http://localhost:8080")
		})
		Host("local-grpc", func() {
			URI("grpc://localhost:8081")
		})
	})
})

var JWTAuth = JWTSecurity("jwt", func() {
	Scope(AuthScopeAPIUser)
})

var _ = Service("api", func() {
	Security(JWTAuth, func() {
		Scope(AuthScopeAPIUser)
	})

	Error(ErrCodeUnauthorized)
	Error(ErrCodeForbidden)

	HTTP(func() {
		Path("/api/v1")
		Response(ErrCodeUnauthorized, StatusUnauthorized)
		Response(ErrCodeForbidden, StatusForbidden)
	})

	GRPC(func() {
		Response(ErrCodeUnauthorized, CodeUnauthenticated)
		Response(ErrCodeForbidden, CodePermissionDenied)
	})

	Method("AuthToken", func() {
		NoSecurity()

		Error(ErrCodeIncompleteAuthInfo)

		Payload(func() {
			Field(1, "provider", String, func() {
				Enum("google")
			})
			Field(2, "access_token", String)
			Required("provider", "access_token")
		})

		Result(func() {
			Field(1, "token", String)
			Required("token")
		})

		HTTP(func() {
			POST("/auth/token")
			Response(StatusOK)
			Response(ErrCodeIncompleteAuthInfo, StatusUnauthorized)
		})

		GRPC(func() {
			Response(CodeOK)
			Response(ErrCodeIncompleteAuthInfo, CodePermissionDenied)
		})
	})

	Method("CounterGet", func() {
		NoSecurity()

		Result(CounterInfo)

		HTTP(func() {
			GET("/counter")
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("CounterIncrement", func() {
		Error(ErrCodeExistingIncrementRequest, func() {
			Temporary()
		})

		Payload(func() {
			TokenField(1, "token", String)
		})

		Result(CounterInfo)

		HTTP(func() {
			POST("/counter")
			Response(StatusAccepted)
			Response(ErrCodeExistingIncrementRequest, StatusTooManyRequests)
		})

		GRPC(func() {
			Response(CodeOK)
			Response(ErrCodeExistingIncrementRequest, CodeAlreadyExists)
		})
	})

	Files("/openapi.json", "gen/http/openapi3.json")
})

var _ = Service("web", func() {
	Error(ErrCodeUnauthorized)

	HTTP(func() {
		Path("/")
		Response(ErrCodeUnauthorized, StatusUnauthorized)
	})

	withSessionCookie := func() {
		Cookie(CookieNameWebSession)
		CookieSameSite(CookieSameSiteLax)
		CookieMaxAge(86400)
		CookieHTTPOnly()
		CookieSecure()
		CookiePath("/")
	}

	Method("Index", func() {
		Result(Bytes)
		HTTP(func() {
			GET("/")
			Response(StatusOK, func() {
				ContentType("text/html")
			})
		})
	})

	Method("Another", func() {
		Result(Bytes)
		HTTP(func() {
			GET("/another")
			Response(StatusOK, func() {
				ContentType("text/html")
			})
		})
	})

	Method("LoginGoogle", func() {
		Result(func() {
			Attribute("redirect_url", String)
			Attribute("session_cookie", String)
			Required("redirect_url", "session_cookie")
		})

		HTTP(func() {
			GET("/login/google")
			Response(StatusFound, func() {
				Header("redirect_url:Location", String)
				withSessionCookie()
			})
		})
	})

	Method("LoginGoogleCallback", func() {
		Payload(func() {
			Attribute("code", String)
			Attribute("state", String)
			Attribute("session_cookie", String)
			Required("code", "state", "session_cookie")
		})

		Result(func() {
			Attribute("redirect_url", String)
			Attribute("session_cookie", String)
			Required("redirect_url", "session_cookie")
		})

		HTTP(func() {
			GET("/login/google/callback")
			Params(func() {
				Param("code", String)
				Param("state", String)
				Required("code", "state")
			})
			Cookie(CookieNameWebSession)
			Response(StatusFound, func() {
				Header("redirect_url:Location", String)
				withSessionCookie()
			})
		})
	})

	Method("Logout", func() {
		Payload(func() {
			Attribute("session_cookie", String)
			Required("session_cookie")
		})

		Result(func() {
			Attribute("redirect_url", String)
			Attribute("session_cookie", String)
			Required("redirect_url", "session_cookie")
		})

		HTTP(func() {
			GET("/logout")
			Cookie(CookieNameWebSession)
			Response(StatusFound, func() {
				Header("redirect_url:Location", String)
				withSessionCookie()
			})
		})
	})

	Method("SessionToken", func() {
		Payload(func() {
			Attribute("session_cookie", String)
			Required("session_cookie")
		})

		Result(func() {
			Attribute("token", String)
			Required("token")
		})

		HTTP(func() {
			GET("/session/token")
			Cookie(CookieNameWebSession)
			Response(StatusOK, func() {
				ContentType("application/json")
			})
		})
	})

	Files("/static/*", "static/")
})

var _ = Service("teapot", func() {
	Error("unwell")

	Method("Echo", func() {
		Payload(func() {
			Field(1, "text", String)
			Required("text")
		})

		Result(func() {
			Field(1, "text", String)
			Required("text")
		})

		HTTP(func() {
			POST("/echo")
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Files("/openapi.json", "gen/http/openapi3.json")
})
