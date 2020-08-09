# https://medium.com/statuscode/golang-docker-for-development-and-production-ce3ad4e69673
FROM golang:1.14.6-alpine3.12
LABEL maintainer="austincollinpena@gmail.com"
ENV REFRESHED_AT 7/27/2020/9:00

ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/github.com/austincollinpena/proxy-middleware
WORKDIR /go/src/github.com/austincollinpena/proxy-middleware

RUN go get github.com/pilu/fresh
RUN go get

RUN go get ./
RUN go build

# db025426d199