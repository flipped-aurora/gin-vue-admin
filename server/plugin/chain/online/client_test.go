package online

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/chain/wallet"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"testing"
)

var usdtcontract = "0x55d398326f99059fF775485246999027B3197955"

func TestClient(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		return
	}
	fmt.Println(client)
	id, err := client.ChainID(context.Background())
	if err != nil {
		return
	}
	fmt.Println(id)
	latestBlock, err := client.BlockByNumber(context.Background(), nil) // nil 表示查询最新的区块
	if err != nil {
		log.Fatalf("Failed to fetch the latest block: %v", err)
	}
	fmt.Println(latestBlock.GasLimit())
	//查询余额
	balance, err := BnbBalance(client, "0x88EA65Ce12BB49C4385424Eb0324F18AbCbC126F")
	if err != nil {
		return
	}
	fmt.Println(balance.String())
	//导入私钥
	//0xC9a5ee343aF59812A4B743eD32611223efB509e4  //节点分红
	//0x79D954564b77C9550327B3e11cFe31472bc1e0d0 //盲盒分红
	store, err := wallet.ImportKeyStore("0x79D954564b77C9550327B3e11cFe31472bc1e0d0")
	if err != nil {
		return
	}
	dcm := decimal.NewFromInt(10000)
	err = Bsc20IncreaseAllowance(client, usdtcontract, store, "0xcF888cBc43b7976f179B7a0f7542eB1202659714", &dcm)
	if err != nil {
		fmt.Println(err)
		return
	}

	client.Close()
}

func TestBTD(t *testing.T) {
	// 示例：创建一个big.Int实例
	bigAmount := big.NewInt(1234567891234567890)

	decAmount := WeiToDecimal(bigAmount)

	fmt.Printf("Big.Int转换为decimal.Decimal: %s\n", decAmount.String())

	// 创建一个decimal.Decimal实例
	*decAmount, _ = decimal.NewFromString("1234567891234567890")
	bigAmount = DecimalToWei(decAmount)

	fmt.Printf("decimal.Decimal转换为Big.Int: %s\n", bigAmount.String())
}

func TestWs(t *testing.T) {
	//WsBn()
	client, err := NewClient()
	if err != nil {
		return
	}
	balance, err := Bsc20Balance(client, usdtcontract, "0x88EA65Ce12BB49C4385424Eb0324F18AbCbC126F")
	if err != nil {
		return
	}
	fmt.Println(balance)
}
