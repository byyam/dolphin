FROM golang:latest AS builder
RUN go version
MAINTAINER Yam <yanzhang.scut@gmail.com>

COPY . /go/src/github.com/byyam/dolphin/
WORKDIR /go/src/github.com/byyam/dolphin/
RUN make os=linux tcp-server

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/byyam/dolphin/bin/ ./bin
COPY --from=builder /go/src/github.com/byyam/dolphin/conf ./conf

EXPOSE 7800
ENTRYPOINT ["./bin/tcp-server"]
