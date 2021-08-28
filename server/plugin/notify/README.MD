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

### 3 参数说明
#### 3-1 全局参数说明

```go
	Url    string `mapstructure:"url" json:"url" yaml:"url"`          // Url
	Token  string `mapstructure:"token" json:"token" yaml:"token"`    // access_token
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"` // 密钥
```
#### 3-2 请求入参说明
```go


```

### 3方法API（可调用方法）
```go

//content 发送的内容
//atMobiles 需要艾特的人的手机号 
//isAtAll 是否艾特全体
SendTextMessage(content string,atMobiles []string,isAtAll bool)

//content 发送的内容
//title 内容标题
//picUrl 配图
//messageUrl 点击跳转路径
SendLinkMessage(content,title,picUrl,messageUrl string)

//content 发送的内容(markdown语法)
//title 内容标题
//atMobiles 需要艾特的人的手机号 
//isAtAll 是否艾特全体
SendMarkdownMessage(content,title string,atMobiles []string,isAtAll bool)

```

### 4. 可直接调用接口

    发送文字消息接口： /notify/sendTextMessage [post] 已配置swagger
    发送图文链接消息接口： /notify/sendLinkMessage [post] 已配置swagger
    发送markdown消息接口： /notify/sendMarkdownMessage [post] 已配置swagger

    入参：
    type Email struct {
        To      string `json:"to"`      // 邮件发送给谁
        Subject string `json:"subject"` // 邮件标题
        Body    string `json:"body"`    // 邮件内容
    }


