# Stage 1: Build the Go binary
FROM golang:1.25.1-bookworm AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o homies-api ./cmd/homies/main.go
 
# Stage 2: Create minimal final image
FROM debian:bookworm-slim
WORKDIR /app

COPY --from=builder /app/homies-api .

EXPOSE 8080

CMD [ "./homies-api" ]
