package models

import (
	"time"
)

type Chat struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	ID       int       `json:"id"`
	Content  string    `json:"content"`
	Username string    `json:"username"`
	Time     time.Time `json:"time"`
	ChatId   int64     `json:"chat_id"`
}
