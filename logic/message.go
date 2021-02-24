package logic

import (
	"time"
)

type Message struct {
	User    *User            `json:"user"`
	Type    int              `json:"type"`
	Content string           `json:"content"`
	MsgTime time.Time        `json:"msg_time"`
	Users   map[string]*User `json:"users"`
}

const (
	MsgTypeNormal = iota
	MsgTypeSystem
	MsgTypeError
	MsgTypeUserList //發送目前使用者列表
)

func NewErrorMessage(content string) *Message {
	return &Message{
		User:    System,
		Type:    MsgTypeError,
		Content: content,
		MsgTime: time.Now(),
	}
}

func NewMessage(user *User, content string) *Message {
	message := &Message{
		User:    user,
		Type:    MsgTypeNormal,
		Content: content,
		MsgTime: time.Now(),
	}
	return message
}

func NewWelcomeMessage(nickname string) *Message {
	return &Message{
		User:    System,
		Type:    MsgTypeSystem,
		Content: "歡迎 " + nickname + " 加入聊天室",
		MsgTime: time.Now(),
	}
}

func NewNoticeMessage(content string) *Message {
	return &Message{
		User:    System,
		Type:    MsgTypeSystem,
		Content: content,
		MsgTime: time.Now(),
	}
}
