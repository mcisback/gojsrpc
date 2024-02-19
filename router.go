package main

import (
	"fmt"
	"net/http"
)

func registerRoute(routes HttpRoutesMap, route string, httpMethod string) {
	methodMap := map[string]string{
		"POST":    http.MethodPost,
		"GET":     http.MethodGet,
		"OPTIONS": http.MethodOptions,
		"DELETE":  http.MethodDelete,
		"CONNECT": http.MethodConnect,
		"PATCH":   http.MethodPatch,
		"PUT":     http.MethodPut,
		"HEAD":    http.MethodHead,
	}

	fmt.Println("Registering route: ", httpMethod, route, "http.Method:", methodMap[httpMethod])

	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {

		if !corsMiddleware(w, r) {
			return
		}

		if r.Method != httpMethod {
			fmt.Println("In Route: ", httpMethod, route)
			fmt.Println("Method not allowed", r.Method)

			w.WriteHeader(http.StatusMethodNotAllowed)

			return
		}

		routes[httpMethod][route](w, r)
	})
}

// Randomicamente non funziona
func dispatchRoutes(routes HttpRoutesMap) {

	for httpMethod := range routes {
		fmt.Println("Loading Route for method: ", httpMethod)

		// fmt.Println("Routes: ", routes[httpMethod])

		for route := range routes[httpMethod] {
			registerRoute(routes, route, httpMethod)
		}
	}
}
