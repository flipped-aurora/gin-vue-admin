企业微信会话存档SDK（基于企业微信C版官方SDK封装），暂时只支持在`linux`环境下使用当前SDK。


### 官方文档地址
https://open.work.weixin.qq.com/api/doc/90000/90135/91774

### 使用方式

1、安装 go module
> go get -u github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2

2、从 `github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/msgaudit/lib` 文件夹下复制 `libWeWorkFinanceSdk_C.so` 动态库文件到系统动态链接库默认文件夹下，或者复制到任意文件夹并在当前文件夹下执行 `export LD_LIBRARY_PATH=$(pwd)`命令设置动态链接库检索地址后即可正常使用

3、编译要求
- 开启CGO: `CGO_ENABLED=1`
- 增加tags参数`msgaudit`： `go build -tags msgaudit`或者`go run -tags msgaudit main.go`

### Example

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/msgaudit"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/config"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	//初始化客户端
	wechatClient := wechat.NewWechat()

	workClient := wechatClient.GetWork(&config.Config{
		CorpID:        "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		CorpSecret:    "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		RasPrivateKey: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	})

	client, err := workClient.GetMsgAudit()
	if err != nil {
		fmt.Printf("SDK 初始化失败：%v \n", err)
		return
	}

	//同步消息
	chatDataList, err := client.GetChatData(0, 100, "", "", 3)
	if err != nil {
		fmt.Printf("消息同步失败：%v \n", err)
		return
	}

	for _, chatData := range chatDataList {
		//消息解密
		chatInfo, err := client.DecryptData(chatData.EncryptRandomKey, chatData.EncryptChatMsg)
		if err != nil {
			fmt.Printf("消息解密失败：%v \n", err)
			return
		}

		if chatInfo.Type == "image" {
			image, _ := chatInfo.GetImageMessage()
			sdkFileID := image.Image.SdkFileID

			isFinish := false
			buffer := bytes.Buffer{}
			indexBuf := ""
			for !isFinish {
				//获取媒体数据
				mediaData, err := client.GetMediaData(indexBuf, sdkFileID, "", "", 5)
				if err != nil {
					fmt.Printf("媒体数据拉取失败：%v \n", err)
					return
				}
				buffer.Write(mediaData.Data)
				if mediaData.IsFinish {
					isFinish = mediaData.IsFinish
				}
				indexBuf = mediaData.OutIndexBuf
			}
			filePath, _ := os.Getwd()
			filePath = path.Join(filePath, "test.png")
			err := ioutil.WriteFile(filePath, buffer.Bytes(), 0666)
			if err != nil {
				fmt.Printf("文件存储失败：%v \n", err)
				return
			}
			break
		}
	}

	//释放SDK实例
	client.Free()
}



```