FROM golang:latest AS builder
MAINTAINER Yam <yanzhang.scut@gmail.com>
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN make os=linux http-server
EXPOSE 12000
ENTRYPOINT "/app/http-server -recipe /app/conf/conf.toml"
