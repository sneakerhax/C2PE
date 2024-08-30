// ChatGPT generated

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Message represents the JSON payload sent to Discord
type Message struct {
	Content string `json:"content"`
}

func main() {
	webhookURL := "" // Replace with your actual webhook URL
	messageContent := "Hello Discord"

	message := Message{Content: messageContent}

	// Convert the message struct to JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		os.Exit(1)
	}

	// Send the POST request
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Discord returns 204 No Content for a successful request
	if resp.StatusCode == http.StatusNoContent {
		fmt.Println("Message sent successfully!")
	} else {
		fmt.Println("Error: Received unexpected response status:", resp.Status)
	}
}
