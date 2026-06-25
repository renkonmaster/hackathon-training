# syntax=docker/dockerfile:1.7

FROM node:24-alpine AS node-base
RUN corepack enable && corepack prepare pnpm@10.33.0 --activate
WORKDIR /workspace/client

FROM golang:1.25-alpine AS go-base
RUN apk add --no-cache git ca-certificates
WORKDIR /workspace

FROM node-base AS client-deps
COPY client/package.json client/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

FROM client-deps AS client-builder
COPY client ./
COPY docs ../docs
RUN pnpm run codegen && pnpm run build

FROM client-deps AS client-dev
CMD ["sh", "-c", "pnpm install --frozen-lockfile && pnpm run codegen && pnpm dev --host 0.0.0.0"]

FROM go-base AS server-dev
RUN go install github.com/air-verse/air@v1.63.0
WORKDIR /workspace/server
CMD ["air", "-c", ".air.toml"]

FROM go-base AS server-builder
COPY server/go.mod server/go.sum server/
WORKDIR /workspace/server
RUN go mod download
WORKDIR /workspace
COPY docs docs
COPY server server
WORKDIR /workspace/server
RUN go generate ./... && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -o /out/server ./cmd/api

FROM alpine:3.22 AS app
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=server-builder /out/server /app/server
COPY --from=client-builder /workspace/client/dist /app/assets
ENV API_ADDR=:8080
ENV ASSETS_DIR=/app/assets
ENTRYPOINT ["/app/server"]
