# Go + JS Simple Custom RPC

A go + js library custom RPC, made for study.
This is not the standard RPC implementation.

Compatible with every JS Framework

You define your function in main.go, and in AppRpc.js,
for example:

```javascript
// myApp.js

import GoJsRpc from "./GoJsRpc";

const rpc = new GoJsRpc("http://localhost:3000/gorpc")

let form = {
    a: 0,
    b: 0
}

let sum = null

async function onFormSubmit() {
	// Simple as that
	// No infinte endpoint, no infinite code, just call the method and get the result
    sum = await rpc.sum(form)

    console.log('sum: ', sum)
}
```

```golang
package main

import (
	"fmt"
	"net/http"
)

// This Method will be called by JS
// Without any REST route, everything (almost) is handled by JS and GO
// You have just to write the logic

func sum(params GoRpcRequestParams) (string, *GoRpcError) {
	fmt.Println("A: ", params["a"].Value)
	fmt.Println("B: ", params["b"].Value)

	var a float64 = params["a"].Value.(float64)

	if ok, err := typeAssert(a, "float64"); !ok {
		return "", err
	}

	var b float64 = params["b"].Value.(float64)

	if ok, err := typeAssert(b, "float64"); !ok {
		return "", err
	}

	return fmt.Sprintf("%d", int(a+b)), nil
}

func main() {
	funcMap := GoRpcFuncMap{
		"sum": sum,
	}

	routes := HttpRoutesMap{
		"POST": HttpRoute{
			"/gorpc": func(w http.ResponseWriter, r *http.Request) {
				goRpc(funcMap, w, r)
			},
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

```