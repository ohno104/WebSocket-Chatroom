package server

import (
	"log"
	"net/http"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"

	"felix.bs.com/felix/BeStrongerInGO/WebSocket-Chatroom/logic"
)

func WebSocketHandleFunc(w http.ResponseWriter, req *http.Request) {
	conn, err := websocket.Accept(w, req, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Println("websocket accept error:", err)
		return
	}

	// 新使用者, 建置該使用者的實例
	nickname := req.FormValue("nickname")
	if l := len(nickname); l < 2 || l > 20 {
		log.Println("nickname illegal: ", nickname)

		wsjson.Write(req.Context(), conn, logic.NewErrorMessage("非法暱稱, 暱稱長度4-20"))
		conn.Close(websocket.StatusUnsupportedData, "nickname illegal")
		return
	}
	if !logic.Broadcaster.CanEnterRoom(nickname) {
		log.Println("暱稱已經存在:", nickname)
		wsjson.Write(req.Context(), conn, logic.NewErrorMessage("暱稱已經存在"))
		conn.Close(websocket.StatusUnsupportedData, "nickname exists")
		return
	}

	user := logic.NewUser(conn, nickname, req.RemoteAddr)

	// 開啟給使用者送訊息的goroutin
	go user.SendMessage(req.Context())

	// 給新使用者發送歡迎訊息
	user.MessageChannel <- logic.NewWelcomeMessage(nickname)

	// 向所有使用者告知新使用者加入
	msg := logic.NewNoticeMessage(nickname + " 加入了聊天室")
	logic.Broadcaster.Broadcast(msg)

	// 將使用者加入廣播器的使用者列表
	logic.Broadcaster.UserEntering(user)
	log.Println("user: ", nickname, "joins chat")

	// 接收使用者訊息
	err = user.ReceiveMessage(req.Context())

	// 使用者離開
	logic.Broadcaster.UserLeaving(user)
	msg = logic.NewNoticeMessage(user.Nickname + " 離開了聊天室")
	logic.Broadcaster.Broadcast(msg)
	log.Println("user:", nickname, "leaves chat")

	if err == nil {
		conn.Close(websocket.StatusNormalClosure, "")
	} else {
		log.Println("Read from client error:", err)
		conn.Close(websocket.StatusInternalError, "Read from client error")
	}
}
