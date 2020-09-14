# build stage
FROM golang:1.13.2-alpine3.10 as builder
RUN apk add --no-cache ca-certificates openssl

ENV GO111MODULE=on

RUN mkdir -p /migration

WORKDIR /migration

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /migration

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build /migration/main.go

# final stage
FROM scratch
COPY --from=builder /migration/main /migration_cli
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/migration_cli", "up"]
