package rpc

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

type RPCError struct {
    Code int `json:"code"`
    Message string `json:"message"`
}

type RPCResponse struct {
    Jsonrpc string `json:"jsonrpc"`
    Id int `json:"id"`
    Result string `json:"result"`
    Error *RPCError `json:"error"`
}

//TODO should probably return a struct representing the response
func (ep *Endpoint) Request(Params []Parameters, rpcDetail RPCMethod) (*RPCResponse, error) {
    RPCRequest := buildRPCRequest(Params, rpcDetail)
    json, err := json.Marshal(RPCRequest)
    if err != nil {
        return nil, err
    }

    if ep.isIPC {
        return ep.UnixSocketRequest(json)
    } else {
        return ep.HttpRequest(rpcDetail.HTTPMethod, json)
    }
}

/*
* Used when trying to dialog with the node via UnixSocket
* Return an RPCResponse structure with the structure or an error if something went wrong
* note that if the RPC request went wrong, this function return a RPCResponse with an RPCERROR in it
*/
func (ep *Endpoint) UnixSocketRequest(RPCjson []byte) (*RPCResponse, error) {
    response := new(RPCResponse)

    conn, err := net.Dial("unix", ep.endpoint)
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    _, err = conn.Write(RPCjson)
    if err != nil {
        return nil, err
    }

    //read response
    //use of a scanner in case the response is very large.
    //The json send as a response is always on one line.
    //We consider that no more than one line is received for each answer
    scanner := bufio.NewScanner(conn)
    scanner.Scan()
    line := scanner.Text()
    err = json.Unmarshal([]byte(line), response)
    if err != nil {
        return nil, fmt.Errorf("Error when Unmarshalling the response : %s. Response: %v", err, response)
    }
    return response, nil 
}

/*
* Used when connecting to the endpoint via http
* return an RPCResponse if everything goes well and nil and an error otherwise
*/
func (ep *Endpoint) HttpRequest(httpMethod string, RPCjson []byte) (*RPCResponse, error) {

    req, err := http.NewRequest(httpMethod, ep.endpoint, bytes.NewBuffer(RPCjson))
    if err != nil {
        return nil, err
    }

    req.Header.Add("Content-Type", "application/json")
    resp, err := http.DefaultClient.Do(req)
    if err != nil{
        return nil, err
    }

    defer resp.Body.Close()
    bodybytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    rpcResponse := new(RPCResponse)
    json.Unmarshal(bodybytes, rpcResponse)

    return rpcResponse, nil
}

func buildRPCRequest(params []Parameters, method RPCMethod) (RPCTransaction){
    rpc := RPCTransaction{
        Jsonrpc:"2.0",
        Id:66, //TODO maybe mettre un random in a given range
        Method:method.Method,
        Params:params,
    }
    return rpc
}
