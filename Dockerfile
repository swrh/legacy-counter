# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o counter

# Final stage
FROM scratch

WORKDIR /app

COPY --from=builder /app/counter .

ENTRYPOINT ["./counter"]
