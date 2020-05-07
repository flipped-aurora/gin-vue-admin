FROM golang:alpine as builder
RUN  apk add --update --no-cache yarn make g++
RUN yarn global add cross-env node-sass

ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1
WORKDIR /go/src/gin-vue-admin
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
COPY server/ ./
RUN go env && go list && go build -v -a -ldflags "-extldflags \"-static\" " -o gvadmin .

WORKDIR /web
COPY web/ ./
RUN yarn install && yarn run build


FROM nginx:alpine
LABEL MAINTAINER="rikugun"

RUN apk add --no-cache  gettext tzdata   && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" >  /etc/timezone && \
    date && \
    apk del tzdata

COPY docker/etc/nginx/nginx.conf.tpl /etc/nginx/nginx.conf.tpl
WORKDIR /app
#copy web
COPY --from=builder /web/dist/ /var/www/
#copy go app
COPY --from=builder /go/src/gin-vue-admin/gvadmin ./
COPY --from=builder /go/src/gin-vue-admin/db.db ./
COPY --from=builder /go/src/gin-vue-admin/config.yaml ./
COPY --from=builder /go/src/gin-vue-admin/resource ./resource
COPY docker/docker-start.sh ./

ENV API_SERVER="http://localhost:8888/"
EXPOSE 80

ENTRYPOINT ["./docker-start.sh"]
