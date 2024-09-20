package main

import (
	"go-http-txt-message/server/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/send", handlers.Send)
	http.HandleFunc("/list/", handlers.List)
	http.HandleFunc("/message", handlers.GetMessage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe Error : ", err)
	}
}