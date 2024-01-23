package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	port, portExists := os.LookupEnv("PORT")

	if portExists {
		fmt.Println(port)
		fmt.Println(portExists)
	} else {
		port = "7000"
		fmt.Println("PORT not found. Server started on reserved PORT:", port)
	}

	http.HandleFunc("/", Handler)
	log.Println("Server is starting on PORT: " + port)
	log.Fatal(http.ListenAndServe((":" + port), nil))
}
