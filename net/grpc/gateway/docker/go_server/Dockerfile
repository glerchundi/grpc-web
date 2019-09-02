FROM golang:1.13rc2

ENV CGO_ENABLED 0
ENV GO111MODULE on

COPY . /go/src/github.com/grpc/grpc-web
RUN cd /go/src/github.com/grpc/grpc-web \
    && go mod init \
    && go build ./... \
    && go install github.com/grpc/grpc-web/net/grpc/gateway/examples/echo/go-server

EXPOSE 9090
CMD ["go-server"]
