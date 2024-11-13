package wallet

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const AccountPassword = "gva65322"
const KeyPath = "./tmp"

// GetInfoByKeyStore 生成一个新的账户通过keystore文件
func GetInfoByKeyStore() {

	// 假设我们已经有了一个KeyStore实例和一个账户
	keyStore := keystore.NewKeyStore(KeyPath, keystore.StandardScryptN, keystore.StandardScryptP)

	// 生成一个新的账户
	account, err := keyStore.NewAccount(AccountPassword)
	if err != nil {
		panic(err)
	}
	// 获取账户的地址
	address := account.Address
	fmt.Println(account)
	// 输出地址
	fmt.Println("BNB Address:", address.Hex())

}

/*
GenerateMultipleAccounts
生成多个地址
keydir: keystore文件的路径
numAccounts: 要生成的账户数量
password: 密码
*/
func GenerateMultipleAccounts(numAccounts int) {
	keyStore := keystore.NewKeyStore(KeyPath, keystore.StandardScryptN, keystore.StandardScryptP)

	for i := 0; i < numAccounts; i++ {
		account, err := keyStore.NewAccount(AccountPassword)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Account %d BNB Address: %s\n", i+1, account.Address.Hex())
	}
}

// ImportKeyStore 导入keystore文件
func ImportKeyStore(address string) (*keystore.Key, error) {
	// 查找keystore文件
	filePath, err := findKeyStoreFile(address)
	if err != nil {
		fmt.Println("查找keystore文件失败:", err)
		return nil, err
	}

	// 读取keystore文件
	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// 解锁keystore
	key, err := keystore.DecryptKey(jsonBytes, AccountPassword)
	if err != nil {
		fmt.Println("解锁keystore失败:", err)
		return nil, err
	}

	return key, err
}

// 生成随机私钥和地址
func GetWallet() {
	// 生成一个新的随机私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("Failed to generate private key:", err)
		return
	}

	// 从私钥中获取公钥的字节表示
	publicKeyBytes := crypto.FromECDSA(privateKey)

	// 将公钥字节转换为以太坊地址
	address := common.BytesToAddress(publicKeyBytes).Hex()

	// 打印私钥的16进制表示（在实际应用中不要打印或保存私钥，这是不安全的）
	fmt.Println("Private Key (Hex):", hex.EncodeToString(publicKeyBytes))

	// 打印以太坊地址
	fmt.Println("Ethereum Address:", address)
}

// 生成助记词
func GetNewMnemonic() {
	// 生成随机256位熵
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		panic(err)
	}

	// 从熵生成助记词
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}

	// 输出助记词
	fmt.Println("Mnemonic Phrase:", mnemonic)
}

// findKeyStoreFile 根据地址模糊匹配查找对应的keystore文件路径
func findKeyStoreFile(address string) (string, error) {
	name := address[2:] // 去掉前缀"0x"

	dir, err := ioutil.ReadDir(KeyPath)
	if err != nil {
		return "", err
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		if strings.Contains(strings.ToLower(file.Name()), strings.ToLower(name)) {
			return filepath.Join(KeyPath, file.Name()), nil
		}
	}

	return "", fmt.Errorf("未找到与地址 %s 匹配的keystore文件", address)
}
