FROM golang:1.25-alpine3.21 as builder

WORKDIR app
COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /api ./cmd/api

FROM alpine:3.21
RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /api /app/api

RUN addgroup -S appgroup && adduser -S appuser -G appgroup && chown - R appuser:appgroup /app
USER appuser

CMD ["/app/api"]
