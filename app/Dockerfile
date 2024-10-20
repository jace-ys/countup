FROM golang:1.23.2 AS builder

ARG LDFLAGS
ENV LDFLAGS=${LDFLAGS}

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -ldflags="${LDFLAGS}" -o dist/ ./cmd/countup/...

FROM scratch
WORKDIR /app
COPY --from=builder /src/dist /app/bin
ENTRYPOINT ["/app/bin/countup"]