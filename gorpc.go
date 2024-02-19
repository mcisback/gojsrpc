package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func goRpc(funcMap GoRpcFuncMap, w http.ResponseWriter, r *http.Request) {
	fmt.Println("New GoRpc Request: ", r.Method)

	// if r.Method != http.MethodPost {
	// 	fmt.Println("Method not allowed", r.Method)

	// 	return
	// }

	reqBody, e := io.ReadAll(r.Body)

	fmt.Println("reqBody: ", string(reqBody))

	if e != nil {
		fmt.Println("reqBody error: ", e)
	}

	var goRpcRequest GoRpcRequest
	json.Unmarshal(reqBody, &goRpcRequest)

	fmt.Println("JSON REQUEST: ", goRpcRequest.Method)
	fmt.Println("JSON REQUEST: ", goRpcRequest.Params)

	value, err := funcMap[goRpcRequest.Method](goRpcRequest.Params)

	var response *GoRpcResponse

	if err != nil {
		response = &GoRpcResponse{
			Success: false,
			Data:    err.Message,
		}
	} else {
		response = &GoRpcResponse{
			Success: true,
			Data:    value,
		}
	}

	json.NewEncoder(w).Encode(response)
}
