FROM golang:1.23.6-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]