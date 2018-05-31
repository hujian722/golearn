package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/ethereum/go-ethereum/rpc"
)

type Block struct {
	Number string
}
type Args struct {
	S string
}

type Result struct {
	String string
	Int    int
	Args   *Args
}

func main() {
	// Connect the client
	client, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("could not create ipc client: %v", err)
	}
	var lastBlock interface{} //Block
	err = client.Call(&lastBlock, "eth_getBlockTransactionCountByNumber", "0x2")
	if err != nil {
		fmt.Println("can't get latest block:", err)
		return
	}
	spew.Dump(lastBlock)
	// Print events from the subscription as they arrive.
	//fmt.Printf("latest block: %v\n", lastBlock.Number)
	//=============
	// var resp Result
	// if err := client.Call(&resp, "service_echo", "hello", 10, &Args{"world"}); err != nil {
	// 	fmt.Println("service_echo:", err)
	// 	return
	// }
	// fmt.Printf("resp resp: %+v\n", resp)
}
