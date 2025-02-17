package rpc

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"time"
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
    Input string `json:"data"`
    // gasLimit int
    // MaxPriorityFeePerGas int
    // maxFeePerGas int
    // nonce string
}

// Represent the transaction to create a contract
// When creating a contract, the TO field should be absent
type ContractCreationTr struct {
    From  string `json:"from"`
    Value string `json:"value"`
    Input string `json:"data"`
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

// TransactionReceiptResponse represent the Response when
// querying the receipt of a transaction
type TransactionReceiptResponse struct {
    BlockHash string `json:"blockHash"`
    BlockNumber string `json:"blockNumber"`
    ContractAddress string `json:"contractAddress"`
    CumulativeGasUser string `json:"cumulativeGasUser"`
    EffectiveGasPrice string `json:"effectiveGasPrice"`
    From string `json:"from"`
    GasUsed string `json:"gasUsed"`
    LogsBloom string `json:"logsBloom"`
    Status string `json:"status"`
    To string `json:"to"`
    TransactionHash string `json:"transactionHash"`
    TransactionIndex string `json:"transactionIndex"`
    Type string `json:"type"`
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

//return by default the IPC endpoint 
// created when running geth locally
func ConnectLocalEndpoint() (*Endpoint, error) {
    return &Endpoint{
        endpoint: "/tmp/geth.ipc",
        isIPC: true,
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

func (ep *Endpoint) SendTransaction(transaction interface{}) (*RPCResponse, error) { 
        return ep.Request([]Parameters{transaction}, RPCendpoint["SendTransaction"])
}

// wrapper around sendTransaction to facilitate the sending
// of token
// the two arguments are value, the amount in wei to send
// from the sender of the amount
// to the address to send the amount to
func (ep *Endpoint) SendToken(amount string, from string,  to string) (*RPCResponse, error){
    tr := BuildTransaction(from, to, amount, "") 
    return ep.SendTransaction(tr)
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

func (ep *Endpoint) GetAccounts() (*RPCResponse, error) {
    return ep.Request([]Parameters{}, RPCendpoint["Accounts"])
}

func (ep *Endpoint) EstimateGas(tr Transaction) (*RPCResponse, error) {
    params := []Parameters{tr, "latest"}
    return ep.Request(params, RPCendpoint["EstimateGas"])
}

func (ep *Endpoint) GetBlockTransactionCountByNumber(blockNumber string) (*RPCResponse, error) {
    params := []Parameters{blockNumber}
    return ep.Request(params, RPCendpoint["GetBlockTransactionCountByNumber"])
}

func (ep *Endpoint) GetBlockByNumber(blockNumber string) (*RPCResponse, error) {
    params := []Parameters{blockNumber}
    return ep.Request(params, RPCendpoint["GetBlockByNumber"])
}

func (ep *Endpoint) GetProof(address string, storageKey []string, blockNumber string) (*RPCResponse, error) {
    params := []Parameters{address, storageKey, blockNumber}
    return ep.Request(params, RPCendpoint["GetProof"])
}

//APPAREMENT C EST PAS REQUIRE LE BLOCKHASH ??
func (ep *Endpoint) GetBlockTransactionCountByHash(blockHash string) (*RPCResponse, error) {
    params := []Parameters{blockHash}
    return ep.Request(params, RPCendpoint["GetBlockTransactionCountByHash"])
}

func (ep *Endpoint) FeeHistory(blockCount string, newestBlock string, rewardPercentile []float32) (*RPCResponse, error) {
    params := []Parameters{blockCount, newestBlock, rewardPercentile}
    return ep.Request(params, RPCendpoint["FeeHistory"])
}

// GetTransactionReceipt query the receipt of a transaction 
// based on the hash of the transaction.
// Upon successful completion, it return a pointer to a TransactionReceiptResponse struct
// In case of failure, return nil and an appropriate error
func (ep *Endpoint) GetTransactionReceipt(transactionHash string) (*TransactionReceiptResponse, error){
    params := []Parameters{transactionHash}
    response, err := ep.Request(params, RPCendpoint["GetTransactionReceipt"])
    if err != nil {
        return nil, err
    }

    if response.Error != nil {
        return nil, fmt.Errorf("Error in the request: %s", response.Error.Message)
    }

    resultMap, ok := response.Result.(map[string]any); if !ok {
        return nil, fmt.Errorf("Answer is not a map : %s", response.Result)
    }
    receipt := new(TransactionReceiptResponse)
    jsonData, err := json.Marshal(resultMap)
    if err != nil {
        return nil, fmt.Errorf("Error When decrypting receipt:%s. Error: %s",resultMap, err)
    }
    json.Unmarshal(jsonData, receipt)
    return receipt, nil
}

func (ep *Endpoint) Call(tr Transaction) (*RPCResponse, error) {
    params := []Parameters{tr}
    return ep.Request(params, RPCendpoint["Call"])
}

// DeployContract deploy a contract to the blockchain
// it takes the account which is deploying the contract as first argument
// The second argument is the path to the compiled contract
// it return a TransactionReceipt struct (which contain the contract address) upon successfull
// completion
// if an error occur, return nil and an error
func (ep *Endpoint) DeployContract(sender string, contractByteCodePath string) (*TransactionReceiptResponse, error){

    contractContent, err := os.ReadFile(contractByteCodePath)
    if err != nil {
        return nil, err
    }
    contractPayload := fmt.Sprintf("0x%s", contractContent)
    //create a contract happen when no receiver are present in the field
    // the 0x0 address does not work to create a contract
    tr := ContractCreationTr{
        From: sender,
        Value: "0x0",
        Input: contractPayload,
    }
    r, err := ep.SendTransaction(tr)
    if err != nil {
        return nil, err
    }
    if r.Error != nil {
        return nil, fmt.Errorf("Error in the Transaction: %s", r.Error.Message)
    }

    //type assertion needed
    result, ok := r.Result.(string); if !ok{
        return nil, fmt.Errorf("Transaction hash is not a string: %v", r)
    }

    time.Sleep(5 * time.Millisecond) //otherwise it goes too fast

    trReceipt, err := ep.GetTransactionReceipt(result)
    if err != nil {
        return nil, err
    }
    return trReceipt, nil
}
