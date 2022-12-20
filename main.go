package main

import (
	"fmt"
	"github.com/ethereum/go-binds/service"
)

func main() {
	ret, err := service.GetBlackList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}
