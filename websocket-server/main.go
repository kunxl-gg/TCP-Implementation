package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world")
}

func reader(ws *websocket.Conn) {
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(message)

		err = ws.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func wsConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Established a TCP Connection")

	reader(ws)
}

func setUpRoutes() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/ws", wsConnection)
}

func main() {
	fmt.Println("Go websocket server")
	setUpRoutes()
	http.ListenAndServe(":8080", nil)
	return
}
