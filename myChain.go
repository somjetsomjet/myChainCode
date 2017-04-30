package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type MyChainCode struct {
}

func (t *MyChainCode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Init called")
	
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(args[2], []byte(args[2]))
	if err != nil {
		return nil, err
	}
	fmt.Printf("Init end")
	return nil, nil
}

func (t *MyChainCode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Query called")

	Avalbytes, err := stub.GetState(args[0])
	
	jsonResp := "{\"Key\":\"" + args[0] + "\",\"Value\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return Avalbytes, nil
}

func (t *MyChainCode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called")
	
	if function == "init" {
		return t.Init(stub, args)
	} else if function == "update" {
		return t.Update(stub, args)
	} 
	return nil, errors.New("Received unknown function invocation")
}

func (t *MyChainCode) Update(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Update called")
	
	err = stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func main() {
    err := shim.Start(new(MyChainCode))
    if err != nil {
        fmt.Printf("Error starting MyChainCode chaincode: %s", err)
    }
}
