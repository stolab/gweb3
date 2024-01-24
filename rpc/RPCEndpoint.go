package rpc

type RPCMethod struct {
    Method string
    HTTPMethod string
}

//https://ethereum.github.io/execution-apis/api-documentation/
var RPCendpoint = map[string]RPCMethod {
    "ClientVersion": RPCMethod{"web3_clientVersion","get"},
    "Sha3" : RPCMethod{"web3_sha3", "post"},
    "NetworkId": RPCMethod{"net_id", "get"},
    "MostRecentBlock": RPCMethod{"eth_blockNumber", "get"},
    "GetBalance": RPCMethod{"eth_getBalance", "get"},
    "GetStorageAt": RPCMethod{"eth_getStorageAt", "get"},
    "GetTransactionCount": RPCMethod{"eth_getTransactionCount", "get"},
    "GetCode": RPCMethod{"eth_getCode", "get"},
    "GetBlockByHash": RPCMethod{"eth_getBlockByHash", "get"},
    //## Transaction ##
    "Sign": RPCMethod{"eth_sign", "post"},
    "SignTransaction": RPCMethod{"eth_signTransaction", "post"},
    "SendTransaction": RPCMethod{"eth_sendTransaction", "post"},
    "SendRawTransaction": RPCMethod{"eth_sendRawTransaction", "post"},
    "GasPrice": RPCMethod{"eth_gasPrice", "get"},
    "Coinbase": RPCMethod{"eth_coinbase", "get"},
    "GetTransactionByHash": RPCMethod{"eth_getTransactionByHash", "get"},
    "GetBlockReceipts": RPCMethod{"eth_getBlockReceipts", "get"},
    "GetTransactionReceipt": RPCMethod{"eth_getTransactionReceipt", "post"},
    "Call": RPCMethod{"eth_call", "post"},

    //TODO Still need to be implemented
    "Account" : RPCMethod{"eth_accounts", "get"},
    "EstimateGas": RPCMethod{"eth_estimateGas", "get"},
    "FeeHistory": RPCMethod{"eth_feeHistory", "get"},
    "GetBlockByNumber": RPCMethod{"eth_getBlockByNumber", "get"},
    "GetBlockTransactionCountByHash": RPCMethod{"eth_getBlockTransactionCountByhash", "get"},
    "GetBlockTransactionCountByNumber": RPCMethod{"eth_getBlockTransactionCountByNumber", "get"},
    "GetProof": RPCMethod{"eth_getProof", "get"},

}
