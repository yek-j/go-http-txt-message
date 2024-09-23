package main

import (
	"flag"
	"fmt"
	"go-http-txt-message/client/handlers"
	"os"
)

func main() {
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	messageCmd := flag.NewFlagSet("message", flag.ExitOnError)

	listUser := listCmd.String("u", "", "Username for list")
	messageUser := messageCmd.String("u", "", "Username for message")
	messageTitle := messageCmd.String("t", "", "Title for message")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'send', 'list', 'message' subcommands")
		os.Exit(1)
	}

	baseURL := os.Getenv("APP_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	
	switch os.Args[1] {
	case "send":
		sendCmd.Parse(os.Args[2:])
		handlers.HandleSend(baseURL)
	case "list":
		listCmd.Parse(os.Args[2:])
		if *listUser == "" {
			fmt.Println("Please provide a username with -u flag")
			os.Exit(1)
		}
		handlers.HandlerList(baseURL, *listUser)
	case "message":
		messageCmd.Parse(os.Args[2:])
		if *messageUser == "" || *messageTitle == "" {
			fmt.Println("Please provide both username (-u) and title (-t)")
			os.Exit(1)
		}
		handlers.HandleMessage(baseURL, *messageUser, *messageTitle)
	default:
		fmt.Println("Expected 'send', 'list', 'message subcommands")
		os.Exit(1)
	}
}