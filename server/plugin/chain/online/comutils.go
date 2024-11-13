package online

import (
	"github.com/shopspring/decimal"
	"math/big"
)

// WeiToDecimal 将以Wei表示的金额转换为decimal.Decimal形式，适用于如Ether的转换（默认1 Ether = 10^18 Wei）
func WeiToDecimal(wei *big.Int) *decimal.Decimal {
	// 1 Ether 等于 10^18 Wei
	etherFactor := decimal.NewFromInt(10).Pow(decimal.NewFromInt(18))

	// 将Wei转换为big.Float，然后除以10^18得到Ether值
	weiFloat := decimal.NewFromBigInt(wei, 0)
	etherValue := weiFloat.Div(etherFactor)

	return &etherValue
}

// DecimalToWei 将decimal.Decimal形式的金额转换回以Wei表示的整数
func DecimalToWei(dec *decimal.Decimal) *big.Int {
	// 1 Ether 等于 10^18 Wei
	etherFactor := decimal.NewFromInt(10).Pow(decimal.NewFromInt(18))

	// 将金额乘以10^18得到Wei值
	weiValue := dec.Mul(etherFactor)

	// 转换为big.Int
	weiBigInt := weiValue.BigInt()

	return weiBigInt
}
