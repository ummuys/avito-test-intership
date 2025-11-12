# BUILDER
FROM golang:1.25.1-alpine AS builder
WORKDIR /grant

COPY go.mod go.sum ./ 
COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN go mod download

RUN go mod tidy \
    && go build -o /grant/app ./cmd

# RUNNER
FROM alpine:3.22 AS runner
WORKDIR /root/
COPY --from=builder ./grant .
EXPOSE 1337
CMD ["./app"]