FROM golang:1.24.6@sha256:8d9e57c5a6f2bede16fe674c16149eee20db6907129e02c4ad91ce5a697a4012 AS builder

ARG LDFLAGS
ENV LDFLAGS=${LDFLAGS}

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -ldflags="${LDFLAGS}" -o dist/ ./cmd/countup/...

FROM gcr.io/distroless/static-debian12@sha256:4b2a093ef4649bccd586625090a3c668b254cfe180dee54f4c94f3e9bd7e381e
WORKDIR /app
USER nonroot:nonroot
COPY --from=builder --chown=nonroot:nonroot /src/dist /app/bin
ENTRYPOINT ["/app/bin/countup"]
