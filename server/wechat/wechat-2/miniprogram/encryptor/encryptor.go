package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/miniprogram/context"
)

// Encryptor struct
type Encryptor struct {
	*context.Context
}

// NewEncryptor 实例
func NewEncryptor(context *context.Context) *Encryptor {
	basic := new(Encryptor)
	basic.Context = context
	return basic
}

var (
	// ErrAppIDNotMatch appid不匹配
	ErrAppIDNotMatch = errors.New("app id not match")
	// ErrInvalidBlockSize block size不合法
	ErrInvalidBlockSize = errors.New("invalid block size")
	// ErrInvalidPKCS7Data PKCS7数据不合法
	ErrInvalidPKCS7Data = errors.New("invalid PKCS7 data")
	// ErrInvalidPKCS7Padding 输入padding失败
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

// PlainData 用户信息/手机号信息
type PlainData struct {
	OpenID          string `json:"openId"`
	UnionID         string `json:"unionId"`
	NickName        string `json:"nickName"`
	Gender          int    `json:"gender"`
	City            string `json:"city"`
	Province        string `json:"province"`
	Country         string `json:"country"`
	AvatarURL       string `json:"avatarUrl"`
	Language        string `json:"language"`
	PhoneNumber     string `json:"phoneNumber"`
	OpenGID         string `json:"openGId"`
	MsgTicket       string `json:"msgTicket"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
	Watermark       struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark"`
}

// pkcs7Unpad returns slice of the original data without padding
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return data[:len(data)-n], nil
}

// GetCipherText returns slice of the cipher text
func GetCipherText(sessionKey, encryptedData, iv string) ([]byte, error) {
	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	cipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	if len(ivBytes) != aes.BlockSize {
		return nil, fmt.Errorf("bad iv length %d", len(ivBytes))
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, err = pkcs7Unpad(cipherText, block.BlockSize())
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

// Decrypt 解密数据
func (encryptor *Encryptor) Decrypt(sessionKey, encryptedData, iv string) (*PlainData, error) {
	cipherText, err := GetCipherText(sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}
	var plainData PlainData
	err = json.Unmarshal(cipherText, &plainData)
	if err != nil {
		return nil, err
	}
	if plainData.Watermark.AppID != encryptor.AppID {
		return nil, ErrAppIDNotMatch
	}
	return &plainData, nil
}
