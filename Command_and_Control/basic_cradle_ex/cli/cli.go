//ChatGPT generated

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

const baseURL = "http://localhost:8080"

func listClients() {
	resp, err := http.Get(baseURL + "/clients")
	if err != nil {
		color.Red("Error: %v", err)
		return
	}
	defer resp.Body.Close()

	var clients [][]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&clients); err != nil {
		color.Red("Error decoding response: %v", err)
		return
	}

	if len(clients) == 0 {
		color.Yellow("No clients found.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Client ID", "IP Address"})

	for i, client := range clients {
		clientID, ipAddress := client[0].(string), client[1].(string)
		table.Append([]string{fmt.Sprintf("%d", i+1), clientID, ipAddress})
	}

	color.Green("Connected Clients:")
	table.Render()
}

func addCommand(agentId string, commandArgs []string) {
	command := strings.Join(commandArgs, " ")

	form := url.Values{}
	form.Add("agentId", agentId)
	form.Add("command", command)

	resp, err := http.PostForm(baseURL+"/add-command", form)
	if err != nil {
		color.Red("Error: %v", err)
		return
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	color.Green("Response: %s", strings.TrimSpace(buf.String()))
}

func showCommands() {
	resp, err := http.Get(baseURL + "/show-commands")
	if err != nil {
		color.Red("Error: %v", err)
		return
	}
	defer resp.Body.Close()

	var commands [][]string
	if err := json.NewDecoder(resp.Body).Decode(&commands); err != nil {
		color.Red("Error decoding response: %v", err)
		return
	}

	if len(commands) == 0 {
		color.Yellow("No commands found.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Agent ID", "Command"})

	for i, command := range commands {
		agentID, cmd := command[0], command[1]
		table.Append([]string{fmt.Sprintf("%d", i+1), agentID, cmd})
	}

	color.Green("Pending Commands:")
	table.Render()
}

func main() {
	if len(os.Args) < 2 {
		color.Red("Usage: cli <command> [arguments]")
		return
	}

	switch os.Args[1] {
	case "list":
		listClients()
	case "add":
		if len(os.Args) < 4 {
			color.Red("Usage: cli add <agentId> <command> [args...]")
			return
		}
		agentId := os.Args[2]
		commandArgs := os.Args[3:]
		addCommand(agentId, commandArgs)
	case "show":
		showCommands()
	default:
		color.Red("Unknown command: %s", os.Args[1])
	}
}
