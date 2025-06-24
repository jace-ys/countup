FROM golang:1.24.3@sha256:81bf5927dc91aefb42e2bc3a5abdbe9bb3bae8ba8b107e2a4cf43ce3402534c6 AS builder

ARG LDFLAGS
ENV LDFLAGS=${LDFLAGS}

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -ldflags="${LDFLAGS}" -o dist/ ./cmd/countup/...

FROM gcr.io/distroless/static-debian12@sha256:b7b9a6953e7bed6baaf37329331051d7bdc1b99c885f6dbeb72d75b1baad54f9
WORKDIR /app
USER nonroot:nonroot
COPY --from=builder --chown=nonroot:nonroot /src/dist /app/bin
ENTRYPOINT ["/app/bin/countup"]