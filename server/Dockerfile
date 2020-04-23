FROM centos:7.6.1810

# 设置go mod proxy 国内代理
# 设置golang path
ENV GOPROXY=https://goproxy.io GOPATH=/gopath PATH="${PATH}:/usr/local/go/bin"
# 定义使用的Golang 版本
ARG GO_VERSION=1.13.3

# 安装 golang 1.13.3
RUN yum install -y wget && \
    yum install -y wqy-microhei-fonts wqy-zenhei-fonts && \
    wget "https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz" && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf "go$GO_VERSION.linux-amd64.tar.gz" && \
    rm -rf *.tar.gz && \
    go version && go env;


WORKDIR $GOPATH
COPY . gin-vue

RUN cd server && go build -o app;

EXPOSE 8888

CMD ["gin-vue/app"]