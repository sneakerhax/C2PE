# Golang TCP C2

I used this TCP client/server [example](https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/) and built a very basic C2

## Starting the server

```
go run tcp_server.go 1234
[+] Listening on port 1234
```

Starting the server

## Building the client

```
go build tcp_client.go
```

Build the client

## Usage

```
./tcp_client 127.0.0.1:1234
```

Running the client and connecting to the server

```
go run tcp_server.go 1234
[+] Listening on port 1234
>> ls
[+] Output from ls

README.md
tcp_client.go
tcp_server.go
```

Running the server and executing the ls command on the client
