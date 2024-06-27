FROM golang:1.22.4-bullseye AS base
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod download

FROM base AS dev
RUN --mount=type=cache,target=/root/.cache/go-build \
    go install github.com/air-verse/air@latest
# RUN go install github.com/air-verse/air@latest
COPY . .
CMD ["air", "-c", ".air.toml"]

FROM base AS builder
WORKDIR /usr/src/app
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    GOOS=linux CGO_ENABLED=0 go build \
    -ldflags="-s -w" \
    -trimpath \
    -o main ./src/cmd/api/main.go

FROM debian:bullseye-slim AS runner
COPY --from=builder /usr/src/app/main /usr/src/app/main
ENTRYPOINT ["/usr/src/app/main"]