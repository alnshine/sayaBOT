package repository

import (
	"fmt"
	"time"

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

func (r *MessagePostgres) GetMessagesForTimeInterval(startTime, endTime time.Time, chatID int64) ([]models.Message, error) {
	var lists []models.Message
	strQuery := `
		SELECT id, content, username, time, chat_id
		FROM %s
		WHERE time BETWEEN $1 AND $2 AND chat_id = $3
	`

	query := fmt.Sprintf(strQuery, messageTable)
	err := r.db.Select(&lists, query, startTime, endTime, chatID)
	return lists, err
}
