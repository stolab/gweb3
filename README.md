# gweb3
Golang library to interact with an ethereum RPC endpoint.

***

## Description
By importing package present in this module, you can interact with the ethereum blockchain directly from your go application.
By using the [rpc](./rpc) package, you can directly interact with RPC method exposed by the endpoint. 
With the [contracts](./contracts) package you can directly interact with a smart contract given that you provide the ABI of the contract.

## Installation
To use one package of this module, simply import it in your project like any other go package.
```go
import github.com/stolab/gweb3/contracts
```
If that's the first time you use this module, you will also need to download this package with the following command
`go get github.com/stolab/gweb3`

## Usage
You have support for quite a lot of RPC call defined in the [Ethereum JSON-RPC specification](https://ethereum.github.io/execution-apis/api-documentation/).
You first need to create your endpoint (ie: your node that act as your entry point to the blockchain)
You will have to provide the address of the RPC in the following schema:
```bash
SCHEME://HOST:[PORT]
```
Note that if you are running a node locally, you can use the IPC file when creating the endpoint.
To create a new endpoint based on an IPC file, just provide the path to the IPC file when creating the endpoint.

where:
* `SCHEME`: is one of HTTP or HTTPS
* `HOST`: is a hostname or an IP
* `PORT`: is optional and is the port number

and then you can call whatever JSON-RPC that is implemented:
```go
endpoint := ConnectEndpoint("http://MysuperEndpoint.com:1717")
response, err := endpoint.ClientVersion() //Will return the clientVersion of the endpoint
```

All the current RPC call implemented are described in the following table
| | |
|-|-|
| ClientVersion | v |
| Sha3          | v |
| NetworkId     | v |
| MostRecentBlock | v |
| GetBalance    | v |
| GetStorageAt  | v |
| GetTransactionCount | v |
| GetCode       | v |
| GetBlockByHash | v |
| Sign          | v |
| SignTransaction | v |
| SendTransaction | v |
| SendRawTransaction | v |
| GasPrice | v |
| Coinbase | v |
| GetTransactionByHash | v |
| GetBlockReceipts | v |
| GetTransactionReceipt | v |
| Call | v |

### Smart Contract Interaction
You can also interact with smart contract. As shown in the example below, you will need the ABI of the contract in order to interact with it, as well as an endpoint.
Once all of this is setup, you can query any of the function defined in the ABI as shown in the example below
```go
endpoint := ConnectEndpoint($ENDPOINT_URL/IP, $ENDPOINT_PORT)
contract, err := InitializeContract(endpoint, $CONTRACT_ADDR, $FROM_ADDR, $ABI)
//calling a contract function named HelloWorld which take 1 argument
contract.Function["HelloWorld"].Call("MySuperName")
```

> There is no wallet support at this time. If you want to sign a transaction, you have to use a node with an unlocked account on it. 
