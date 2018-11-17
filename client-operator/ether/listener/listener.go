package listener

import (
	"log"

	"./abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Balance string

func UpdateTokenBalance(b *string) {
	for {
		*b = TokenBalance()
	}
}

func TokenBalance() string {

	ethClient, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Println(err)
	}

	instance, err := abi.NewToken(common.HexToAddress("0xe50466567878fca56c55ee3c5f988d3ce88b5958"), ethClient) // TLCP Token
	if err != nil {
		log.Println(err)
	}

	tb, err := instance.BalanceOf(nil, common.HexToAddress("0xa0560e36f9d48dd82d8ed8c94a5d6b56020ed43b")) // Random address
	if err != nil {
		log.Println(err)
	}

	return tb.String()
}
