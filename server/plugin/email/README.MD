## GVA 邮件发送功能插件
#### 开发者：GIN-VUE-ADMIN 官方

### 使用步骤

#### 1. 前往GVA主程序下的initialize/router.go 在Routers 方法最末尾按照你需要的及安全模式添加本插件
    例：
    本插件可以采用gva的配置文件 也可以直接写死内容作为配置 建议为gva添加配置文件结构 然后将配置传入
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
		))

    同样也可以再传入时写死

    PluginInit(PrivateGroup, email.CreateEmailPlug(
    "a@qq.com",
    "b@qq.com",
    "smtp.qq.com",
    "global.GVA_CONFIG.Email.Secret",
    "登录密钥",
    465,
    true,
    ))

### 2. 配置说明

#### 2-1 全局配置结构体说明
    //其中 Form 和 Secret 通常来说就是用户名和密码

    type Email struct {
	    To       string  // 收件人:多个以英文逗号分隔 例：a@qq.com b@qq.com 正式开发中请把此项目作为参数使用 此处配置主要用于发送错误监控邮件
	    From     string  // 发件人  你自己要发邮件的邮箱
	    Host     string  // 服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议
	    Secret   string  // 密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥
	    Nickname string  // 昵称    发件人昵称 自定义即可 可以不填
	    Port     int     // 端口     请前往QQ或者你要发邮件的邮箱查看其smtp协议 大多为 465
	    IsSSL    bool    // 是否SSL   是否开启SSL
    }
#### 2-2 入参结构说明
    //其中 Form 和 Secret 通常来说就是用户名和密码

    type Email struct {
        To      string `json:"to"`      // 邮件发送给谁
        Subject string `json:"subject"` // 邮件标题
        Body    string `json:"body"`    // 邮件内容
    }


### 3. 方法API

    utils.EmailTest(邮件标题，邮件主体) 发送测试邮件
    例:utils.EmailTest("测试邮件"，"测试邮件")
    utils.ErrorToEmail(邮件标题,邮件主体) 错误监控
    例:utils.ErrorToEmail("测试邮件"，"测试邮件")
    utils.Email(目标邮箱多个的话用逗号分隔，邮件标题，邮件主体) 发送测试邮件
    例:utils.Email(”a.qq.com,b.qq.com“,"测试邮件"，"测试邮件")

### 4. 可直接调用的接口

    测试接口： /email/emailTest [post] 已配置swagger

    发送邮件接口接口： /email/emailSend [post] 已配置swagger
    入参：
    type Email struct {
        To      string `json:"to"`      // 邮件发送给谁
        Subject string `json:"subject"` // 邮件标题
        Body    string `json:"body"`    // 邮件内容
    }
   
