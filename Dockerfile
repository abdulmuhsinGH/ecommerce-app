FROM golang:1.13.2-alpine3.10

RUN mkdir -p go/src/ecormmerce-rest-api

WORKDIR /go/src/ecormmerce-rest-api

COPY . /go/src/ecormmerce-rest-api

RUN go install ecormmerce-rest-api/cmd/ecommerce-api

CMD /go/bin/ecommerce-api

EXPOSE 8080