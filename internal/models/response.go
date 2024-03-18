package models

import "time"

type Response struct {
	ResponseValues `json:"response"`
}

type ResponseValues struct {
	Retelling string    `json:"retelling"`
	TimeStart time.Time `json:"time-start"`
	TimeEnd   time.Time `json:"time-end"`
	ChatID    int64     `json:"chat-id"`
}
