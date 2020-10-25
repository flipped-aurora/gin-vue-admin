.PHONY: all build run gotool clean help

BINARY="server"
GVA = "gva"

all: check gva initdb run

gva:
	go build -o ${GVA} cmd/main.go

initdb:
	@if [ -f ${GVA} ] ; then ./gva initdb && rm ${GVA} ; fi

linux-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

windows-build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe

mac-build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BINARY}

run:
	@go run main.go

check:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	@if [ -f ${GVA} ] ; then rm ${GVA} ; fi

help:
	@echo "make - 构建gva终端工具并初始化数据,初始化数据后删除gva终端工具,启动server项目"
	@echo "make gva - 构建gva终端工具"
	@echo "make initdb - 如果有gva终端工具就初始化数据,初始化数据后删除gva终端工具,没有则不会执行"
	@echo "make linux-build - 编译 Go 代码, 生成Linux系统的二进制文件"
	@echo "make windows-build - 编译 Go 代码, 生成Windows系统的exe文件"
	@echo "make mac-build - 编译 Go 代码, 生成Mac系统的二进制文件"
	@echo "make run - 直接运行 main.go"
	@echo "make clean - 移除二进制文件"
	@echo "make check - 运行 Go 工具 'fmt' and 'vet'"