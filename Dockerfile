FROM golang:1.11-alpine

WORKDIR /go/src/github.com/0daryo/cinemawith/
COPY . .
ENV GO111MODULE=on

RUN apk --no-cache --update upgrade \
  && apk add --no-cache git alpine-sdk \
  && go get -u github.com/codegangsta/gin

CMD gin -i run main.go routing.go dependency.go