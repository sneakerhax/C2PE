# Basic C2 cradle (http)

**In Development**

## Building server with Docker

```
docker build -t c2server .
```

## Running server with Docker

```
docker run -p 8080:8080 -it c2server
docker run --name c2server -p 8080:8080 -it c2server
[2024-08-30 20:38:40 +0000] [1] [INFO] Starting gunicorn 23.0.0
[2024-08-30 20:38:40 +0000] [1] [INFO] Listening at: http://0.0.0.0:8080 (1)
[2024-08-30 20:38:40 +0000] [1] [INFO] Using worker: sync
[2024-08-30 20:38:40 +0000] [7] [INFO] Booting worker with pid: 7
ZRZ04A 192.168.65.1 has connected with interval 60
```

## Building the client

```
go build client.go
```

## Sending a command to the server (with curl)

```
curl -X POST -d "agentId=<agentId>&command=id" http://127.0.0.1/add-command
```

## Using the cli (ChatGPT generated)

Currently supports list-clients (list connected clients), list-commands (show commands), add-command (Add commmand for client)

```
./cli list-clients
+-----------+--------------+----------+
| CLIENT ID |  IP ADDRESS  | INTERVAL |
+-----------+--------------+----------+
|  ZRZ04A   | 192.168.65.1 |    60    |
|  Y5VVRH   | 192.168.65.2 |    60    |
+-----------+--------------+----------+
```

List all agents connected to the server. 

```
./cli add-command --agent-id Y5VVRH --command "curl sneakerhax.com"
Command added successfully
```

Add command to be executed by an agent/s

```
./cli.go list-commands                                               
+----------+---------------------+
| AGENT ID |       COMMAND       |
+----------+---------------------+
|  Y5VVRH  | curl sneakerhax.com |
+----------+---------------------+
```

List all commands to be executed by an agent/s
