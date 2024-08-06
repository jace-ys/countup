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
		Services("api", "web")
		Host("dev-http", func() {
			URI("http://localhost:8080")
		})
		Host("dev-grpc", func() {
			URI("grpc://localhost:8081")
		})
	})
})

var CounterInfo = ResultType("application/vnd.countup.counter-info`", "CounterInfo", func() {
	Field(1, "count", Int32)
	Field(2, "last_increment_by", String)
	Field(3, "last_increment_at", String)
	Field(4, "next_finalize_at", String)
	Required("count", "last_increment_by", "last_increment_at", "next_finalize_at")
})

var _ = Service("api", func() {
	Error("unauthorized")
	Error("existing_increment_request", func() {
		Temporary()
	})

	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
		Response("existing_increment_request", StatusTooManyRequests)
	})

	GRPC(func() {
		Response("unauthorized", CodePermissionDenied)
		Response("existing_increment_request", CodeAlreadyExists)
	})

	Method("CounterGet", func() {
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
			Field(1, "user", String)
			Required("user")
		})

		Result(CounterInfo)

		HTTP(func() {
			POST("/counter/inc")
			Response(StatusAccepted)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

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

var _ = Service("web", func() {
	Method("index", func() {
		Result(Bytes)
		HTTP(func() {
			GET("/")
			Response(StatusOK, func() {
				ContentType("text/html")
			})
		})
	})

	Method("another", func() {
		Result(Bytes)
		HTTP(func() {
			GET("/another")
			Response(StatusOK, func() {
				ContentType("text/html")
			})
		})
	})

	Files("/static/{*path}", "static/")
})
