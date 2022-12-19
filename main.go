package main

import (
	"fmt"
	GA "github.com/ethereum/go-binds/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("43.198.66.226:8545")
	if err != nil {
		log.Fatal(err)
	}

	//该地址为合约地址-目前未部署，待填入正确的地址
	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	instance, err := GA.NewGlobalAccess(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
	//合约成功加载后，就可以instance.调用各种方法
}
