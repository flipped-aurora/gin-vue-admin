## GVA 阿里云短信发送功能插件
#### 开发者：GIN-VUE-ADMIN 官方

### 使用步骤

#### 1. 前往GVA主程序下的initialize/router.go 在Routers 方法最末尾按照你需要的及安全模式添加本插件
    例：
    本插件可以采用gva的配置文件 也可以直接写死内容作为配置 建议为gva添加配置文件结构 然后将配置传入
		PluginInit(PublicGroup, sms.CreateAliSmsPlug("短信的AccessKey ID", "短信的AccessKey Secret", "短信的 SignName"))

### 2. 配置说明

#### 2-1 全局配置结构体说明

    type AliSms struct {
        AccessKeyId  string `mapstructure:"accessKeyId" json:"accessKeyId" yaml:"accessKeyId"`    // 短信的AccessKey ID
        AccessSecret string `mapstructure:"accessSecret" json:"accessSecret" yaml:"accessSecret"` // 短信的AccessKey Secret
        SignName     string `mapstructure:"signName" json:"signName" yaml:"signName"`             // 短信的 SignName
    }

#### 2-2 入参结构说明

type AliModel struct {
    Phones        []string `json:"phones"` // 需要发送的手机（可传入多个）
    TemplateCode  string   `json:"templateCode"` // 短信模板的code
    TemplateParam string   `json:"templateParam"` // 短信模板的填充
}

templateParam: 模板使用json字符串 {"code":"xxx"} 对应你模板里面的变量key和value

### 3. 方法API
// 无 

### 4. 可直接调用的接口

    发送邮件接口接口： /aliSms/sendSms [post] 已配置swagger
    入参：
    type AliModel struct {
        Phones        []string `json:"phones"` // 需要发送的手机（可传入多个）
        TemplateCode  string   `json:"templateCode"` // 短信模板的code
        TemplateParam string   `json:"templateParam"` // 短信模板的填充
    }
   
