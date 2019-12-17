package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, WebAssembly!")

	http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))
}
