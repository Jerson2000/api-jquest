# Build Stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install dependencies for building if necessary (e.g. git for private repos, though not needed here really)
# RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
# CGO_ENABLED=0 for static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Run Stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .
COPY --from=builder /app/.env.example .env
COPY --from=builder /app/casbin ./casbin

# Expose port (default 8080 usually)
EXPOSE 8080

CMD ["./main"]
