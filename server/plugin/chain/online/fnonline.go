package online

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
)

// // BnbBalance 获取账户Bnb余额
func BnbBalance(client *ethclient.Client, address string) (balance *decimal.Decimal, err error) {
	//获取余额
	res, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return nil, err
	}
	//res, err := client.PendingBalanceAt(context.Background(), common.HexToAddress(address))
	//if err != nil {
	//	fmt.Println(err)
	//	//return nil, err
	//}

	balance = WeiToDecimal(res)
	return balance, nil
}

// Bsc20Balance 获取账户代币余额
func Bsc20Balance(client *ethclient.Client, constractaddress, address string) (balance *decimal.Decimal, err error) {
	//实例化一个transactor
	instance, err := NewGoethereumCaller(common.HexToAddress(constractaddress), client)
	if err != nil {
		return balance, err
	}
	//获取余额
	res, err := instance.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		return balance, err
	}
	balance = WeiToDecimal(res)
	return balance, nil
}

// 代币授权
func Bsc20IncreaseAllowance(client *ethclient.Client, contract string, keystore *keystore.Key, spender string, amount *decimal.Decimal) error {
	instance, err := NewGoethereum(common.HexToAddress(contract), client)
	if err != nil {
		return err
	}
	res, err := instance.Allowance(nil, keystore.Address, common.HexToAddress(spender))
	if err != nil {
		return err
	}
	if WeiToDecimal(res).Cmp(decimal.NewFromInt(10000)) > 0 {
		fmt.Println("已经授权成功", keystore.Address)
	}

	//获取nonce
	nonce, _ := client.PendingNonceAt(context.Background(), keystore.Address)
	//获取gasPrice
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	chainID, _ := client.NetworkID(context.Background())
	//构造auth
	auth, _ := bind.NewKeyedTransactorWithChainID(keystore.PrivateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(70000) // in units
	auth.GasPrice = gasPrice
	allowance, err := instance.Approve(auth, common.HexToAddress(spender), DecimalToWei(amount))
	if err != nil {
		return err
	}
	fmt.Println("授权成功", allowance.Hash().Hex())
	return nil
}

// 查询授权额度
func GetApproval(client *ethclient.Client, from, spender string) error {
	instance, err := NewGoethereum(common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"), client)
	if err != nil {
		return err
	}
	res, err := instance.Allowance(nil, common.HexToAddress(from), common.HexToAddress(spender))
	if err != nil {
		return err
	}
	if WeiToDecimal(res).Cmp(decimal.NewFromInt(1000000)) > 0 {
		fmt.Println("已经认证成功", from)
		return nil
	}
	return errors.New("认证额度不足")
}

// Bsc20Transfer 代币转账
func Bsc20Transfer(client *ethclient.Client, keystore *keystore.Key, contract string, to string, amount *decimal.Decimal) (txhash string, err error) {

	//实例化一个transactor
	instance, err := NewGoethereum(common.HexToAddress(contract), client)
	if err != nil {

		return "", err
	}

	am := DecimalToWei(amount)

	//获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), keystore.Address)
	if err != nil {
		return "", err
	}
	//获取gasPrice
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	chainID, _ := client.NetworkID(context.Background())
	//构造auth
	auth, _ := bind.NewKeyedTransactorWithChainID(keystore.PrivateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)    // in wei
	auth.GasLimit = uint64(70000) // in units
	auth.GasPrice = gasPrice

	//调用合约transfer方法
	tx, err := instance.Transfer(auth, common.HexToAddress(to), am)
	if err != nil {
		return "", err
	}
	// 检查交易是否成功
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return "", err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		// 转账失败，可以根据具体情况进行处理
		fmt.Printf("转账失败: %s\n", tx.Hash().Hex())
		return "", errors.New("转账失败")
	}

	fmt.Printf("转账tx sent: %s\n", tx.Hash().Hex())
	return receipt.TxHash.Hex(), nil
}

// BnbTransfer  BNB转账
func BnbTransfer(cli *ethclient.Client, keystore *keystore.Key, toAddress string, amount *decimal.Decimal) (txhash string, err error) {

	//1.构造私钥

	//2.构造公钥(如果只有私钥可以用这里)
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	nonce, err := cli.PendingNonceAt(context.Background(), keystore.Address)
	if err != nil {
		fmt.Println("PendingNonceAt", err)
		return "", err
	}
	//fmt.Println("转出地址", fromAddress)
	//value := new(big.Int)
	//value.SetString(amount, 10)

	gasLimit := uint64(24000) // in units
	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	toAddress = strings.ToLower(toAddress)
	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), DecimalToWei(amount), gasLimit, gasPrice, nil)
	chainID, err := cli.NetworkID(context.Background())
	if err != nil {
		fmt.Println("NewTransaction", err)
		return "", err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), keystore.PrivateKey)
	if err != nil {
		fmt.Println("SignTx", err)
		return signedTx.Hash().Hex(), err
	}
	err = cli.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("SendTransaction1111", err)
		return signedTx.Hash().Hex(), err
	}
	return signedTx.Hash().Hex(), nil
}
