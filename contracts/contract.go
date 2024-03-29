//Package contracts provide a way to interact with a smart contract
package contracts

import (
	"fmt"
	"io"

    "github.com/stolab/gweb3/rpc"

	gethABI "github.com/ethereum/go-ethereum/accounts/abi"
)

//ContractFunction Structure is used to represent
// the function a contract might have. 
type ContractFunction struct {
    Name string
    contract *Contract
}

//Contract structure represent a contract with different 
//useful attribute
type Contract struct {
    ep *rpc.Endpoint
    contractAddr string
    Abi gethABI.ABI
    Function map[string]ContractFunction
    From string
}

//Call can only be used on a ContractFunction struct.
//It will simply called the ContractFunction cf with the 
//undefined number of argument the user will give to it.
func (cf ContractFunction) Call(arguments ...interface{}) (*rpc.RPCResponse, error){
    encInput, err := cf.contract.Abi.Pack(cf.Name, arguments...)
    if err != nil {
        return nil,err
    }

    Transaction := rpc.BuildTransaction(cf.contract.From, cf.contract.contractAddr, "0x0", fmt.Sprintf("0x%x", encInput)) 

    var resp *rpc.RPCResponse
    //return the encoded value
    if cf.contract.Abi.Methods[cf.Name].StateMutability == "view" || cf.contract.Abi.Methods[cf.Name].StateMutability == "pure"{
        resp, err = cf.contract.ep.Call(Transaction) //No gas required
    } else {
        resp, err = cf.contract.ep.SendTransaction(Transaction) //gas required
    }

    if err != nil {
        return nil,err
    }

    return resp,nil
}

//Use to Initialize a new contract
// Multiple parameters are needed:
// 1) the endpoint structure
// 2) the contract address
// 3) the sendingAddr: the address used to send transaction to this contract
// 4) the ABI definition of the contract
func InitalizeContract(ep *rpc.Endpoint, contractAddr string, sendingAddr string, abiDefinition io.Reader) (*Contract, error) { 
    contract := new(Contract)
    contract.contractAddr = contractAddr
    contract.ep = ep
    contract.From = sendingAddr

    ABI, err := gethABI.JSON(abiDefinition)
    if err != nil {
        return nil, err
    }
    contract.Abi = ABI

    contract.Function = make(map[string]ContractFunction)
    for key := range ABI.Methods {
        contract.Function[key] = ContractFunction{
            Name: key,
            contract: contract,
        }
    }

    return contract, nil
}
