package service

import (
	"time"

	"github.com/alnshine/sayaBOT/internal/communication"
	"github.com/alnshine/sayaBOT/internal/models"
	"github.com/alnshine/sayaBOT/internal/repository"
)

type MessageService struct {
	repo repository.Message
}

func NewMessageService(repo repository.Message) *MessageService {
	return &MessageService{
		repo: repo,
	}
}

func (s *MessageService) CreateMessage(message models.Message) error {
	return s.repo.CreateMessage(message)
}

func (s *MessageService) GetMessagesForTimeInterval(chatID int64) (string, error) {
	endTime := time.Now()
	startTime := endTime.Add(-time.Hour)
	messages, err := s.repo.GetMessagesForTimeInterval(startTime, endTime, chatID)
	if err != nil {
		return "", err
	}

	responseMsg, err := communication.GetRetelling(messages)
	if err != nil {
		return "", err
	}

	return responseMsg, err
}
