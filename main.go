package main

import (
	"fmt"

	"github.com/noot/token/test"

	"github.com/ChainSafeSystems/leth/core"
)

func migrate() {
	err := core.Migrate("testnet", "Token")
	if err != nil {
		fmt.Println("could not deploy Token.sol to testnet")
	}
}

func main() {
	//migrate()

	fmt.Println("executing tests...")
	test.Test()
}