# Multistage docker file
# Base
FROM golang:1.26-alpine AS base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Development
FROM base as development

ARG API_PORT

EXPOSE ${API_PORT}

CMD ["go", "tool", "air"]

# Builder
FROM base AS builder

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags="-s -w" -o api ./cmd/api/

# Production
FROM alpine:3.23 AS production

ARG API_PORT

WORKDIR /

COPY --from=builder /app/api .

EXPOSE ${API_PORT}

CMD ["./api"]
