# 小程序-云开发 SDK

Tencent Cloud Base [文档](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/)

## 使用说明

**初始化配置**

```golang
//使用memcache保存access_token，也可选择redis或自定义cache
memCache=cache.NewMemcache("127.0.0.1:11211")

//配置小程序参数
config := &wechat.Config{
    AppID:     "your app id",
    AppSecret: "your app secret",
    Cache:     memCache,
}
wc := wechat.NewWechat(config)
wcTcb := wc.GetTcb()
```

### 举例

#### 触发云函数

```golang
res, err := wcTcb.InvokeCloudFunction("test-xxxx", "add", `{"a":1,"b":2}`)
if err != nil {
    panic(err)
}
```

更多使用方法参考[PKG.DEV](https://pkg.go.dev/github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/tcb)