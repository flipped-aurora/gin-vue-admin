FROM golang:1.12-alpine

WORKDIR /go/src/gin-vue-admin
COPY . .

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn
RUN go mod download

RUN go build -ldflags="-s -w" -o gin-vue-admin
ENTRYPOINT [ "./gin-vue-admin" ]
