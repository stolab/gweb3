package rpc

import (
	"fmt"
	"net/url"
	"os"
)

// Represent the parameters to allow multiple type
type Parameters interface{}

//represent the endpoint used as the entry point to the network.
//This can be a public endpoint as well as a personal node
type Endpoint struct {
    ParsedURL *url.URL
    endpoint string
    isIPC bool
    requestId int
}

// Represent the transaction 
type Transaction struct {
    From  string `json:"from"`
    To    string `json:"to"`
    Value string `json:"value"`
    Input string `json:"input"`
    // gasLimit int
    // MaxPriorityFeePerGas int
    // maxFeePerGas int
    // nonce string
}

//The structure of the RPC call made to the endpoint 
type RPCTransaction struct {
    Jsonrpc string       `json:"jsonrpc"`
    Method  string       `json:"method"`
    Params  []Parameters `json:"params"`
    Id      int          `json:"id"` //not sure what this is used for, seems like it can be anything.
}

//Initialize the endpoint.
func ConnectEndpoint(Rawurl string) (*Endpoint, error) {

    _, error := os.Stat(Rawurl)
    //case this is a IPC file
    if error == nil {
        return &Endpoint{
            endpoint: Rawurl,
            isIPC: true,
            requestId: 1,
        }, nil
    }

    //case this is a URL
    u, err := url.Parse(Rawurl)
    if err != nil {
        return nil, err
    }
    return &Endpoint{
        ParsedURL: u, 
        endpoint: u.String(),
        isIPC: false,
        requestId: 1,
    }, nil
}

func (ep *Endpoint) ClientVersion() (*RPCResponse, error) {
    return ep.Request([]Parameters{}, RPCendpoint["ClientVersion"])
}

func (ep *Endpoint) NetworkId() (*RPCResponse, error) {
    return ep.Request([]Parameters{}, RPCendpoint["NetworkId"])
}

func (ep *Endpoint) MostRecentBlock() (*RPCResponse, error) {
    return ep.Request([]Parameters{}, RPCendpoint["MostRecentBlock"])
}

func (ep *Endpoint) GetTransactionCount(address string) (*RPCResponse, error) {
    params := []Parameters{address, "latest"}

    return ep.Request(params, RPCendpoint["GetTransactionCount"])
}

func (ep *Endpoint) SignTransaction(transaction Transaction) (*RPCResponse, error) { //must be used after in accordance with sendRawTransaction
    params := []Parameters{transaction}
    return ep.Request(params, RPCendpoint["SignTransaction"])
}

func (ep *Endpoint) SendRawTransaction(rawSignedTransaction string) (*RPCResponse, error) {
    params := []Parameters{rawSignedTransaction}
    return ep.Request(params, RPCendpoint["SendRawTransaction"])
}

func BuildTransaction(from string, to string, value string, input string) Transaction {
    transaction := Transaction{
        From:  from,
        To:    to,
        Value: value,
        Input: input,
    }
    return transaction
}

func (ep *Endpoint) SendTransaction(transaction Transaction) (*RPCResponse, error) { //must be used with getTransactionReceipt to get the contract address after creating it.
        return ep.Request([]Parameters{transaction}, RPCendpoint["SendTransaction"])
}

func (ep *Endpoint) GetTransactionByHash(transactionHash string) (*RPCResponse, error) {
    params := []Parameters{transactionHash}
    return ep.Request(params, RPCendpoint["GetTransactionByHash"])
}

func (ep *Endpoint) GetBalance(address string) (*RPCResponse, error) {
    params := []Parameters{address, "latest"} 

    return ep.Request(params, RPCendpoint["GetBalance"])
}

func (ep *Endpoint) GetStorageAt(contractAdress string, storageAddr int) (*RPCResponse, error) {
    storageAddrString := fmt.Sprintf("0x%x", storageAddr)

    params := []Parameters{contractAdress, storageAddrString, "latest"}
    return ep.Request(params, RPCendpoint["GetStorageAt"])

}

func (ep *Endpoint) Sha3(data string) (*RPCResponse, error) {
    params := []Parameters{data}
    return ep.Request(params, RPCendpoint["Sha3"])
}

func (ep *Endpoint) GetCode(address string) (*RPCResponse, error) {
    params := []Parameters{address, "latest"}
    return ep.Request(params, RPCendpoint["GetCode"])
}

func (ep *Endpoint) GetGasPrice() (*RPCResponse, error) {
    return ep.Request([]Parameters{}, RPCendpoint["GasPrice"])
}

func (ep *Endpoint) GetCoinbase() (*RPCResponse, error) {
    return ep.Request([]Parameters{}, RPCendpoint["Coinbase"])
}

func (ep *Endpoint) GetBlockReceipts(blockNumber string) (*RPCResponse, error) {
    params := []Parameters{blockNumber}
    return ep.Request(params, RPCendpoint["GetBlockReceipts"])
}

func (ep *Endpoint) GetTransactionReceipt(transactionHash string) (*RPCResponse, error){
    params := []Parameters{transactionHash}
    return ep.Request(params, RPCendpoint["GetTransactionReceipt"])
}

func (ep *Endpoint) Call(tr Transaction) (*RPCResponse, error) {
    params := []Parameters{tr}
    return ep.Request(params, RPCendpoint["Call"])
}

func (ep *Endpoint) DeployContract(sender string, contractByteCodePath string) (*RPCResponse, error){

    contractContent, err := os.ReadFile(contractByteCodePath)
    if err != nil {
        return nil, err
    }
    contractPayload := fmt.Sprintf("0x%s", contractContent)
    tr := BuildTransaction(sender, "0x0000000000000000000000000000000000000000", "0x0", contractPayload)
    rep, err := ep.SendTransaction(tr)
    if err != nil {
        return nil, err
    }

    t, err := ep.GetTransactionReceipt(rep.Result)
    if err != nil {
        return nil, err
    }
    return t, nil

}
