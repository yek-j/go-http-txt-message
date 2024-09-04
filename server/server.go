package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/yek-j/go-http-txt-message/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "임시 서버")
    })

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe Error : ", err)
	}
}