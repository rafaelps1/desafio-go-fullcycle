FROM golang:1.21-alpine AS builder

WORKDIR /src

COPY . .

RUN go mod tidy

EXPOSE 8080
