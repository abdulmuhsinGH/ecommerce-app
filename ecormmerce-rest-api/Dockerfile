# build stage
FROM golang:1.13.2-alpine3.10 as builder

ENV GO111MODULE=on

RUN mkdir -p /ecormmerce-rest-api

WORKDIR /ecormmerce-rest-api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /ecormmerce-app/ecormmerce-rest-api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build /ecormmerce-app/ecormmerce-rest-api/cmd/ecommerce-api/main.go

# final stage
FROM scratch
COPY --from=builder /ecormmerce-app/ecormmerce-rest-api/main /ecormmerce-rest-api
EXPOSE 8081
ENTRYPOINT ["/ecormmerce-rest-api"]
