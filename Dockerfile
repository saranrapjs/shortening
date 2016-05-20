FROM golang:1.6

RUN mkdir -p /go/src/github.com/saranrapjs/shortening
WORKDIR /go/src/github.com/saranrapjs/shortening
ADD . /go/src/github.com/saranrapjs/shortening

RUN go get ./...
RUN go build .

EXPOSE 8080

ENTRYPOINT ["shortening"]