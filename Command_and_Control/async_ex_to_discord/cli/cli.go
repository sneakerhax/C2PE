package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

const c2ServerURL = "http://localhost:8080"

var httpClient = &http.Client{}

// Client represents a client with ID, IP, and interval
type Client struct {
	ID       string
	IP       string
	Interval int
}

// Command represents a command with agent ID and command string
type Command struct {
	AgentID string
	Command string
}

// FetchClients fetches the list of clients from the server
func FetchClients() ([]Client, error) {
	resp, err := httpClient.Get(fmt.Sprintf("%s/clients", c2ServerURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch clients, status code: %d", resp.StatusCode)
	}

	var clients interface{}
	if err := json.NewDecoder(resp.Body).Decode(&clients); err != nil {
		return nil, err
	}

	if str, ok := clients.(string); ok && str == "no clients" {
		return nil, nil // No clients to return
	}

	clientData, ok := clients.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	clientList := make([]Client, len(clientData))
	for i, client := range clientData {
		clientArray, ok := client.([]interface{})
		if !ok || len(clientArray) < 3 {
			return nil, fmt.Errorf("invalid client data format")
		}
		clientList[i] = Client{
			ID:       clientArray[0].(string),
			IP:       clientArray[1].(string),
			Interval: int(clientArray[2].(float64)),
		}
	}

	return clientList, nil
}

// FetchCommands fetches the list of commands from the server
func FetchCommands() ([]Command, error) {
	resp, err := httpClient.Get(fmt.Sprintf("%s/show-commands", c2ServerURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var commands [][]string
	if err := json.NewDecoder(resp.Body).Decode(&commands); err != nil {
		return nil, err
	}

	commandList := make([]Command, len(commands))
	for i, cmd := range commands {
		if len(cmd) < 2 {
			return nil, fmt.Errorf("invalid command data format")
		}
		commandList[i] = Command{
			AgentID: cmd[0],
			Command: cmd[1],
		}
	}

	return commandList, nil
}

// AddCommand sends a command to the server
func AddCommand(agentID string, commandArgs ...string) error {
	command := strings.Join(commandArgs, " ")

	data := url.Values{
		"agentId": {agentID},
		"command": {command},
	}

	resp, err := httpClient.PostForm(fmt.Sprintf("%s/add-command", c2ServerURL), data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to add command, status code: %d", resp.StatusCode)
	}

	return nil
}

// RemoveAgent deregisters an agent from the server
func RemoveAgent(agentID string) error {
	data := url.Values{
		"agentId": {agentID},
	}

	resp, err := httpClient.PostForm(fmt.Sprintf("%s/deregister", c2ServerURL), data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to remove agent, status code: %d", resp.StatusCode)
	}

	return nil
}

// printTable prints the clients or commands in a table format
func printTable(data [][]string, header []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	if len(data) == 0 {
		empty := make([]string, len(header))
		table.Append(empty)
	} else {
		for _, v := range data {
			table.Append(v)
		}
	}
	table.Render()
}

func main() {
	app := &cli.App{
		Name:  "C2 CLI",
		Usage: "Interact with the Flask C2 server",
		Commands: []*cli.Command{
			{
				Name:  "list-clients",
				Usage: "List all registered clients",
				Action: func(c *cli.Context) error {
					clients, err := FetchClients()
					if err != nil {
						return err
					}

					if clients == nil {
						fmt.Println("No clients available")
						return nil
					}

					var data [][]string
					for _, client := range clients {
						data = append(data, []string{client.ID, client.IP, fmt.Sprintf("%d", client.Interval)})
					}

					printTable(data, []string{"Client ID", "IP Address", "Interval"})
					return nil
				},
			},

			{
				Name:  "list-commands",
				Usage: "List all pending commands",
				Action: func(c *cli.Context) error {
					commands, err := FetchCommands()
					if err != nil {
						return err
					}

					var data [][]string
					for _, command := range commands {
						data = append(data, []string{command.AgentID, command.Command})
					}

					printTable(data, []string{"Agent ID", "Command"})
					return nil
				},
			},
			{
				Name:  "add-command",
				Usage: "Add a command to a specific agent",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "agent-id",
						Usage:    "The ID of the agent",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "command",
						Usage:    "The command to execute",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					agentID := c.String("agent-id")
					command := c.String("command")

					if err := AddCommand(agentID, command); err != nil {
						return err
					}

					fmt.Println("Command added successfully")
					return nil
				},
			},
			{
				Name:  "remove-agent",
				Usage: "Remove an agent by agent ID",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "agent-id",
						Usage:    "The ID of the agent to remove",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					agentID := c.String("agent-id")

					if err := RemoveAgent(agentID); err != nil {
						return err
					}

					fmt.Printf("Agent %s removed successfully\n", agentID)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
