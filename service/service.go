package service

import (
	"fmt"
	GA "github.com/ethereum/go-binds/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func GetBlackList() (map[string]GA.GlobalAccessV1StorageUserStatus, error) {
	client, err := ethclient.Dial("http://43.198.66.226:8545")
	if err != nil {
		log.Fatal(err)
	}

	//该地址为proxy合约地址
	address := common.HexToAddress("0x19B6B108e438E6615eAfa89b20450A5DE3B89933")
	instance, err := GA.NewGlobalAccess(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")

	//读取
	n, err := instance.GetUserNumber(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	length := n.Int64()

	res := make(map[string]GA.GlobalAccessV1StorageUserStatus)
	for i := int64(0); i < length; i++ {
		addr, err := instance.UsersList(nil, big.NewInt(i))
		if err != nil {
			log.Fatal(err)
		}
		status, err := instance.GetUserStatus(nil, addr)
		if err != nil {
			log.Fatal(err)
		}
		res[addr.String()] = status
	}
	return res, nil
}
