// ChatGPT 4 generated

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Client struct {
	ID   int
	Conn net.Conn
}

var (
	clientIDCounter int
	clients         = make(map[int]Client)
	clientsMutex    sync.Mutex
)

func handleClient(client Client) {
	defer func() {
		clientsMutex.Lock()
		delete(clients, client.ID)
		clientsMutex.Unlock()
		client.Conn.Close()
		fmt.Printf("\nClient %d disconnected\n", client.ID)
	}()

	scanner := bufio.NewScanner(client.Conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("Received from client %d: %s\n", client.ID, text)
	}
	if scanner.Err() != nil {
		fmt.Printf("Error with client %d: %v\n", client.ID, scanner.Err())
	}
}

func startTCPServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("[+] Server started on port", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		clientsMutex.Lock()
		clientIDCounter++
		client := Client{ID: clientIDCounter, Conn: conn}
		clients[client.ID] = client
		clientsMutex.Unlock()

		fmt.Printf("\n[+] Client %d connected\n", client.ID)

		go handleClient(client)
	}
}

func listClients() {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	data := [][]string{}
	for id, client := range clients {
		data = append(data, []string{fmt.Sprintf("%d", id), client.Conn.RemoteAddr().String()})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Address"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT})
	table.AppendBulk(data)
	table.Render()
}

func sendMessageToClient(clientID int, message string) {
	clientsMutex.Lock()
	client, exists := clients[clientID]
	clientsMutex.Unlock()

	if exists {
		client.Conn.Write([]byte(message + "\n"))
		fmt.Printf("Message sent to client %d\n", clientID)
	} else {
		fmt.Printf("Client %d not found\n", clientID)
	}
}

func startClientCLI() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- Client Commands ---")
		fmt.Println("1. List clients")
		fmt.Println("2. Send message/command to client")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			listClients()
		case "2":
			fmt.Print("Enter client ID: ")
			clientIDInput, _ := reader.ReadString('\n')
			clientIDInput = strings.TrimSpace(clientIDInput)
			clientID, err := strconv.Atoi(clientIDInput)
			if err != nil {
				fmt.Println("Invalid client ID")
				continue
			}

			fmt.Print("Enter message: ")
			message, _ := reader.ReadString('\n')
			message = strings.TrimSpace(message)

			sendMessageToClient(clientID, message)
		case "3":
			fmt.Println("Exiting CLI")
			return
		}
	}
}

func main() {
	go startTCPServer("8080")
	time.Sleep(2 * time.Second)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("c2>")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "client":
			startClientCLI()
		}
	}
}
