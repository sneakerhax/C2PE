package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("hostname:", name)
	fmt.Println("username:", user.Username)
}
