# ========== STAGE 1: Build ==========
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod dan go.sum (biar caching dependensi)
COPY go.mod go.sum ./
RUN go mod download

# Copy semua kode
COPY . .

# Build binary (statis, biar bisa jalan di alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# ========== STAGE 2: Production ==========
FROM alpine:latest

# Install ca-certificates (buat HTTPS request)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary dari stage builder
COPY --from=builder /app/main .

EXPOSE 8081

CMD ["./main"]