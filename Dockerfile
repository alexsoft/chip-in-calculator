FROM golang:1.26-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o chip-in-calculator ./

# Production stage
FROM scratch AS prod

LABEL org.opencontainers.image.source=https://github.com/alexsoft/chip-in-go

WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/chip-in-calculator /chip-in-calculator

ENTRYPOINT ["/chip-in-calculator"]
