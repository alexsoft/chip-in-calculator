FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o bin/chip-in-calculator ./

# Production stage
FROM scratch AS prod

LABEL org.opencontainers.image.source=https://github.com/alexsoft/chip-in-go

WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/bin/chip-in-calculator /chip-in-calculator

ENTRYPOINT ["/chip-in-calculator"]
