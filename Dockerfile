# Stage 1: Build the Go binaries
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod .
# COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/function-waiter ./cmd/function-waiter
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/function-sequential ./cmd/entrypoint-sequential
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/function-parallel ./cmd/entrypoint-parallel


# Stage 2: Create a minimal image for each function

FROM alpine:latest AS neptuneplus-simple-dependencies-waiter
EXPOSE 8080
WORKDIR /app

COPY --from=builder /app/bin/function-waiter .
ENTRYPOINT ["./function-waiter"]

FROM alpine:latest AS neptuneplus-simple-dependencies-sequential
EXPOSE 8080
WORKDIR /app

COPY --from=builder /app/bin/function-sequential .
ENTRYPOINT ["./function-sequential"]

FROM alpine:latest AS neptuneplus-simple-dependencies-parallel
EXPOSE 8080
WORKDIR /app

COPY --from=builder /app/bin/function-parallel .
ENTRYPOINT ["./function-parallel"]