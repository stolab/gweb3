package rpc

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"testing"
)

// global variable
var (
    endpoint *Endpoint
    err error
)

// flag available
var (
    chainURL = flag.String("chainURL", "https://rpc2.sepolia.org", "the RPC endoint to use. (Default gnosis chain : https://rpc2.sepolia.org)")
    ipcPath = flag.String("ipc", "/tmp/geth.ipc", "the path to the IPC file. (default : /tmp/geth.ipc)")
    transactionHash = flag.String("trHash", "0xe88c6ef5de26f616184690eac79846de1c531dba0f1407179c2ea87e62a29e6d", "The transaction hash to use for test. (Default : 0xe88c6ef5de26f616184690eac79846de1c531dba0f1407179c2ea87e62a29e6d)")
    account = flag.String("account", "0xA23c9035AfD3e34690d80804B33Bdf1b93c0A604", "The address to use for some check. (Default: 0xA23c9035AfD3e34690d80804B33Bdf1b93c0A604)")
    contractAddr = flag.String("contractAddr", "0xEd4Af1Ef4c47d25fC8f7D2Bae7Da3dDF9B34fB0B", "Address of the contract to use. (Default : 0xEd4Af1Ef4c47d25fC8f7D2Bae7Da3dDF9B34fB0B)")
)

//helper to help marshalling a answer into a string
//NOTE could be move as an helper in RPC package ?
func marshalling(answer *RPCResponse) (string, error){
    marshalled, err := json.Marshal(answer)
    if err != nil {
        return "", nil
    }
    return string(marshalled), nil
}

//Setup function
func TestMain(m *testing.M){

    flag.Parse()

    endpoint, err =  ConnectEndpoint(*chainURL)
    if err != nil {
        fmt.Println("Error when connecting to the endpoint: %w", err)
        os.Exit(1) 
    }

    os.Exit(m.Run())
}

func TestClientVersion(t *testing.T){
	//To check if everything went fine we just check if the json response contain a error field
	r, err := endpoint.ClientVersion()
	if err != nil {
		t.Fatalf(`Got an error : %q`, err)
	}

    if r.Error != nil {
        t.Fatalf(`Got an error in the answer: %q`, r.Error)
    }

    stringResult, err := marshalling(r)
    if err != nil {
        t.Fatalf("Got an error from unmarshalling: %q", err)
    }

	t.Logf(`Answer from the request : %s`, stringResult)
}

func TestMostRecentBlock(t *testing.T){
    r, err := endpoint.MostRecentBlock()
    if err != nil {
        t.Fatalf(`Got an error: %q`, err)
    }
    stringResult, err := marshalling(r)
    if err != nil {
        t.Fatalf("Error when unmarshalling: %q", err)
    }

    t.Logf(`Response: %s`, stringResult)
}

func TestGetTransactionCount(t *testing.T){
	r, err := endpoint.GetTransactionCount(*account)
	if err != nil {
		t.Fatalf(`Got an error : %q`, err)
	}

    if r.Error != nil {
        t.Fatalf(`Got an error in the rpc answer: %q`, r.Error)
    }

    stringResult, err := marshalling(r)
    if err != nil {
        t.Fatalf("Got an error from unmarshalling: %q", err)
    }
	t.Logf(`Answer from the request : %s`, stringResult)
}

func TestSendTransaction(t *testing.T){
    transaction := BuildTransaction(*account, *account, "0x918e72a", "")
	r, err := endpoint.SendTransaction(transaction)
	if err != nil {
		t.Fatalf(`Got an error : %q`, err)
	}

    if r.Error != nil {
        t.Fatalf(`Got an error in the rpc answer : %q`, r.Error)
    }

    stringResult, err := marshalling(r)
    if err != nil {
        t.Fatalf("Got an error from unmarshalling: %q", err)
    }
	t.Logf(`Answer from the request : %s`, stringResult)
}

func TestGetTransactionByHash(t *testing.T){
	r, err := endpoint.GetTransactionByHash(*transactionHash)
	if err != nil {
		t.Fatalf(`Got an error : %q`, err)
	}

    if r.Error != nil {
        t.Fatalf(`Got an error in the rpc answer : %q`, r.Error)
    }
    stringResult, err := marshalling(r)
    if err != nil {
        t.Fatalf("Got an error from unmarshalling: %q", err)
    }
	t.Logf(`Answer from the request : %s`, stringResult)
}

func TestGetBalance(t *testing.T){
	r, err := endpoint.GetBalance(*account)
	if err != nil {
		t.Fatalf(`Got an error : %q`, err)
	}

    if r.Error != nil {
        t.Fatalf(`Got an error in the rpc answer : %q`, r.Error)
    }

    stringResult, err := marshalling(r)
    if err != nil {
        t.Fatalf("Got an error from unmarshalling: %q", err)
    }
	t.Logf(`Answer from the request : %s`, stringResult)
}

func TestGetTransactionReceipt(t *testing.T) {
	r, err := endpoint.GetTransactionReceipt(*transactionHash)
	if err != nil {
		t.Fatalf("Error when getting the receipt %s", err)
	}
	t.Logf(`Answer from the request : %v`, r)
}

func TestGetStorageAt(t *testing.T){

	r, err := endpoint.GetStorageAt(*contractAddr, 1)
	if err != nil {
		t.Fatalf("Got an error in the transaction: %s", err)
	}

    if r.Error != nil {
        t.Fatalf(`Got an error in the rpc answer : %q`, r.Error)
    }

    stringResult, err := marshalling(r)
    if err != nil {
        t.Fatalf("Got an error from unmarshalling: %q", err)
    }
	t.Logf(`Answer from the request : %s`, stringResult)
}

//Should run geth --dev to test it.
func TestIPC(t *testing.T){
    e, err := ConnectEndpoint(*ipcPath)
    if err != nil {
        t.Fatalf("Got and error : %q", err)
    }
    
    t.Logf("Testing the client version")
    r, err := e.ClientVersion()
    if err != nil {
        t.Fatalf("Got an error: %q", err)
    }

    stringResult, err := marshalling(r)
    if err != nil {
        t.Fatalf("Got an error from unmarshalling: %q", err)
    }
    t.Logf("Received an answer: %s", stringResult)
}

func TestDeployContract(t *testing.T){
    rep, err := endpoint.DeployContract(*account, "./test/HelloWorld.bin")
    if err != nil {
        t.Fatalf("error when deploying contract: %q", err)
    }
    t.Logf("Contract address: %s", rep.ContractAddress)
}
