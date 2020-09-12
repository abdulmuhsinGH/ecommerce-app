# build stage
FROM golang:1.13.2-alpine3.10 as builder
RUN apk add --no-cache ca-certificates openssl

ENV GO111MODULE=on

RUN mkdir -p /ecormmerce-rest-api

WORKDIR /ecormmerce-rest-api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /ecormmerce-rest-api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build /ecormmerce-rest-api/cmd/ecommerce-api/main.go

# final stage
FROM scratch
COPY --from=builder /ecormmerce-rest-api/main /ecormmerce-rest-api
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
EXPOSE 8080
ENTRYPOINT ["/ecormmerce-rest-api"]
