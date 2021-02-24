package logic

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type User struct {
	UID            int           `json:"uid"`
	Nickname       string        `json:"nickname"`
	EnterAt        time.Time     `json:"enter_at"`
	Addr           string        `json:"addr"`
	MessageChannel chan *Message `json:"-"`

	conn *websocket.Conn
}

var globalUID uint32 = 0
var System = &User{}

func NewUser(conn *websocket.Conn, nickname, addr string) *User {
	return &User{
		UID:            int(atomic.AddUint32(&globalUID, 1)),
		Nickname:       nickname,
		Addr:           addr,
		EnterAt:        time.Now(),
		MessageChannel: make(chan *Message, 32),

		conn: conn,
	}
}

func (u *User) SendMessage(ctx context.Context) {
	for msg := range u.MessageChannel {
		wsjson.Write(ctx, u.conn, msg)
	}
}

func (u *User) CloseMessageChannel() {
	close(u.MessageChannel)
}

func (u *User) ReceiveMessage(ctx context.Context) error {
	var (
		receiveMsg map[string]string
		err        error
	)
	for {
		err = wsjson.Read(ctx, u.conn, &receiveMsg)
		if err != nil {
			//判斷連接是否已經關閉
			var closeErr websocket.CloseError
			if errors.As(err, &closeErr) {
				return nil
			}

			return err
		}

		sendMsg := NewMessage(u, receiveMsg["content"])
		Broadcaster.Broadcast(sendMsg)
	}
}
