# build stage
FROM golang:1.13.2-alpine3.10 as builder
RUN apk add --no-cache ca-certificates openssl

ENV GO111MODULE=on

RUN mkdir -p /auth-server

WORKDIR /auth-server

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /auth-server

# RUN rm -r vendor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build /auth-server/cmd/auth/main.go

# final stage
FROM scratch
COPY --from=builder /auth-server/main /auth-server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY ./view/ /view/
# RUN apk add --no-cache ca-certificates openssl

EXPOSE 8080

ENTRYPOINT ["/auth-server"]
