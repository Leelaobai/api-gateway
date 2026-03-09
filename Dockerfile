FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o gateway .

FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/gateway .
COPY etc/gateway-api.yaml etc/
EXPOSE 8080
CMD ["./gateway", "-f", "etc/gateway-api.yaml"]
