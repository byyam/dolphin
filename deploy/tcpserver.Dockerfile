FROM golang:latest AS builder
RUN go version
MAINTAINER Yam <yanzhang.scut@gmail.com>

COPY . /go/src/github.com/byyam/dolphin/
WORKDIR /go/src/github.com/byyam/dolphin/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -o app ./service/tcp-server

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/byyam/dolphin/app .
COPY --from=builder /go/src/github.com/byyam/dolphin/conf ./conf

EXPOSE 7800
ENTRYPOINT ["./app"]
