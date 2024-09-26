// ChatGPT generated

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

		fmt.Printf("\n [+] Client %d connected\n", client.ID)

		go handleClient(client)
	}
}

func listClients() {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	fmt.Println("Connected Clients:")
	fmt.Println("ID\tAddress")
	for id, client := range clients {
		fmt.Printf("%d\t%s\n", id, client.Conn.RemoteAddr().String())
	}
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

func startCLI() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- Command Line Interface ---")
		fmt.Println("1. List clients")
		fmt.Println("2. Send message to client")
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
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func main() {
	go startTCPServer("8080")
	time.Sleep(2 * time.Second)
	startCLI() // This will drop into the CLI after the server starts
}
