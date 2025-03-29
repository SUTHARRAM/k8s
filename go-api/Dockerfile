# Build stage 
FROM golang:1.20 as builder
WORKDIR /app
COPY . . 
RUN go mod download 
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Run stage 
FROM alpine:latest 
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]