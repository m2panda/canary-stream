package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	if err != nil {
		log.Printf("Error launching serv %v", err)
	}
}
