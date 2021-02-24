package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"felix.bs.com/felix/BeStrongerInGO/WebSocket-Chatroom/logic"
)

func homeHandleFunc(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles(rootDir + "/template/home.html")
	if err != nil {
		fmt.Fprint(w, "範例解析錯誤")
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, "範例執行錯誤")
		return
	}
}

func userListHandleFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	userList := logic.Broadcaster.GetUserList()
	b, err := json.Marshal(userList)

	if err != nil {
		fmt.Fprint(w, `[]`)
	} else {
		fmt.Fprint(w, string(b))
	}
}
