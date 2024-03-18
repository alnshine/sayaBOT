package service

import (
	"github.com/alnshine/sayaBOT/internal/models"
	"github.com/alnshine/sayaBOT/internal/repository"
)

type Message interface {
	CreateMessage(message models.Message) error
	GetMessagesForTimeInterval(chatID int64) (string, error)
}

type Service struct {
	Message
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Message: NewMessageService(repo),
	}
}
