package myutil

import (
	"context"
	"fmt"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/sms"
	"github.com/redis/go-redis/v9"
	"math/rand"
)

const accessKey = "uABw8IJWASS2zDBY5I9EtFNb8D8tS2opMyPsCQLh"
const secretKey = "N8EKUp-lG66L9gVE86sw6aHzRfm7snurSXojCeOE"
const SignatureID = "1826539897645711360"
const TemplateID = "1826556244987490304"

func SendMS(phone string, code string) {
	// 短信发送
	// 1. 引入阿里云短信服务
	// 2. 配置阿里云短信服务
	// 3. 发送短信
	//https://api.smsbao.com/wsms?u=18688831993&p=55a0886fa13747269b497231ad59edae&m={mobile}&c={content}

	result1 := netutil.IsPingConnected("www.baidu.com")

	fmt.Println(result1)

}

// 生成请求认证签名
var manager *sms.Manager

func Qiuniusms(code string, phone string) error {
	var msg = sms.MessagesRequest{
		SignatureID: SignatureID,
		TemplateID:  TemplateID,
		Parameters:  map[string]interface{}{"code": code},
		Mobiles:     []string{phone},
	}
	mac := auth.New(accessKey, secretKey)
	manager = sms.NewManager(mac)
	// 发送短信
	_, err := manager.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

// 生成四位随机数字
func GetRandomNumber(length int) string {
	var result string
	for i := 0; i < length; i++ {
		result += fmt.Sprintf("%d", rand.Intn(10))
	}
	return result
}

// 验证缓存验证码是否正确
func CheckCode(phone string, code string) (bool, error) {
	// 验证缓存验证码是否正确
	// 构造键名
	key := "sms" + phone
	ctx := context.Background()
	fmt.Println(key)
	// 从 Redis 中获取验证码
	storedCode, err := global.GVA_REDIS.Get(ctx, key).Result()
	fmt.Println(storedCode, err)
	if err == redis.Nil {
		// 如果键不存在，则返回 false 和错误信息
		return false, fmt.Errorf("验证码已过期")
	} else if err != nil {
		// 如果发生其他错误，则返回错误
		return false, err
	}

	// 比较输入的验证码与 Redis 中的验证码
	if storedCode == code {
		// 如果验证码正确，则删除 Redis 中的验证码
		err := global.GVA_REDIS.Del(ctx, key).Err()
		if err != nil {
			return false, err
		}
		return true, nil
	}

	// 如果验证码不正确，则返回 false
	return false, fmt.Errorf("验证码错误")
}
