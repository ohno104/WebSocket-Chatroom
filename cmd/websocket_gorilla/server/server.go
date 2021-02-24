package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "HTTP, HELLO")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) {
		c, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer func() {
			log.Println("server disconnect...")
			c.Close()
		}()

		var v interface{}
		err = c.ReadJSON(&v)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("接收到用戶端: %s\n", v)
		if err := c.WriteJSON("Hello Websocket Client"); err != nil {
			log.Println(err)
			return
		}

	})
	log.Println("Server start at :2021")
	log.Fatal(http.ListenAndServe(":2021", nil))
}
