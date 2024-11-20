# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install git and dependencies
RUN apk add --no-cache git

# Копируем только необходимые файлы
COPY proto/go.mod proto/go.mod
COPY proto/pb proto/pb
COPY user-service user-service

RUN echo 'go 1.23.3' > go.work && \
    echo '' >> go.work && \
    echo 'use (' >> go.work && \
    echo '    /app/user-service' >> go.work && \
    echo '    /app/proto' >> go.work && \
    echo ')' >> go.work

# Активируем workspace mode
ENV GOWORK=/app/go.work

# Download dependencies
WORKDIR /app/user-service
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service cmd/api/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/user-service/user-service .

# Expose port
EXPOSE 8082 9092

# Command to run
CMD ["./user-service"]