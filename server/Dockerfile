FROM golang:alpine as builder

WORKDIR /gva-server
COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy