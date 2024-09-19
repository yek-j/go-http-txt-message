package main

import (
	"log"
	"net/http"

	"go-http-txt-message/server/handlers"
)

func main() {
	http.HandleFunc("/send", handlers.Send)
	http.HandleFunc("/list/", handlers.List)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe Error : ", err)
	}
}