# build stage
FROM golang:1.13.2-alpine3.10 as builder

ENV GO111MODULE=on

RUN mkdir -p /auth-server

WORKDIR /auth-server

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /authserver

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth-server/bin/auth-server /auth-server/cmd/auth/main.go

# final stage
FROM scratch
COPY --from=builder /auth-server/bin/auth-server /auth-server
EXPOSE 9096
ENTRYPOINT ["/auth-server"]
