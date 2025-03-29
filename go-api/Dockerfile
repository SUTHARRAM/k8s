# Build stage 
FROM golang:1.20 AS builder
WORKDIR /app
COPY . . 
RUN go mod download 
RUN go mod tidy  # Ensure go.mod is tidy and includes all dependencies
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Run stage 
FROM alpine:latest 
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]