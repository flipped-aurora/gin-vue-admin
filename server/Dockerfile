FROM golang:1.18-alpine as build-stage

WORKDIR /gva-server

ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o app .

FROM alpine:latest

LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /gva-server

COPY --from=build-stage /gva-server/app ./
COPY --from=build-stage /gva-server/resource ./resource/
COPY --from=build-stage /gva-server/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ./app -c config.docker.yaml
