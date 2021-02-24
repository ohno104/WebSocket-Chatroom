package server

import (
	"net/http"
	"os"
	"path/filepath"

	"felix.bs.com/felix/BeStrongerInGO/WebSocket-Chatroom/logic"
)

var rootDir string

func RegisterHandle() {
	inferRootDir()

	go logic.Broadcaster.Start()

	http.HandleFunc("/", homeHandleFunc)
	http.HandleFunc("/user_list", userListHandleFunc)
	http.HandleFunc("/ws", WebSocketHandleFunc)

}

func inferRootDir() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var infer func(d string) string
	infer = func(d string) string {
		if exists(d + "/template") {
			return d
		}
		return infer(filepath.Dir(d))
	}

	rootDir = infer(cwd)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
