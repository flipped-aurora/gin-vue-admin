## GVA 钉钉群通知插件

本插件用于向钉钉群推送消息

### 1. 使用场景

- 当服务运行异常时，可以向钉钉推送异常信息，便于及时发现解决问题
- 推送一些关键业务的运行日志等

### 2. 配置说明

钉钉 token 等相关信息的获取，请参考 [钉钉官网](https://developers.dingtalk.com/document/robots/custom-robot-access?spm=ding_open_doc.document.0.0.7f8710afbfzduV#topic-2026027)

在`plugin/notify/global/global.go` 文件中配置钉钉通知的URL ，Token 等

```go
var GlobalConfig_ = &config.DingDing{
	Url:    "https://oapi.dingtalk.com/robot/send",
	Token:  "xxx",
	Secret: "xxx",
}
```
### 3. 使用说明

在代码中调用 `SendTextMessage` 方法即可
```go
func NotifyController(c *gin.Context) {
    if err := service.ServiceGroupApp.SendTextMessage("test"); err != nil {
        global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
        response.FailWithMessage("发送失败", c)
    } else {
        response.OkWithData("发送成功", c)
    }
}

```
### 方法API

