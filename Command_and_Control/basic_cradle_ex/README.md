# Basic C2 cradle (http)

**In Development**

## Building server with Docker

```
docker build -t c2server .
```

## Running server with Docker

```
docker run -p 80:8080 -it c2server
```

## Building the client

```
go build client.go
```

## Sending a command to the server (with curl)

```
curl -X POST -d "agentId=<agentId>&command=id" http://127.0.0.1/add-command
```

## Using the cli

```
./cli list
Connected Clients:
+---+-----------+--------------+
| # | CLIENT ID |  IP ADDRESS  |
+---+-----------+--------------+
| 1 | 9ZPOX7    | 192.168.65.1 |
+---+-----------+--------------+
```

List all clients conneted to the server. Currently supports list (list connected clients), show (show commands), add (Add commmand for client)