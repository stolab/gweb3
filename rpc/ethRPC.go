package rpc

import (
    "fmt"
    "net/http"
)

// Represent the parameters to allow multiple type
type Parameters interface{}

//represent the endpoint used as the entry point to the network.
//This can be a public endpoint as well as a personal node
type Endpoint struct {
    url      string
    port     string
    endpoint string
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

// NOTE does it really make sense to keep this Endpoint in this form?
//Initialize the endpoint.
func ConnectEndpoint(url string, port string) *Endpoint {
    return &Endpoint{
        url:      url,
        port:     port,
        endpoint: "http://" + url + ":" + port,
    }
}

func (ep *Endpoint) ClientVersion() (*http.Response, error) {
    return ep.HttpRequest([]Parameters{}, RPCendpoint["ClientVersion"])
}

func (ep *Endpoint) MostRecentBlock() (*http.Response, error) {
    return ep.HttpRequest([]Parameters{}, RPCendpoint["MostRecentBlock"])
}

func (ep *Endpoint) GetTransactionCount(address string) (*http.Response, error) {
    params := []Parameters{address, "latest"}

    return ep.HttpRequest(params, RPCendpoint["GetTransactionCount"])
}

// TODO should do a build transaction function
func (ep *Endpoint) SignTransaction(transaction Transaction) (*http.Response, error) { //must be used after in accordance with sendRawTransaction
    params := []Parameters{transaction}
    return ep.HttpRequest(params, RPCendpoint["SignTransaction"])
}

func (ep *Endpoint) SendRawTransaction(rawSignedTransaction string) (*http.Response, error) {
    params := []Parameters{rawSignedTransaction}
    return ep.HttpRequest(params, RPCendpoint["SendRawTransaction"])
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

func (ep *Endpoint) SendTransaction(transaction Transaction) (*http.Response, error) { //must be used with getTransactionReceipt to get the contract address after creating it.
        return ep.HttpRequest([]Parameters{transaction}, RPCendpoint["SendTransaction"])
}

func (ep *Endpoint) GetTransactionByHash(transactionHash string) (*http.Response, error) {
    params := []Parameters{transactionHash}
    return ep.HttpRequest(params, RPCendpoint["GetTransactionByHash"])
}

func (ep *Endpoint) GetBalance(address string) (*http.Response, error) {
    params := []Parameters{address, "latest"} 

    return ep.HttpRequest(params, RPCendpoint["GetBalance"])
}

func (ep *Endpoint) GetStorageAt(contractAdress string, storageAddr int) (*http.Response, error) {
    storageAddrString := fmt.Sprintf("0x%x", storageAddr)

    params := []Parameters{contractAdress, storageAddrString, "latest"}
    return ep.HttpRequest(params, RPCendpoint["GetStorageAt"])

}

func (ep *Endpoint) Sha3(data string) (*http.Response, error) {
    params := []Parameters{data}
    return ep.HttpRequest(params, RPCendpoint["Sha3"])
}

func (ep *Endpoint) GetCode(address string) (*http.Response, error) {
    params := []Parameters{address, "latest"}
    return ep.HttpRequest(params, RPCendpoint["GetCode"])
}

func (ep *Endpoint) GetGasPrice() (*http.Response, error) {
    return ep.HttpRequest([]Parameters{}, RPCendpoint["GasPrice"])
}

func (ep *Endpoint) GetCoinbase() (*http.Response, error) {
    return ep.HttpRequest([]Parameters{}, RPCendpoint["Coinbase"])
}

func (ep *Endpoint) GetBlockReceipts(blockNumber string) (*http.Response, error) {
    params := []Parameters{blockNumber}
    return ep.HttpRequest(params, RPCendpoint["GetBlockReceipts"])
}

func (ep *Endpoint) GetTransactionReceipt(transactionHash string) (*http.Response, error){
    params := []Parameters{transactionHash}
    return ep.HttpRequest(params, RPCendpoint["GetTransactionReceipt"])
}

func (ep *Endpoint) Call(tr Transaction) (*http.Response, error) {
    params := []Parameters{tr}
    return ep.HttpRequest(params, RPCendpoint["Call"])
}

