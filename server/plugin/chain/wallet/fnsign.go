package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

const signmsg = "LoginFriends"

// signMessage 对消息进行签名
func signMessage(privateKey *ecdsa.PrivateKey, message string) (hexutil.Bytes, error) {
	// 对消息进行KECCAK-256哈希，这是以太坊签名前的标准步骤
	hashedMessage := crypto.Keccak256Hash([]byte(message))

	// 使用私钥对消息的哈希值进行签名，得到的签名包含r、s和v值
	// 注意：以太坊的crypto.Sign函数会自动设置正确的v值（通常是27或28）
	sig, err := crypto.Sign(hashedMessage.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	// 假设需要手动调整V值，这是不常见的做法
	// 如果V值是0或1，按照以太坊标准调整为27或28
	// 确保签名长度为65字节（32字节r + 32字节s + 1字节v）
	if len(sig) != 65 {
		return nil, fmt.Errorf("invalid signature length: %d", len(sig))
	}

	return sig, nil
}

// VerifyEthSignature 验证以太坊签名是否有效，返回true表示签名有效，false表示无效
func VerifyEthSignature(signatureHex, fromAddress string) bool {
	// 解析用户声称的地址
	fromAddr := common.HexToAddress(fromAddress)

	// 解码签名
	sigBytes, err := hexutil.Decode(signatureHex)
	if err != nil {
		fmt.Println("Error decoding signature:", err)
		return false
	}
	// 确保V值正确（27或28）
	if sigBytes[64] >= 27 {
		sigBytes[64] -= 27
	}

	// 计算消息的签名哈希，直接在方法内实现
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(signmsg), signmsg)

	msgHash := crypto.Keccak256([]byte(msg))

	// 恢复公钥
	pubKey, err := crypto.SigToPub(msgHash, sigBytes)
	if err != nil {
		fmt.Println("Error recovering public key:", err)
		return false
	}

	// 将公钥转换为以太坊地址
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	fmt.Println(" Address:", fromAddr)
	fmt.Println("Recovered Address:", recoveredAddr.Hex())
	return strings.EqualFold(recoveredAddr.Hex(), fromAddress) // 直接比较地址的十六进制表示，忽略大小
}
