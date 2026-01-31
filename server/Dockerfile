FROM golang:1.21-alpine AS builder

WORKDIR /app

# Dependencies
RUN apk add --no-cache git

# Copy source code
COPY . .

# Build
# Clean go.mod/sum and rebuild to avoid checksum issues
RUN go clean -modcache && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server/main.go

# Final stage
FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]EXPOSE 8080

CMD ["./server"]
