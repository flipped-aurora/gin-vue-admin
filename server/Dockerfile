# stage 1: build
FROM golang:1.18-alpine as build-stage

WORKDIR /build

ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o server .

# stage 2: distribute
FROM alpine:latest

LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /app

COPY --from=build-stage /build/server ./
COPY --from=build-stage /build/resource ./resource/
COPY --from=build-stage /build/config.docker.yaml ./

EXPOSE 8888

ENTRYPOINT ./server -c config.docker.yaml
