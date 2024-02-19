package main

import (
	"fmt"
	"net/http"
)

func sum(params GoRpcRequestParams) (string, *GoRpcError) {
	fmt.Println("A: ", params["a"].Value)
	fmt.Println("B: ", params["b"].Value)

	var a float64 = params["a"].Value.(float64) // type assertion

	if ok, err := typeAssert(a, "float64"); !ok {
		return "", err
	}

	var b float64 = params["b"].Value.(float64)

	if ok, err := typeAssert(b, "float64"); !ok {
		return "", err
	}

	// Missing Type Assertion

	return fmt.Sprintf("%d", int(a+b)), nil
}

func concat(params GoRpcRequestParams) (string, *GoRpcError) {
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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello e basta")
}

func main() {
	funcMap := GoRpcFuncMap{
		"sum":    sum,
		"concat": concat,
	}

	routes := HttpRoutesMap{
		"POST": HttpRoute{
			"/gorpc": func(w http.ResponseWriter, r *http.Request) {
				goRpc(funcMap, w, r)
			},
		},
		"GET": HttpRoute{
			"/":      homePage,
			"/hello": hello,
		},
	}

	fmt.Println("Starting GO RPC")

	// http.HandleFunc("/", dispatchHttp)

	dispatchRoutes(routes)

	// http.HandleFunc("/gorpc", goRpc)

	// fs := http.FileServer(http.Dir("./frontend/public"))
	// http.Handle("/", fs)

	fmt.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}

}
