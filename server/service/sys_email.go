package service

import (
	"gin-vue-admin/utils"
)

// @title    EmailTest
// @description   发送邮件测试
// @auth                    （2020/09/08  13:58
// @return    err            error

func EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = utils.EmailTest(subject, body)
	return err
}
