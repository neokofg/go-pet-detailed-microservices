# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install git and dependencies
RUN apk add --no-cache git

# Копируем только необходимые файлы
COPY proto/go.mod proto/go.mod
COPY proto/pb proto/pb
COPY api-gateway api-gateway

RUN echo 'go 1.23.3' > go.work && \
    echo '' >> go.work && \
    echo 'use (' >> go.work && \
    echo '    /app/api-gateway' >> go.work && \
    echo '    /app/proto' >> go.work && \
    echo ')' >> go.work

# Активируем workspace mode
ENV GOWORK=/app/go.work

# Download dependencies
WORKDIR /app/api-gateway
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway cmd/api/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/api-gateway/api-gateway .

# Expose port
EXPOSE 8080

# Command to run
CMD ["./api-gateway"]