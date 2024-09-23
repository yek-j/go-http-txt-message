package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HandleMessage(baseUrl, username, title string) {
	url := fmt.Sprintf("%s/message?username=%s&title=%s", baseUrl, username, title)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("http /message error : %v", err)
		return 
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Message Content : ", string(body))
}