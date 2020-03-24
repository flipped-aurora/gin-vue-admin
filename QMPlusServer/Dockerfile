FROM golang:1.12 AS build

WORKDIR /go/src/gin-vue-admin
COPY . .

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
RUN go get ./... \
    && go build -o gin-vue-admin

FROM golang:1.12 AS serve

COPY --from=build /go/src/gin-vue-admin /go/src/gin-vue-admin

EXPOSE 8080
ENTRYPOINT [ "./gin-vue-admin" ]
