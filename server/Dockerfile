FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /go/src/gin-vue-admin
COPY . .
RUN go env && go build -o server .

FROM alpine:latest
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/gin-vue-admin
COPY --from=0 /go/src/gin-vue-admin/server ./
COPY --from=0 /go/src/gin-vue-admin/config.yaml ./
COPY --from=0 /go/src/gin-vue-admin/resource ./resource

EXPOSE 8888

ENTRYPOINT ./server
