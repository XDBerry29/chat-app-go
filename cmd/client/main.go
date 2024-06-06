package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/XDBerry29/chat-app-go/internal/cli"
	"github.com/XDBerry29/chat-app-go/internal/network"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-chat-app <port>")
		os.Exit(1)
	}

	port := os.Args[1]

	localIP, err := getLocalIP()
	if err != nil {
		log.Fatalf("Error getting local IP address: %v", err)
	}
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Starting server on %s:%s\n", localIP, port)

	go network.StartServer(port)

	cli.StartCLI(localIP, port)
}

func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// Check the address type and if it is not a loopback then return it
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no valid local IP address found")
}
