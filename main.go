package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server is running in PORT 8080...☺️")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
