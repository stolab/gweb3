package contracts

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/stolab/gweb3/rpc"
)

var (
    endpoint *rpc.Endpoint
    err error
    abi *os.File
    contract *Contract 
)

//flag available
var (
    chainURL = flag.String("chainURL", "http://127.0.0.1:8545", "the RPC endoint to use. (Default gnosis chain : http://127.0.0.1:8545)")
    contractAddr = flag.String("contractAddr", "", "the address of the contract to use for testing. (Default: \"\")")
    signerAddr = flag.String("signer", "", "address used to sign transaction")
    ipcPath = flag.String("ipc", "/tmp/geth.ipc", "path to the IPC file. (Default: /tmp/geth.ipc)")
)

// Setup
//IDEA could start geth by default if present in the path
func TestMain(m *testing.M){

    flag.Parse()

    if *contractAddr == "" {
        fmt.Println("Please provide the contract address (ie: --contractAddr addr)")
        os.Exit(1)
    }

    if *signerAddr == ""{
        fmt.Println("Please provide the signer address (ie: --signer addr)")
        os.Exit(1)
    }

    endpoint, err = rpc.ConnectEndpoint(*chainURL)
    if err != nil {
        fmt.Println("Error when connecting to endpoint: %w", err)
    }

    abi, err = os.Open("./test/abi.json")
    if err != nil {
        os.Exit(1)
    }

    contract, err = InitalizeContract(endpoint, *contractAddr, *signerAddr, abi)
    if err != nil {
        fmt.Println("Error when initializing the contract: %w", err)
        os.Exit(1)
    }

    os.Exit(m.Run())

}

func TestCallEmptyParametersFunction(t *testing.T) {

    result, err := contract.Function["getMessage"].Call()
    if err != nil {
        t.Fatalf("Got an error when calling getMessage : %s ", err)
    }
    t.Logf("got the following response, %v", result)
}

func TestContractCall(t *testing.T){

    result, err := contract.Function["setMessage"].Call("My New message")
    if err != nil {
        t.Fatalf("Got and error in the request %s", err)
    }

    t.Logf("Got an answer : %v", result)

    t.Logf("Testing getMessage (view function)")
    result, err = contract.Function["getMessage"].Call()
    if err != nil {
        t.Fatalf("Got and error in the request getMessage %s", err)
    }
    t.Logf("Got an answer : %v", result)

}

//TODO should update the smart contract which should return a big array
func TestBigAnswerOverIPC(t  *testing.T) {

    abi, err = os.Open("./test/abi.json")
    if err != nil {
        t.Fatalf("Error opening abi file: %s", err) 
    }

    IPCconn, err := rpc.ConnectEndpoint(*ipcPath)
    if err != nil {
        t.Fatalf("Error When connecting to IPC %s", err)
    }

    mycontract, err := InitalizeContract(IPCconn, *contractAddr, *signerAddr, abi)
    if err != nil {
        t.Fatalf("Got an error when intializing the contract : %s ", err)
    }

    result, err := mycontract.Function["getMessage"].Call()
    if err != nil {
        t.Fatalf("Got an error when calling getMessage : %s ", err)
    }
    t.Logf("got the following response, %v", result)
}
