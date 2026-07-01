FROM golang:1.25-bookworm AS builder

ARG GOPROXY=https://goproxy.cn,direct

ENV CGO_ENABLED=0 \
    GOPROXY=${GOPROXY}

WORKDIR /src

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags="-s -w" -o /out/tdx-api ./api

FROM alpine:3.22

ENV GIN_MODE=release \
    TDX_API_ADDR=:8181 \
    TZ=Asia/Shanghai

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata wget \
    && mkdir -p /app/data/database

COPY --from=builder /out/tdx-api /app/tdx-api

VOLUME ["/app/data/database"]

EXPOSE 8181

CMD ["/app/tdx-api"]
