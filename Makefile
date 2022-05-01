SHELL = /bin/bash

CONFIG_FILE         = config.yaml
PROJECT_NAME        = github.com/flipped-aurora/gin-vue-admin/server

#SCRIPT_DIR         = $(shell pwd)/etc/script
BUILD_IMAGE_SERVER  = golang:1.16
BUILD_IMAGE_WEB     = node:16
IMAGE_NAME          = all
REPOSITORY          = gin-vue-admin

# support build custom tags
ifneq ($(TAGS),)
TAGS_OPT 	    = -tags ${TAGS}
endif

#前后端共同打包
build: build-web build-server
	docker run --name build-local --rm -v $(shell pwd):/go/src/${PROJECT_NAME} -w /go/src/${PROJECT_NAME} ${BUILD_IMAGE_SERVER} make build-local

#容器打包前端
build-web:
	docker run --name build-web-local --rm -v $(shell pwd):/go/src/${PROJECT_NAME} -w /go/src/${PROJECT_NAME} ${BUILD_IMAGE_WEB} make build-web-local

#容器打包后端
build-server:
	docker run --name build-server-local --rm -v $(shell pwd):/go/src/${PROJECT_NAME} -w /go/src/${PROJECT_NAME} ${BUILD_IMAGE_SERVER} make build-server-local

#构建web镜像
build-image-web:
	@cd web/ && docker build --rm -t ${REPOSITORY}/web:${TAGS_OPT} .

#构建server镜像
build-image-server:
	@cd server/ && docker build --rm -t ${REPOSITORY}/server:${TAGS_OPT} .

#本地环境打包前后端
build-local:
	if [ -d "build" ];then rm -rf build; else echo "OK!"; fi \
	&& if [ -f "/.dockerenv" ];then echo "OK!"; else  make build-web-local && make build-server-local; fi \
	&& mkdir build && cp -r web/dist build/ && cp server/server build/ && cp -r server/resource build/resource 

#本地环境打包前端
build-web-local:
	@cd web/ && if [ -d "dist" ];then rm -rf dist; else echo "OK!"; fi \
	&& yarn config set registry http://mirrors.cloud.tencent.com/npm/ \
	&& yarn install && yarn build

#本地环境打包后端
build-server-local:
	@cd server/ && if [ -f "server" ];then rm -rf server; else echo "OK!"; fi \
	&& go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct \
	&& go env -w CGO_ENABLED=0 && go env  && go mod tidy && go build

build-web-local-action:
	@cd web/ && if [ -d "dist" ];then rm -rf dist; else echo "OK!"; fi \
	&& yarn install && yarn build

build-server-local-action:
	@cd server/ && if [ -f "server" ];then rm -rf server; else echo "OK!"; fi \
	&& go mod tidy && go build
#后端打包待优化版
build-server-locallll:
	GO111MODULE=on CGO_ENABLED=1 go build ${TAGS_OPT} \
		-ldflags "-B 0x$(shell head -c20 /dev/urandom|od -An -tx1|tr -d ' \n') -X main.Version=${MAJOR_VERSION}(${GIT_VERSION})" \
		-v -o ${TARGET} 
#镜像待优化自动化版本
images: build
	docker build -t registry.cn-hangzhou.aliyuncs.com/tscuite/gva:v1 .

