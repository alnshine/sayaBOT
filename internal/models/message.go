package models

import (
	"time"
)

type Chat struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	ID       int       `json:"id" db:"id"`
	Content  string    `json:"content" db:"content"`
	Username string    `json:"username" db:"username"`
	Time     time.Time `json:"time" db:"time"`
	ChatId   int64     `json:"chat_id" db:"chat_id"`
}
