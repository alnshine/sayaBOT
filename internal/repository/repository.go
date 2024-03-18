package repository

import (
	"time"

	"github.com/alnshine/sayaBOT/internal/models"
	"github.com/jmoiron/sqlx"
)

type Message interface {
	CreateMessage(message models.Message) error
	GetMessagesForTimeInterval(startTime, endTime time.Time, chatID int64) ([]models.Message, error)
}

type Repository struct {
	Message
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Message: NewMessagePostgres(db),
	}
}
