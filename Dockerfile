# Production builder
FROM golang:1.12 as builder

WORKDIR /go/src/github.com/mao-wfs/mao-ctrl

COPY . .

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v \
    -ldflags="-w -s" \
    github.com/mao-wfs/mao-ctrl/cmd/mao-ctrl

# Production image
FROM scratch

COPY --from=builder /go/bin/mao-ctrl /mao-ctrl

EXPOSE 3000

CMD ["/mao-ctrl"]
