package online

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

// 实例化client
func NewClient() (client *ethclient.Client, err error) {
	//url := "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"  eth
	url1 := "https://bsc-dataseed.binance.org/"
	url2 := "https://bitter-smart-haze.bsc.quiknode.pro/4a63a040fc9f5cc7a7840a0c88c35680c50c7ea6/"
	client, err = ethclient.Dial(url1)
	if err != nil {
		client, err = ethclient.Dial(url2)
		if err != nil {
			return
		}
		return
	}
	return
}
