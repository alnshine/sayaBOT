package communication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alnshine/sayaBOT/internal/models"
)

func GetRetelling(messages []models.Message) (string, error) {
	chat := models.Chat{
		Messages: messages,
	}

	chatJSON, err := json.Marshal(chat)
	if err != nil {
		return "", err
	}

	url := "http://localhost:5000/process_json"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(chatJSON))
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	var response models.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	var responseMsg string
	responseMsg += fmt.Sprintf("Время начала: %s\n", response.TimeStart.Format("2006-01-02 15:04:05"))
	responseMsg += fmt.Sprintf("Время окончания: %s\n\n", response.TimeEnd.Format("2006-01-02 15:04:05"))
	responseMsg += fmt.Sprintf("Пересказ беседы:\n%s\n\n", response.Retelling)

	return responseMsg, nil
}
