package test

import (
	"fmt"
	"log"
	"io/ioutil"
	"strings"
	"math/big"
	"path/filepath"
	
	"github.com/noot/token/bindings"

	"github.com/ChainSafeSystems/leth/core"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Log struct {
    // Consensus fields:
    // address of the contract that generated the event
    Address common.Address `json:"address" gencodec:"required"`
    // list of topics provided by the contract.
    Topics []common.Hash `json:"topics" gencodec:"required"`
    // supplied by the contract, usually ABI-encoded
    Data []byte `json:"data" gencodec:"required"`

    // Derived fields. These fields are filled in by the node
    // but not secured by consensus.
    // block in which the transaction was included
    BlockNumber uint64 `json:"blockNumber"`
    // hash of the transaction
    TxHash common.Hash `json:"transactionHash" gencodec:"required"`
    // index of the transaction in the block
    TxIndex uint `json:"transactionIndex" gencodec:"required"`
    // hash of the block in which the transaction was included
    BlockHash common.Hash `json:"blockHash"`
    // index of the log in the receipt
    Index uint `json:"logIndex" gencodec:"required"`

    // The Removed field is true if this log was reverted due to a chain reorganisation.
    // You must pay attention to this field if you receive logs through a filter query.
    Removed bool `json:"removed"`
}

func Test() {
	conn, err := core.NewConnection("default")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}	

	address, err := core.ContractAddress("Token", "testnet")
	if err != nil {
		log.Fatal(err)
	}

	logs := make(chan []types.Log)
	go core.WatchAllEvents(conn, address, big.NewInt(0), logs)
	go func(logsChan chan []types.Log) {
		logs :=<- logsChan
		for _, l := range logs {
			logBytes, _ := l.MarshalJSON()
			fmt.Println(string(logBytes))
		}
	}(logs)

	token, err := bindings.NewToken(address, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a contract: %v", err)
	}

	owner, err := token.Owner(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve owner: %v", err)
	}
	fmt.Println("Contract owner:", owner.Hex())

	path, _ := filepath.Abs("./keystore/UTC--2018-05-17T21-58-52.188632298Z--8f9b540b19520f8259115a90e4b4ffaeac642a30")
	key, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("could not find keystore json file", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(string(key)), "password")
	if err != nil {
		log.Fatalf("could not unlock account")
	}

	tx, err := token.Transfer(auth, common.HexToAddress("0x0000000000000000000000000000000000000000"), big.NewInt(1))
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())

	// Wrap the Token contract instance into a session
	// session := &TokenSession{
	// 	Contract: token,
	// 	CallOpts: bind.CallOpts{
	// 		Pending: true,
	// 	},
	// 	TransactOpts: bind.TransactOpts{
	// 		From:     auth.From,
	// 		Signer:   auth.Signer,
	// 		GasLimit: big.NewInt(3141592),
	// 	},
	// }
	// // Call the previous methods without the option parameters
	// session.Name()
	// session.Transfer("0x0000000000000000000000000000000000000000"), big.NewInt(1))
}

func watchTransfer() {

}