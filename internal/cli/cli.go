package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/XDBerry29/chat-app-go/internal/network"
)

func StartCLI(localIP, localPort string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.SplitN(input, " ", 2)
		command := parts[0]

		switch command {
		case "connect":
			handleConnect(parts, localIP, localPort)
		case "message":
			handleMessage(parts)
		case "list":
			handleList()
		case "disconnect":
			handleDisconnect(parts)
		case "help":
			fmt.Println(" Available commands: connect, message, list, disconnect")
		default:
			fmt.Println("Unknown command. Try 'help'.")
		}
	}
}

func handleConnect(parts []string, localIP, localPort string) {
	if len(parts) < 2 {
		fmt.Println("Usage: connect <ip:port>")
		return
	}
	address := parts[1]
	client := network.ConnectToServer(address, localIP, localPort)
	if client != nil {
		connectionID := network.RegisterClient(client)
		fmt.Printf("Connected to %s as connection %d\n", address, connectionID)
	}
}

func handleMessage(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Usage: message <connection_id> <message>")
		return
	}
	msgParts := strings.SplitN(parts[1], " ", 2)
	if len(msgParts) < 2 {
		fmt.Println("Usage: message <connection_id> <message>")
		return
	}
	id := msgParts[0]
	message := msgParts[1]
	idInt := atoi(id)
	err := network.SendMessageToClient(idInt, message)
	if err != nil {
		fmt.Println("Invalid connection ID")
	}
}

func handleList() {
	connections := network.GetAllConnections()
	if len(connections) == 0 {
		fmt.Println("No active connections.")
		return
	}
	fmt.Println("Active connections:")
	for id, client := range connections {
		fmt.Printf("%d: %s\n", id, client.Conn.RemoteAddr().String())
	}
}

func handleDisconnect(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Usage: disconnect <connection_id>")
		return
	}
	id := parts[1]
	idInt := atoi(id)
	err := network.DisconnectClient(idInt)
	if err != nil {
		fmt.Println("Invalid connection ID")
	}
}

func atoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return num
}
