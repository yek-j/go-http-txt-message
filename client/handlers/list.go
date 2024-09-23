package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HandlerList(baseUrl, username string) {
	resp, err := http.Get(fmt.Sprintf("%s/list/%s", baseUrl, username))
	if err != nil {
		fmt.Printf("http /list error : %v", err)
		return 
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Message List : ", string(body))
}