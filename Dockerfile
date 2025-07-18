# -------- Stage 1: Build --------
FROM golang:1.23-alpine AS builder

WORKDIR /build

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./playgobackend ./cmd/main/main.go

# -------- Stage 2: Runtime --------
FROM golang:1.23-alpine AS runtime

WORKDIR /app

RUN apk add --no-cache jq

COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /build/playgobackend /app/playgobackend
COPY ./db/migrations /migrations
COPY ./entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh
EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
