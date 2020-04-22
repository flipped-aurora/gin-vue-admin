package tools

import (
	"fmt"
	"gin-vue-admin/config"
	"math/rand"
	"strings"
	"submail_go_sdk/submail/sms"
	"time"
)

func SMSSend(phone string) string {
	conf := make(map[string]string)

	conf["appid"] = config.GinVueAdminconfig.SMSConfig.AppId
	conf["appkey"] = config.GinVueAdminconfig.SMSConfig.AppKey
	conf["signType"] = config.GinVueAdminconfig.SMSConfig.SignType

	submail := sms.CreateXsend(conf)

	submail.SetProject("q7T214")
	submail.SetTo(phone)
	print(phone)
	code := GenValidateCode(6)
	submail.AddVar("code", code)
	submail.AddVar("time", "1分钟")
	submail.Xsend()
	return code
}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
