# Basic C2 cradle (http request)

**In Development**

## Building server with Docker

```docker build -t c2server .```

## Running server with Docker

```docker run -p 80:8080 -it c2server```

## Building the client

```go build client.go```

## Sending a command to the server (with curl)

```curl -X POST -d "agentId=<agentId>&command=id" http://127.0.0.1/add-command```

