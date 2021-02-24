package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:2021/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	err = c.WriteJSON("Hello WebSocket Server")
	if err != nil {
		log.Println(err)
		return
	}

	var v interface{}
	err = c.ReadJSON(&v)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("收到服務端的回應: %s\n", v)
}
