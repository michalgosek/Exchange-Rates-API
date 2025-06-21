FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
 
COPY . .

RUN go build -o micgos-web-app ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/micgos-web-app .

EXPOSE 3000

CMD ["./micgos-web-app"]