package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type User struct {
	Username       string
	MessageChannel chan *Message
	LeaveChannel   chan *User
	Connection     *websocket.Conn
}

func NewUser(Username string, MessageChannel chan *Message, LeaveChannel chan *User, Connection *websocket.Conn) *User {
	return &User{
		Username:       Username,
		MessageChannel: MessageChannel,
		LeaveChannel:   LeaveChannel,
		Connection:     Connection,
	}
}

func (u *User) ReceiveMessage() {
	for {
		if _, content, err := u.Connection.ReadMessage(); err != nil {
			fmt.Printf("there was an error: %s", err)
			break
		} else {
			message := &Message{}
			json.Unmarshal(content, &message)
			u.MessageChannel <- message
		}
	}

	u.LeaveChannel <- u
}

func (u *User) SendMessage() {

}
