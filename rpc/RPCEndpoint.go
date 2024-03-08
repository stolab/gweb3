package rpc

type RPCMethod struct {
    Method string
    HTTPMethod string
}

//https://ethereum.github.io/execution-apis/api-documentation/
var RPCendpoint = map[string]RPCMethod {
    "ClientVersion": {"web3_clientVersion","GET"},
    "Sha3" : {"web3_sha3", "POST"},
    "NetworkId": {"net_id", "GET"},
    "MostRecentBlock": {"eth_blockNumber", "POST"},
    "GetBalance": {"eth_getBalance", "POST"},
    "GetStorageAt": {"eth_getStorageAt", "GET"}, //Should probably be POST?
    "GetTransactionCount": {"eth_getTransactionCount", "GET"},
    "GetCode": {"eth_getCode", "GET"},
    "GetBlockByHash": {"eth_getBlockByHash", "GET"},
    //## Transaction ##
    "Sign": {"eth_sign", "POST"},
    "SignTransaction": {"eth_signTransaction", "POST"},
    "SendTransaction": {"eth_sendTransaction", "POST"},
    "SendRawTransaction": {"eth_sendRawTransaction", "POST"},
    "GasPrice": {"eth_gasPrice", "GET"},
    "Coinbase": {"eth_coinbase", "GET"},
    "GetTransactionByHash": {"eth_getTransactionByHash", "GET"},
    "GetBlockReceipts": {"eth_getBlockReceipts", "GET"},
    "GetTransactionReceipt": {"eth_getTransactionReceipt", "POST"},
    "Call": {"eth_call", "POST"},

    //TODO Still need to be implemented
    "Account" : {"eth_accounts", "GET"},
    "EstimateGas": {"eth_estimateGas", "GET"},
    "FeeHistory": {"eth_feeHistory", "GET"},
    "GetBlockByNumber": {"eth_getBlockByNumber", "GET"},
    "GetBlockTransactionCountByHash": {"eth_getBlockTransactionCountByhash", "GET"},
    "GetBlockTransactionCountByNumber": {"eth_getBlockTransactionCountByNumber", "GET"},
    "GetProof": {"eth_getProof", "GET"},

}
