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

	// Print system information
	fmt.Println("\n========= Remote Execution Payload ========")
	fmt.Println("Time:", time.Now().Format(time.RFC1123))
	fmt.Println("Hostname:", hostname)
	fmt.Println("Username:", currentUser.Username)
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("==========================================")

	// List processes (as an example of something system-specific)
	fmt.Println("\nListing contents of current directory:")
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			continue
		}

		fmt.Printf("- %s (%d bytes)\n", file.Name(), info.Size())
	}

	fmt.Println("\nPayload execution complete!")
}
