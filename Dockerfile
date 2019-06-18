# builder
FROM golang:1.12 AS builder

WORKDIR /go/src/mao-ctrl

ENV GO111MODULE=on

RUN groupadd -g 10001 mao-ctrl \
    && useradd -u 10001 -g mao-ctrl mao-ctrl

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install \
    -ldflags="-w -s" \
    ./cmd/mao-ctrl

# runtime image
FROM scratch

COPY --from=builder /go/bin/mao-ctrl /mao-ctrl
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 3000

USER mao-ctrl

ENTRYPOINT ["/mao-ctrl"]
