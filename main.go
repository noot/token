package main

import (
	"fmt"

	"github.com/noot/token/test"

	"github.com/ChainSafeSystems/leth/core"
)

func main() {
	err := core.Migrate("testnet", "Token")
	if err != nil {
		fmt.Println("could not deploy Token.sol to testnet")
	}

	fmt.Println("executing tests...")
	test.Test()
}