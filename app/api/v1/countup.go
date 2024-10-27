package apiv1

import (
	"embed"

	. "goa.design/goa/v3/dsl"
)

var (
	//go:embed gen/http/*.json gen/http/*.yaml
	OpenAPIFS embed.FS
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
	Scope("api")
})

var CounterInfo = ResultType("application/vnd.countup.counter-info`", "CounterInfo", func() {
	Field(1, "count", Int32)
	Field(2, "last_increment_by", String)
	Field(3, "last_increment_at", String)
	Field(4, "next_finalize_at", String)
	Required("count", "last_increment_by", "last_increment_at", "next_finalize_at")
})

var _ = Service("api", func() {
	Security(JWTAuth, func() {
		Scope("api")
	})

	Error("unauthorized")
	Error("existing_increment_request", func() {
		Temporary()
	})

	HTTP(func() {
		Path("/api/v1")
		Response("unauthorized", StatusUnauthorized)
		Response("existing_increment_request", StatusTooManyRequests)
	})

	GRPC(func() {
		Response("unauthorized", CodePermissionDenied)
		Response("existing_increment_request", CodeAlreadyExists)
	})

	Method("AuthToken", func() {
		NoSecurity()

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
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("CounterGet", func() {
		Payload(func() {
			TokenField(1, "token", String)
			Required("token")
		})

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
		Payload(func() {
			TokenField(1, "token", String)
			Field(2, "user", String)
			Required("token", "user")
		})

		Result(CounterInfo)

		HTTP(func() {
			POST("/counter")
			Response(StatusAccepted)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Files("/openapi.json", "gen/http/openapi3.json")
})

var _ = Service("teapot", func() {
	Error("unwell")

	Method("Echo", func() {
		NoSecurity()

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

var _ = Service("web", func() {
	Error("unauthorized")

	HTTP(func() {
		Path("/")
		Response("unauthorized", StatusUnauthorized)
	})

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
				Cookie("session_cookie:countup.session")
				CookieSameSite(CookieSameSiteLax)
				CookieMaxAge(86400)
				CookieHTTPOnly()
				//TODO: CookieSecure()
				CookiePath("/")
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
			Param("code", String)
			Param("state", String)
			Cookie("session_cookie:countup.session")
			Response(StatusFound, func() {
				Header("redirect_url:Location", String)
				Cookie("session_cookie:countup.session")
				CookieSameSite(CookieSameSiteLax)
				CookieMaxAge(86400)
				CookieHTTPOnly()
				//TODO: CookieSecure()
				CookiePath("/")
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
			Cookie("session_cookie:countup.session")
			Response(StatusFound, func() {
				Header("redirect_url:Location", String)
				Cookie("session_cookie:countup.session")
				CookieSameSite(CookieSameSiteLax)
				CookieMaxAge(86400)
				CookieHTTPOnly()
				//TODO: CookieSecure()
				CookiePath("/")
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
			Cookie("session_cookie:countup.session")
			Response(StatusOK, func() {
				ContentType("application/json")
			})
		})
	})

	Files("/static/*", "static/")
})
