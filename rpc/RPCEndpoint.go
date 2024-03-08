package rpc

type RPCMethod struct {
    Method string
    HTTPMethod string
}

//https://ethereum.github.io/execution-apis/api-documentation/
var RPCendpoint = map[string]RPCMethod {
    "ClientVersion": RPCMethod{"web3_clientVersion","GET"},
    "Sha3" : RPCMethod{"web3_sha3", "POST"},
    "NetworkId": RPCMethod{"net_id", "GET"},
    "MostRecentBlock": RPCMethod{"eth_blockNumber", "POST"},
    "GetBalance": RPCMethod{"eth_getBalance", "POST"},
    "GetStorageAt": RPCMethod{"eth_getStorageAt", "GET"},
    "GetTransactionCount": RPCMethod{"eth_getTransactionCount", "GET"},
    "GetCode": RPCMethod{"eth_getCode", "GET"},
    "GetBlockByHash": RPCMethod{"eth_getBlockByHash", "GET"},
    //## Transaction ##
    "Sign": RPCMethod{"eth_sign", "POST"},
    "SignTransaction": RPCMethod{"eth_signTransaction", "POST"},
    "SendTransaction": RPCMethod{"eth_sendTransaction", "POST"},
    "SendRawTransaction": RPCMethod{"eth_sendRawTransaction", "POST"},
    "GasPrice": RPCMethod{"eth_gasPrice", "GET"},
    "Coinbase": RPCMethod{"eth_coinbase", "GET"},
    "GetTransactionByHash": RPCMethod{"eth_getTransactionByHash", "GET"},
    "GetBlockReceipts": RPCMethod{"eth_getBlockReceipts", "GET"},
    "GetTransactionReceipt": RPCMethod{"eth_getTransactionReceipt", "POST"},
    "Call": RPCMethod{"eth_call", "POST"},

    //TODO Still need to be implemented
    "Account" : RPCMethod{"eth_accounts", "GET"},
    "EstimateGas": RPCMethod{"eth_estimateGas", "GET"},
    "FeeHistory": RPCMethod{"eth_feeHistory", "GET"},
    "GetBlockByNumber": RPCMethod{"eth_getBlockByNumber", "GET"},
    "GetBlockTransactionCountByHash": RPCMethod{"eth_getBlockTransactionCountByhash", "GET"},
    "GetBlockTransactionCountByNumber": RPCMethod{"eth_getBlockTransactionCountByNumber", "GET"},
    "GetProof": RPCMethod{"eth_getProof", "GET"},

}
