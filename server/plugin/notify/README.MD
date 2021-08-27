## GVA 钉钉群通知插件

本插件用于向钉钉群推送消息

### 1. 使用场景

- 当服务运行异常时，可以向钉钉推送异常信息，便于及时发现解决问题
- 推送一些关键业务的运行日志等

### 2. 配置说明

钉钉 token 等相关信息的获取，请参考 [钉钉官网](https://developers.dingtalk.com/document/robots/custom-robot-access?spm=ding_open_doc.document.0.0.7f8710afbfzduV#topic-2026027)

在`plugin/notify/global/global.go` 文件中配置钉钉通知的URL ，Token 等

```go
	//  在gin-vue-admin 主程序的initialize中的plugin的InstallPlugin 函数中写入如下代码
    PluginInit(PublicGroup, notify.CreateDDPlug(
        URL,
        Token,
        密钥))
}
```
### 3方法API
```go

可调用方法
//content 发送的内容
//atMobiles 需要艾特的人的手机号 
//isAtAll 是否艾特全体
SendTextMessage(content string,atMobiles []string,isAtAll bool)

//content 发送的内容
//title 内容标题
//picUrl 配图
//messageUrl 点击跳转路径
//atMobiles 需要艾特的人的手机号 
//isAtAll 是否艾特全体
SendLinkMessage(content,title,picUrl,messageUrl string)

//content 发送的内容(markdown语法)
//title 内容标题
//atMobiles 需要艾特的人的手机号 
//isAtAll 是否艾特全体
SendMarkdownMessage(content,title string,atMobiles []string,isAtAll bool)

```

### 4. 使用示例

在代码中调用 `SendTextMessage` 方法即可
```go
// 示例
func NotifyController(c *gin.Context) {
    if err := service.ServiceGroupApp.SendTextMessage("test",[]string{"手机号"},是否艾特全体); err != nil {
        global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
        response.FailWithMessage("发送失败", c)
    } else {
        response.OkWithData("发送成功", c)
    }
}

```


