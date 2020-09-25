FROM golang:1.15.2

ADD ./ /go/src/aezion

WORKDIR /go/src/aezion

RUN go env
RUN go mod download
RUN go build ./cmd/todo

EXPOSE 8080

ENTRYPOINT ["./todo","server"]