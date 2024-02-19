package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GoRPC struct {
	funcMap *GoRpcFuncMap
}

func (rpc *GoRPC) start(route string, funcMap *GoRpcFuncMap) {

	rpc.funcMap = funcMap

	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {

		if !corsMiddleware(w, r) {
			return
		}

		if r.Method != http.MethodPost {
			fmt.Println("Method not allowed", r.Method)

			w.WriteHeader(http.StatusMethodNotAllowed)

			return
		}

		rpc.handleRPC(w, r)
	})
}

func (rpc *GoRPC) handleRPC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New GoRpc Request: ", r.Method)

	reqBody, e := io.ReadAll(r.Body)

	fmt.Println("reqBody: ", string(reqBody))

	if e != nil {
		fmt.Println("reqBody error: ", e)
	}

	var goRpcRequest GoRpcRequest
	json.Unmarshal(reqBody, &goRpcRequest)

	fmt.Println("JSON REQUEST: ", goRpcRequest.Method)
	fmt.Println("JSON REQUEST: ", goRpcRequest.Params)

	handler := *rpc.funcMap

	value, err := handler[goRpcRequest.Method](goRpcRequest.Params)
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
