# Go Chat Application

This is a simple command-line based chat application built using Go and WebSockets. The application allows multiple clients to connect and communicate with each other. Each client can connect to other clients, send messages, list active connections, and disconnect from any connection.

## Features

- Peer-to-peer communication using WebSockets.
- Client can connect to multiple peers.
- Send and receive messages.
- List all active connections.
- Disconnect from a specific connection.

- 
## Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/XDBerr/chat-app-go.git
    cd go-chat-app
    ```

2. **Build the application:**
    ```sh
    go build -o client ./cmd/client
    ```

## Usage

1. **Start a client:**
    ```sh
    ./client <port>
    ```
    Example:
    ```sh
    ./client 8080
    ```

2. **Commands:**

    - **connect `<ip:port>`**: Connect to another client.
        ```sh
        connect 127.0.0.1:8081
        ```

    - **message `<connection_id> <message>`**: Send a message to a connected client.
        ```sh
        message 1 Hello from 8080
        ```

    - **list**: List all active connections.
        ```sh
        list
        ```

    - **disconnect `<connection_id>`**: Disconnect from a specific connection.
        ```sh
        disconnect 1
        ```

## Example

1. **Start the first client on port 8080:**
    ```sh
    ./client 8080
    ```

2. **Start the second client on port 8081:**
    ```sh
    ./client 8081
    ```

3. **Connect the second client to the first client:**
    ```sh
    connect 127.0.0.1:8080
    ```

4. **Send a message from the second client to the first client:**
    ```sh
    message 1 Hello from 8081
    ```

5. **List active connections in the first client:**
    ```sh
    list
    ```

6. **Disconnect the first client from the second client:**
    ```sh
    disconnect 1
    ```



