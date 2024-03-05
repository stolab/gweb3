package rpc

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (ep *Endpoint)HttpRequest(Params []Parameters, rpcDetail RPCMethod) (*http.Response, error) {
	RPCRequest := buildRPCRequest(Params, rpcDetail)
	json, err := json.Marshal(RPCRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(rpcDetail.HTTPMethod, ep.endpoint, bytes.NewReader(json))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("accept-encoding", "*/*")
	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		return nil, err
	}

	return resp, nil
}

func buildRPCRequest(params []Parameters, method RPCMethod) (RPCTransaction){
	rpc := RPCTransaction{
		Jsonrpc: "2.0",
		Id: 66, //TODO maybe mettre un random in a given range
		Method: method.Method,
		Params: params,
	}
	return rpc
}
