FROM golang:1.19 AS development
WORKDIR /go/src/github.com/escalopa/goblog
COPY . .
RUN go mod download
RUN go install github.com/cespare/reflex@latest
CMD reflex -sr '\.go$' go run ./cmd/main.go

FROM golang:1.19.3-alpine3.16 AS builder
WORKDIR /go/src/github.com/escalopa/goblog
COPY . .
RUN go build -o /go/bin/main ./cmd/main.go

FROM alpine:3.16 AS production
COPY --from=builder /go/bin/main /go/bin/main

EXPOSE 8000
CMD ["/go/bin/main"]