package network

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Incoming chan string
	Outgoing chan string

	ID int
}

func ConnectToServer(address, localIP, localPort string) *Client {
	u := url.URL{Scheme: "ws", Host: address, Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Printf("Dial error: %v", err)
		return nil
	}

	client := NewClient(conn)
	go client.Listen()

	// Send the local listening address as part of the initial handshake
	// handshakeMessage := fmt.Sprintf("HANDSHAKE %s:%s", localIP, localPort)
	// client.SendMessage(handshakeMessage)

	return client
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Conn:     conn,
		Incoming: make(chan string),
		Outgoing: make(chan string),
	}
}

func (c *Client) Listen() {
	go c.readLoop()
	go c.writeLoop()

	for {
		select {
		case msg := <-c.Incoming:
			currtime := time.Now()

			fmt.Printf("[%s] Message from %s: %s\n", currtime.Format(time.DateTime), c.Conn.RemoteAddr().String(), msg)
		}
	}
}

func (c *Client) readLoop() {
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			// Log clean disconnection message
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Read error: %v", err)
			} else {
				log.Printf("Client disconnected from %s", c.Conn.RemoteAddr().String())
				DisconnectClient(c.ID)
			}
			c.Conn.Close()
			break
		}
		message := string(msg)
		if strings.HasPrefix(message, "HANDSHAKE") {
			fmt.Println("Handshake message received:", message)
			// Handle handshake message if necessary
		} else {
			c.Incoming <- message
		}
	}
}

func (c *Client) writeLoop() {
	for {
		select {
		case msg := <-c.Outgoing:
			err := c.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				// Log clean disconnection message
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Write error: %v", err)
				} else {
					log.Printf("Client blah disconnected from %s", c.Conn.RemoteAddr().String())
					DisconnectClient(c.ID)
				}
				c.Conn.Close()
				break
			}
		}
	}
}

func (c *Client) SendMessage(msg string) {
	c.Outgoing <- msg
}
