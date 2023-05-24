package werun

import (
	"encoding/json"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/context"
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/encryptor"
)

// WeRun 微信运动
type WeRun struct {
	*context.Context
}

// Data 微信运动数据
type Data struct {
	StepInfoList []struct {
		Timestamp int `json:"timestamp"`
		Step      int `json:"step"`
	} `json:"stepInfoList"`
}

// NewWeRun 实例化
func NewWeRun(ctx *context.Context) *WeRun {
	return &WeRun{Context: ctx}
}

// GetWeRunData 解密数据
func (werun *WeRun) GetWeRunData(sessionKey, encryptedData, iv string) (*Data, error) {
	cipherText, err := encryptor.GetCipherText(sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}
	var weRunData Data
	err = json.Unmarshal(cipherText, &weRunData)
	if err != nil {
		return nil, err
	}
	return &weRunData, nil
}
