# gweb3
Golang library to interact with the ethereum network.

***

## Description
By importing package present in this module, you can interact with the ethereum blockchain directly from your go application with a level of abstraction. You don't have to reimplement everything to interact with the blockchain.

## Installation
To use this package, simply import it in your project like any other go package.
```go
import github.com/stolab/gweb3
```
If that's the first time you use this library, you will also need to get this package with the following command
`go get github.com/stolab/gweb3`

## Usage
You have support for quite a lot of RPC call defined in the [Ethereum JSON-RPC specification](https://ethereum.github.io/execution-apis/api-documentation/).
You first need to create your endpoint (ie: your node that act as your entry point to the blockchain)
and then you can call whatever JSON-RPC that is implemented:
```go
endpoint := ConnectEndpoint($ENDPOINT_URL/IP, $ENDPOINT_PORT)
response, err := endpoint.ClientVersion() //Will return the clientVersion of the endpoint
```
### Smart Contract Interaction
You can also interact with smart contract. As shown in the example below, you will need the ABI of the contract in order to interact with it, as well as an endpoint.
Once all of this is setup, you can query any of the function defined in the ABI as shown in the example below
```go
endpoint := ConnectEndpoint($ENDPOINT_URL/IP, $ENDPOINT_PORT)
contract, err := InitializeContract(endpoint, $CONTRACT_ADDR, $FROM_ADDR, $ABI)
//calling a contract function named HelloWorld which take 1 argument
contract.Function["HelloWorld"].Call("MySuperName")
```

>  If you have to create and actual transaction (ie: not calling a view/pure function), you have to have your own node running in which you have unlock the sender adress. (In geth this is achieve with the `--unlock` paramters)
