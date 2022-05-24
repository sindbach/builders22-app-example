package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!")
}

func main() {
	port := ":8080"
	http.HandleFunc("/api/v1/hello", handler)
	fmt.Println("Server is running on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
