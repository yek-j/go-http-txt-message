package handlers 

import (
	"net/http"
	"fmt"
)

func Send(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Send Handler")
}