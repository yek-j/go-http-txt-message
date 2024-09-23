package handlers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Message struct {
	Recipient string `json:"Recipient"`
	Sender    string `json:"Sender"`
	Title     string `json:"Title"`
	Content   string `json:"Content"`
}

func HandleSend(baseUrl string) {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Recipient? >> ")
	recipient, _ := reader.ReadString('\n')
	recipient = strings.TrimSpace(recipient)

	fmt.Print("Sender? >> ")
	sender, _ := reader.ReadString('\n')
	sender = strings.TrimSpace(sender)

	fmt.Print("Title? >> ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Content? >> ")
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	msg := Message{
		Recipient: recipient,
		Sender: sender,
		Title: title,
		Content: content,
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("JSON 마샬링 에러 : %v", err)
		return
	}

	resp, err := http.Post(baseUrl + "send", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("http /send 오류 : %v", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response: ", string(body))
}