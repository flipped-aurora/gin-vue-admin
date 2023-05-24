# 微信开放平台

[官方文档](https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Third_party_platform_appid.html)

## 快速入门

### 服务端处理

```go
wc := wechat.NewWechat()
memory := cache.NewMemory()
cfg := &openplatform.Config{
    AppID:     "xxx",
    AppSecret: "xxx",
    Token:     "xxx",
    EncodingAESKey: "xxx",
    Cache: memory,
}


openPlatform := wc.GetOpenPlatform(cfg)
// 传入request和responseWriter
server := openPlatform.GetServer(req, rw)
//设置接收消息的处理方法
server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
    if msg.InfoType == message.InfoTypeVerifyTicket {
        componentVerifyTicket, err := openPlatform.SetComponentAccessToken(msg.ComponentVerifyTicket)
        if err != nil {
            log.Println(err)
            return nil
        }
        //debug 
        fmt.Println(componentVerifyTicket)
        rw.Write([]byte("success"))
        return nil
    }
    //handle other message
    //
    
    
    return nil
})

//处理消息接收以及回复
err := server.Serve()
if err != nil {
    fmt.Println(err)
    return
}
//发送回复的消息
server.Send()


```

### 待授权处理消息

```go

//授权的第三方公众号的appID
appID := "xxx"
openPlatform := wc.GetOpenPlatform(cfg)
openPlatform.GetOfficialAccount(appID)

```