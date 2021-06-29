# Global build ARGs
ARG platform=x86-64
ARG system=linux

################################################
# STEP 1 - Build Go app
################################################
FROM golang:alpine as builder

WORKDIR /build

# ADD GO PACKAGE DEPENDENCIES HERE
RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/go-redis/redis/v8

COPY . .

RUN go mod tidy

# COMPILE GO APPLICATION
RUN CGO_ENABLED=0 GOOS=${system} go build -o ./app -v ./main.go

#########################################
# STEP 2 - Build image for Go app
#########################################
FROM scratch
#
LABEL maintainer="Franz Branntvino"
LABEL Name=urlShortener-demo
LABEL Version=0.0.1
#
WORKDIR /app

COPY --from=builder /build/app .
ADD static-data ./static-data

EXPOSE 8080

CMD ["./app"]
