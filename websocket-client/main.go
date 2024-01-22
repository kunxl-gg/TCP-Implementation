package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"os"
)

func establishConnection() *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080.ws", nil)
	if err != nil {
		fmt.Println(err)
	}

	return conn
}

func readMessage(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message from server:", err)
			return
		}
		fmt.Println("Received from server:", string(message))
	}
}

func readUserInput() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter your message here")
		userMessage, _ := reader.ReadString('\n')
		fmt.Println(color.GreenString(userMessage))
	}
}

func main() {
	ws := establishConnection()
	defer ws.Close()
	go readUserInput()
	fmt.Println("Hello World")
	select {}
}
