package repository

import (
	"github.com/alnshine/sayaBOT/internal/models"
	"github.com/jmoiron/sqlx"
)

type Message interface {
	CreateMessage(message models.Message) error
}

type Repository struct {
	Message
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Message: NewMessagePostgres(db),
	}
}
