package main

import (
	"fmt"
	"log"
	"net/http"

	"felix.bs.com/felix/BeStrongerInGO/WebSocket-Chatroom/server"
)

var (
	addr   = ":2021"
	banner = `ChatRoom, start on: %s`
)

func main() {
	fmt.Printf(banner+"\n", addr)
	server.RegisterHandle()
	log.Fatal(http.ListenAndServe(addr, nil))
}
