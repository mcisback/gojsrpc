package main

import (
	"fmt"
	"net/http"
)

func sum(params GoRpcRequestParams) (any, *GoRpcError) {
	fmt.Println("A: ", params["a"].Value)
	fmt.Println("B: ", params["b"].Value)

	var a float64 = params["a"].Value.(float64) // type assertion

	if ok, err := typeAssert(a, "float64"); !ok {
		return 0, err
	}

	var b float64 = params["b"].Value.(float64)

	if ok, err := typeAssert(b, "float64"); !ok {
		return 0, err
	}

	// Missing Type Assertion

	return int(a + b), nil
}

func concat(params GoRpcRequestParams) (any, *GoRpcError) {
	fmt.Println("A: ", params["a"].Value)
	fmt.Println("B: ", params["b"].Value)

	fmt.Println("a type: ", getType(params["a"].Value))
	fmt.Println("b type: ", getType(params["b"].Value))

	if ok, err := typeAssert(params["a"].Value, "string"); !ok {
		return "", err
	}

	a, _ := params["a"].Value.(string) // type assertion

	if ok, err := typeAssert(params["b"].Value, "string"); !ok {
		return "", err
	}

	b, _ := params["b"].Value.(string) // type assertion

	return (a + b), nil
}

func login(params GoRpcRequestParams) (any, *GoRpcError) {
	fmt.Println("User: ", params["username"].Value)
	fmt.Println("Password: ", params["password"].Value)

	fmt.Println("a type: ", getType(params["username"].Value))
	fmt.Println("b type: ", getType(params["password"].Value))

	if ok, err := typeAssert(params["username"].Value, "string"); !ok {
		return false, err
	}

	username, _ := params["username"].Value.(string) // type assertion

	if ok, err := typeAssert(params["password"].Value, "string"); !ok {
		return false, err
	}

	password, _ := params["password"].Value.(string) // type assertion

	if username == "mario" && password == "bros" {
		return true, nil
	}

	return false, nil
}

func main() {
	funcMap := GoRpcFuncMap{
		"sum":    sum,
		"concat": concat,
		"login":  login,
	}

	rpc := GoRPC{}

	rpc.start("/gorpc", &funcMap)

	fmt.Println("Starting GO RPC")

	// http.HandleFunc("/gorpc", goRpc)

	// fs := http.FileServer(http.Dir("./frontend/public"))
	// http.Handle("/", fs)

	fmt.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}

}
