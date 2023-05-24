package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// RSADecrypt 数据解密
func RSADecrypt(privateKey string, ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("PrivateKey format error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		oldErr := err
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("ParsePKCS1PrivateKey error: %s, ParsePKCS8PrivateKey error: %s", oldErr.Error(), err.Error())
		}
		switch t := key.(type) {
		case *rsa.PrivateKey:
			priv = key.(*rsa.PrivateKey)
		default:
			return nil, fmt.Errorf("ParsePKCS1PrivateKey error: %s, ParsePKCS8PrivateKey error: Not supported privatekey format, should be *rsa.PrivateKey, got %T", oldErr.Error(), t)
		}
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// RSADecryptBase64 Base64解码后再次进行RSA解密
func RSADecryptBase64(privateKey string, cryptoText string) ([]byte, error) {
	encryptedData, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		return nil, err
	}
	return RSADecrypt(privateKey, encryptedData)
}
