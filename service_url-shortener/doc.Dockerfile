#
#######################################################################################
#
# FROM ericgln/godoc

## FROM golang as builder
FROM golang:alpine as builder
#
LABEL maintainer="Franz Branntvino"
LABEL Name=doc-server
LABEL Version=0.0.1
#
#######################################################################################
#
WORKDIR /build

RUN go get golang.org/x/tools/cmd/godoc

COPY . .
#RUN go mod tidy

EXPOSE 6060
ENTRYPOINT /go/bin/godoc -http=:6060
