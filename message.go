package main

type Message struct {
	SenderUser string
	Content    string
}

func NewMessage(SenderUser string, Content string) *Message {
	return &Message{
		SenderUser: SenderUser,
		Content:    Content,
	}
}
