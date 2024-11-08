package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user/profile", handleClientProfile)
	log.Println("Server is running in PORT 8080...☺️")
	log.Fatal(http.ListenAndServe( ":8080", nil))
}
