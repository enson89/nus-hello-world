package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", helloWorldEndPoint)
	fmt.Println("Server :  http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func helloWorldEndPoint(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "hello world")

}
