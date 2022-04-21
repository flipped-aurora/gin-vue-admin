package service

import (
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/global"
	"strings"
)

type AliSmsService struct{}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SendAliSms
//@description: 发送（阿里）短信
//@return: err error

// 模板使用json字符串 {"code":"xxx"} 对应你模板里面的变量key和value
func (e *AliSmsService) SendAliSms(phones []string, templateCode string, templateParam string) (err error) {
	client, _err := CreateClient(tea.String(global.GlobalConfig.AccessKeyId), tea.String(global.GlobalConfig.AccessSecret))
	if _err != nil {
		return _err
	}
	if len(phones) == 0 {
		return errors.New("请输入手机号")
	}

	phonesStr := strings.Join(phones, ",")
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(global.GlobalConfig.SignName),
		PhoneNumbers:  tea.String(phonesStr),
		TemplateCode:  tea.String(templateCode),
		TemplateParam: tea.String(templateParam),
	}
	// 复制代码运行请自行打印 API 的返回值
	_, _err = client.SendSms(sendSmsRequest)
	if _err != nil {
		return _err
	}
	return _err
}
