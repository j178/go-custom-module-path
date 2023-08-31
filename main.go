package main

import (
	"fmt"
	"net/http"

	"go.j178.dev/go-custom-module-path/api"
)

func main() {
	fmt.Println("Hello from go.j178.dev/go-custom-module-path")
	http.ListenAndServe(":8888", http.HandlerFunc(api.Index))
}
