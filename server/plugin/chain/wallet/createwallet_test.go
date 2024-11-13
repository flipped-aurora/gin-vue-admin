package wallet

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestWallet(t *testing.T) {

	//GetInfoByKeyStore() //创建账户 0x881FF4C96CF02F3C005391eeE06Cef4000BfFfC1
	//BNB Address: 0xa4416da064bBeE02d0f71130CF1414279A32280a
	//GenerateMultipleAccounts() //创建多个账户
	store, err := ImportKeyStore("0x881FF4C96CF02F3C005391eeE06Cef4000BfFfC1") //导入keystore
	if err != nil {
		return
	}
	address, err := PrivateKeyToAddress(store.PrivateKey) //获取私钥字符串
	if err != nil {
		return
	}
	fmt.Println(address)

	//GetWallet()
	//GetNewMnemonic()
}

func TestSign(t *testing.T) {
	type args struct {
		Address string
		Amount  decimal.Decimal
	}
	var arg = args{
		Address: "0xAbDc60C4d049c66CdcEa7D7b5F0e371744356946",
		Amount:  decimal.NewFromFloat(0.01),
	}
	//序列化arg
	marshal, err := json.Marshal(arg)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

	store, err := ImportKeyStore("0x7dd8ea3d5725f307bf2213bfc919d90853c36542") //导入keystore
	if err != nil {
		return
	}
	//fmt.Println(PrivateKeyToAddress(store.PrivateKey))

	bytes, err := signMessage(store.PrivateKey, "login")
	if err != nil {
		return
	}
	fmt.Printf("Signature: %s\n", bytes.String())

	// 假设这是你的公钥和签名字符串
	Address := "0xAbDc60C4d049c66CdcEa7D7b5F0e371744356946"
	//message := fmt.Sprintf("%s", string(marshal))

	signatureHex := "0xd0aac1fceefe2d4a16e94172cf3f3ea0478f28aed52f6e06e27acd59ca7ca4576a5c22e133d79fd610efd53860d7f2a3207e854c8a21a196266fe4a90c46ee421c"
	//VerifySignature()

	bol := VerifyEthSignature(signatureHex, Address)
	fmt.Println(bol)
}

func TestGetAddress(t *testing.T) {
	GetWallet()
	GetNewMnemonic()
}
