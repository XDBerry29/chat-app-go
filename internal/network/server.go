package network

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartServer(port string) {
	http.HandleFunc("/ws", handleConnections)
	address := fmt.Sprintf(":%s", port)
	// log.Printf("Server started on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		return
	}
	defer ws.Close()

	client := NewClient(ws)
	fmt.Printf("New connection from %s\n", ws.RemoteAddr().String())
	client.ID = RegisterClient(client)
	client.Listen()
	UnregisterClient(client)
}
