package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Testing out loops")

	//for i := 0; i < 10; i++ {
	//	fmt.Println("Occurrence: ", i+1)
	//printVal(i)
	//}

	printVal(1)
	printVal(3)

	startHttpServer()
}

func printVal(a int) {
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("unsupported operation")
	}
}

func startHttpServer() {
	http.HandleFunc("/", handleReq)

	http.ListenAndServe(":80", nil)
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my website!")
}
