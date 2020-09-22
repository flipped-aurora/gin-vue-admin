package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"gin-vue-admin/global"

	"github.com/jordan-wright/email"
)

func Email(subject string, body string) error {
	to := strings.Split(global.GVA_CONFIG.Email.EmailTo, ",")
	return send(to, subject, body)
}

// ErrorToEmail Error 发送邮件
func ErrorToEmail(subject string, body string) error {
	to := strings.Split(global.GVA_CONFIG.Email.EmailTo, ",")
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		to = to[:len(to)-1]
	}
	return send(to, subject, body)
}

func EmailTest(subject string, body string) error {
	to := []string{global.GVA_CONFIG.Email.EmailFrom}
	return send(to, subject, body)
}

func send(to []string, subject string, body string) error {
	from := global.GVA_CONFIG.Email.EmailFrom
	nickname := global.GVA_CONFIG.Email.EmailNickname
	secret := global.GVA_CONFIG.Email.EmailSecret
	host := global.GVA_CONFIG.Email.EmailHost
	port := global.GVA_CONFIG.Email.EmailPort
	isSSL := global.GVA_CONFIG.Email.EmailIsSSL

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
