FROM golang:alpine as builder

# 设置go mod proxy 国内代理
# 设置golang path
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1
WORKDIR /ginvue
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
COPY . .
RUN go env && go list && go build -o app main.go

EXPOSE 8888
ENTRYPOINT /ginvue/app

# 根据Dockerfile生成Docker镜像
# docker build -t ginvue .
# 根据Docker镜像启动Docker容器
# docker run -itd -p 8888:8888 --name ginvue ginvue
