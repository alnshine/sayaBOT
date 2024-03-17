package models

import "time"

type Response struct {
	Retelling string    `json:"retelling"`
	TimeStart time.Time `json:"time-start"`
	TimeEnd   time.Time `json:"time-end"`
	ChatID    int       `json:"chat-id"`
}
