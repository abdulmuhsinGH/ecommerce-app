# build stage
FROM golang:1.13.2-alpine3.10 as builder

ENV GO111MODULE=on

RUN mkdir -p /auth-server

WORKDIR /auth-server

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /auth-server

# RUN rm -r vendor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build /ecormmerce-app/auth-server/cmd/auth/main.go

# final stage
FROM scratch
COPY --from=builder /ecormmerce-app/auth-server/main /auth-server
COPY ./view/ /view/
EXPOSE 9096
ENTRYPOINT ["/auth-server"]
