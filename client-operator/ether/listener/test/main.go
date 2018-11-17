package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"../abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var balance string

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	var balanceChannel chan string = make(chan string)

	go Balance(balanceChannel)
	go UpdateBalance(balanceChannel, &balance)

	wg.Wait()
}

func Balance(c chan string) {
	for {
		ethClient, err := ethclient.Dial("https://rinkeby.infura.io")
		if err != nil {
			log.Println(err)
		}
		instance, err := abi.NewToken(common.HexToAddress("0xe50466567878fca56c55ee3c5f988d3ce88b5958"), ethClient) // TLCP Token
		if err != nil {
			log.Println(err)
		}
		balance, err := instance.BalanceOf(nil, common.HexToAddress("0xa0560e36f9d48dd82d8ed8c94a5d6b56020ed43b")) // Random address
		if err != nil {
			log.Println(err)
		}
		c <- balance.String()
	}
}

func UpdateBalance(c chan string, balance *string) {
	for {
		msg := <-c
		*balance = msg
		fmt.Println(*balance)
		time.Sleep(time.Second * 2)
	}
}
