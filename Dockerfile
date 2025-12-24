FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o chip-in-calculator ./

# Production stage
FROM alpine:3 AS prod

WORKDIR /prod

COPY --from=builder /build/chip-in-calculator ./chip-in-calculator
COPY --from=builder /build/config.json ./config.json

ENTRYPOINT ["/prod/chip-in-calculator"]
