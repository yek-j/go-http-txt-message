package main

import (
	"go-http-txt-message/server/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/send", handlers.Send)
	mux.HandleFunc("/list/", handlers.List)
	mux.HandleFunc("/message", handlers.GetMessage)

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe Error : ", err)
	}
}