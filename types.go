package main

import (
	"fmt"
	"net/http"
	"strings"
)

type GoRpcRequestValue struct {
	// Type  string // optional ?
	Value any
}

type GoRpcRequestParams map[string]GoRpcRequestValue

type GoRpcRequest struct {
	Method string
	Params GoRpcRequestParams
}

type GoRpcResponse struct {
	Success bool
	Data    string
}

type GoRpcError struct {
	Message string
}

type GoRpcFuncMap map[string]func(GoRpcRequestParams) (string, *GoRpcError)

type HttpFunction func(w http.ResponseWriter, r *http.Request)
type HttpRoutesMap map[string]HttpRoute
type HttpRoute map[string]HttpFunction

func getType(v any) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float64:
		return "float64"
	case float32:
		return "float32"
	case bool:
		return "bool"
	default:
		return "" // Could cause bugs if not handled correctly
	}
}

func typeAssert(v any, expectedType string) (bool, *GoRpcError) {
	if expectedType != getType(v) {
		return false, &GoRpcError{
			Message: strings.Trim(fmt.Sprintf("GoRpcRequestParams Value Interface conversion: interface is %T, not %s\n", v, expectedType), "\r\n"),
		}
	}

	return true, nil
}
