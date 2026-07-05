FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN go test ./... && go build -o /mobility-api ./cmd/server

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /mobility-api /usr/local/bin/mobility-api
EXPOSE 8080
CMD ["mobility-api"]
