package test

import (
	"fmt"
	"log"

	"github.com/ChainSafeSystems/leth/bindings"
	"github.com/ChainSafeSystems/leth/core"
)

func Test() {
	conn, err := core.NewConnection("default")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}	

	address, err := core.ContractAddress("Token", "testnet")
	if err != nil {
		log.Fatal(err)
	}

	token, err := bindings.NewExample(address, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a contract: %v", err)
	}

	owner, err := token.Owner(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve owner: %v", err)
	}
	fmt.Println("Contract owner:", owner.Hex())
}