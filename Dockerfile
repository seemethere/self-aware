FROM golang:1.11

ADD . /build
WORKDIR /build
RUN go build -o main main.go

ENTRYPOINT "./main"
