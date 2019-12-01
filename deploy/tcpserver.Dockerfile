FROM golang:latest AS builder
MAINTAINER Yam <yanzhang.scut@gmail.com>
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN make os=linux tcp-server
EXPOSE 7800
ENTRYPOINT "/app/tcp-server -recipe /app/conf/conf.toml"
