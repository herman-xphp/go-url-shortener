FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

# Allow Go toolchain to auto-download the required version
ENV GOTOOLCHAIN=auto
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env.example .env
COPY --from=builder /app/public ./public

EXPOSE 3000

CMD ["./main"]
