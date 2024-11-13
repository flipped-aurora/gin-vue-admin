package wallet

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// PrivateKeyFromString 私钥字符串转换为PrivateKey
func PrivateKeyFromString(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// PrivateKeyToAddress ecdsa.PrivateKey私钥转换成地址或者私钥字符串
func PrivateKeyToAddress(privateKey *ecdsa.PrivateKey) (string, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Failed to cast public key to ECDSA")
		return "", errors.New("failed to cast public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Address:", address)
	// 获取私钥
	privateKeyBytes := crypto.FromECDSA(privateKey)
	priv := hexutil.Encode(privateKeyBytes)[2:]
	fmt.Println("私钥", priv)
	//获取地址

	return address, nil
}
