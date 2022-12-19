package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	GA "github.com/ethereum/go-binds/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
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
	//合约成功加载后，就可以instance.调用各种方法
	//先写:4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356是operator的私钥
	privateKey, err := crypto.HexToECDSA("4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(8888))
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(8000000) // in units
	auth.GasPrice = gasPrice

	tx, err := instance.SetQuota(auth, big.NewInt(1), big.NewInt(1), big.NewInt(10000000))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(tx.Hash().String())
	//读取
	v, err := instance.Quotas(nil, big.NewInt(1), big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
}
