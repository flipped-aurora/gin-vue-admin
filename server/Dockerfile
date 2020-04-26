FROM golang:alpine as builder
RUN apk add --update --no-cache yarn make g++

ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1
WORKDIR /go/src/gin-vue-admin
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
COPY . .
RUN go env && go list && go build -v -a -ldflags "-extldflags \"-static\" " -o gvadmin .

WORKDIR /app
COPY --from=builder /go/src/gin-vue-admin/gvadmin .
COPY --from=builder /go/src/gin-vue-admin/db.db .
COPY --from=builder /go/src/gin-vue-admin/config.yaml .
COPY --from=builder /go/src/gin-vue-admin/resource ./resource

EXPOSE 8888

CMD ["gvadmin/app"]

