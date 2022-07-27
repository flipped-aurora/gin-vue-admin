package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/global"

	"github.com/jordan-wright/email"
)

//@author: [maplepie](https://github.com/maplepie)
//@function: Email
//@description: Email发送方法
//@param: subject string, body string
//@return: error

func Email(To, subject string, body string) error {
	to := strings.Split(To, ",")
	return send(to, subject, body)
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: ErrorToEmail
//@description: 给email中间件错误发送邮件到指定邮箱
//@param: subject string, body string
//@return: error

func ErrorToEmail(subject string, body string) error {
	to := strings.Split(global.GlobalConfig.To, ",")
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		to = to[:len(to)-1]
	}
	return send(to, subject, body)
}

//@author: [maplepie](https://github.com/maplepie)
//@function: EmailTest
//@description: Email测试方法
//@param: subject string, body string
//@return: error

func EmailTest(subject string, body string) error {
	to := []string{global.GlobalConfig.To}
	return send(to, subject, body)
}

//@author: [maplepie](https://github.com/maplepie)
//@function: send
//@description: Email发送方法
//@param: subject string, body string
//@return: error

func send(to []string, subject string, body string) error {
	from := global.GlobalConfig.From
	nickname := global.GlobalConfig.Nickname
	secret := global.GlobalConfig.Secret
	host := global.GlobalConfig.Host
	port := global.GlobalConfig.Port
	isSSL := global.GlobalConfig.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}
