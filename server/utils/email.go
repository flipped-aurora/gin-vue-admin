package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"gin-vue-admin/global"

	"github.com/jordan-wright/email"
)

//@author: [maplepie](https://github.com/maplepie)
//@function: Email
//@description: Email发送方法
//@param: subject string, body string
//@return: error

func Email(subject string, body string) error {
	to := strings.Split(global.GVA_CONFIG.Email.To, ",")
	return send(to, subject, body)
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: ErrorToEmail
//@description: 给email中间件错误发送邮件到指定邮箱
//@param: subject string, body string
//@return: error

func ErrorToEmail(subject string, body string) error {
	to := strings.Split(global.GVA_CONFIG.Email.To, ",")
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
	to := []string{global.GVA_CONFIG.Email.From}
	return send(to, subject, body)
}

//@author: [maplepie](https://github.com/maplepie)
//@function: send
//@description: Email发送方法
//@param: subject string, body string
//@return: error

func send(to []string, subject string, body string) error {
	from := global.GVA_CONFIG.Email.From
	nickname := global.GVA_CONFIG.Email.Nickname
	secret := global.GVA_CONFIG.Email.Secret
	host := global.GVA_CONFIG.Email.Host
	port := global.GVA_CONFIG.Email.Port
	isSSL := global.GVA_CONFIG.Email.IsSSL

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
