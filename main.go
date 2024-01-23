package main

import (
	"log"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello, world!\n"))
}

func main() {
	http.HandleFunc("/", Handler)
	log.Println("Server is starting on PORT: 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
