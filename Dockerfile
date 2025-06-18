# -------- Stage 1: Build the Go binary --------
FROM golang:1.23-alpine AS builder

WORKDIR /build

# Install Goose in builder stage (just for copying into final image)
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./playgobackend ./cmd/main/main.go

# -------- Stage 2: Final runtime image --------
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache jq

# Install Goose runtime binary
COPY --from=builder /go/bin/goose /usr/local/bin/goose

# Copy built Go backend
COPY --from=builder /build/playgobackend /app/playgobackend

# Copy migrations folder
COPY ./db/migrations /migrations

# Entrypoint script
COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]


