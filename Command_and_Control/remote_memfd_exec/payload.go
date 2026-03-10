// Claude 3.7 Sonnet generated
package main

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"time"
)

func main() {
	// Get system information
	hostname, _ := os.Hostname()
	currentUser, _ := user.Current()
	OS := runtime.GOOS
	arch := runtime.GOARCH

	// Print system information
	fmt.Println("[+] Executing remote payload in memory using memfd_exec")
	fmt.Println("\n========= System Information ========")
	fmt.Println("Time:", time.Now().Format(time.RFC1123))
	fmt.Println("Hostname:", hostname)
	fmt.Println("Username:", currentUser.Username)
	fmt.Println("OS:", OS)
	fmt.Println("Architecture:", arch)
	fmt.Println("==========================================")
}
