FROM golang:1.19.5-buster

WORKDIR /go/src/app

ADD . /go/src/app/

RUN go mod download

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz && mv migrate.linux-amd64 migrate
