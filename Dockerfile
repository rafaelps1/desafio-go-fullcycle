FROM golang:1.21-alpine AS builder

WORKDIR /src

# Restore dependencies
COPY . .

RUN go mod tidy

# Build executable
# RUN go build -o /src/route ./

# FROM scratch:latest
# WORKDIR /src
# COPY --from=builder /src/route ./
EXPOSE 8080
