package repository

import (
	"fmt"

	"github.com/alnshine/sayaBOT/internal/models"
	"github.com/jmoiron/sqlx"
)

type MessagePostgres struct {
	db *sqlx.DB
}

func NewMessagePostgres(db *sqlx.DB) *MessagePostgres {
	return &MessagePostgres{
		db: db,
	}
}

func (r *MessagePostgres) CreateMessage(message models.Message) error {
	query := fmt.Sprintf("INSERT INTO %s (content, username, time, chat_id) values ($1, $2, $3, $4) RETURNING id", messageTable)
	row := r.db.QueryRow(query, message.Content, message.Username, message.Time, message.ChatId)
	var id int // Assuming id is of type SERIAL
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}
