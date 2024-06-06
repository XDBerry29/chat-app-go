package network

import (
	"fmt"
	"sync"
)

type Hub struct {
	clients      map[int]*Client
	connectionID int
	mu           sync.Mutex
}

var hub = &Hub{
	clients: make(map[int]*Client),
}

func RegisterClient(client *Client) int {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	hub.connectionID++
	hub.clients[hub.connectionID] = client
	fmt.Printf("Registered new connection from %s as connection %d\n", client.Conn.RemoteAddr().String(), hub.connectionID)
	return hub.connectionID
}

func UnregisterClient(client *Client) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	for id, c := range hub.clients {
		if c == client {
			delete(hub.clients, id)
			fmt.Printf("Disconnected client %d (%s)\n", id, client.Conn.RemoteAddr().String())
			break
		}
	}
}

func GetClient(id int) (*Client, bool) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	client, ok := hub.clients[id]
	return client, ok
}

func GetAllConnections() map[int]*Client {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	copy := make(map[int]*Client)
	for id, client := range hub.clients {
		copy[id] = client
	}
	return copy
}

func BroadcastMessage(msg string) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	for _, client := range hub.clients {
		client.Outgoing <- msg
	}
}

func SendMessageToClient(id int, msg string) error {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	client, ok := hub.clients[id]
	if !ok {
		return fmt.Errorf("client not found")
	}
	client.Outgoing <- msg
	return nil
}

func DisconnectClient(id int) error {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	client, ok := hub.clients[id]
	if !ok {
		return fmt.Errorf("client not found")
	}
	client.Conn.Close()
	delete(hub.clients, id)
	fmt.Printf("Disconnected client %d\n", id)
	return nil
}
