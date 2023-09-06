package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
)

type TelegramUseCase interface {
	ParseMessage(message map[string]interface{}) (int64, string)
	SendMessage(chatID int64, text string) error
	SendImage(chatID int64) error
	SendAudio(chatID int64) error
	SendVideo(chatID int64) error
	SendPoll(chatID int64) error
	SendButton(chatID int64) error
}

type telegramUseCase struct {
	telegramToken string
}

func NewTelegramUseCase(telegramToken string) TelegramUseCase {
	return &telegramUseCase{telegramToken: telegramToken}
}

func (uc *telegramUseCase) ParseMessage(message map[string]interface{}) (int64, string) {
	chat := message["message"].(map[string]interface{})["chat"].(map[string]interface{})
	chatID := chat["id"].(int64)
	text := message["message"].(map[string]interface{})["text"].(string)
	return chatID, text
}

func (uc *telegramUseCase) SendMessage(chatID int64, text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", uc.telegramToken)
	payload := map[string]interface{}{
		"chat_id": chatID,
		"text":    text,
	}

	messageJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(messageJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send. Status code: %d", resp.StatusCode)
	}

	return nil
}

func (uc *telegramUseCase) SendImage(chatID int64) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", uc.telegramToken)
	data := map[string]interface{}{
		"chat_id": chatID,
		"photo":   "https://i.pinimg.com/originals/c2/cf/70/c2cf70612bcdcafbd892b8aaca092221.jpg",
		"caption": "This is a sample image",
	}

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	for key, val := range data {
		_ = writer.WriteField(key, fmt.Sprintf("%v", val))
	}

	writer.Close()
	resp, err := http.Post(url, writer.FormDataContentType(), &requestBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send image. Status code: %d", resp.StatusCode)
	}

	return nil
}

func (uc *telegramUseCase) SendAudio(chatID int64) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendAudio", uc.telegramToken)
	data := map[string]interface{}{
		"chat_id": chatID,
		"audio":   "http://www.largesound.com/ashborytour/sound/brobob.mp3",
	}

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	for key, val := range data {
		_ = writer.WriteField(key, fmt.Sprintf("%v", val))
	}

	writer.Close()
	resp, err := http.Post(url, writer.FormDataContentType(), &requestBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send audio. Status code: %d", resp.StatusCode)
	}

	return nil
}

func (uc *telegramUseCase) SendVideo(chatID int64) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendVideo", uc.telegramToken)
	data := map[string]interface{}{
		"chat_id": chatID,
		"video":   "https://www.appsloveworld.com/wp-content/uploads/2018/10/640.mp4",
	}

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	for key, val := range data {
		_ = writer.WriteField(key, fmt.Sprintf("%v", val))
	}

	writer.Close()
	resp, err := http.Post(url, writer.FormDataContentType(), &requestBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send video. Status code: %d", resp.StatusCode)
	}

	return nil
}

func (uc *telegramUseCase) SendPoll(chatID int64) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPoll", uc.telegramToken)
	options := []string{"North", "South", "East", "West"}
	data := map[string]interface{}{
		"chat_id":           chatID,
		"question":          "In which direction does the sun rise ?",
		"options":           options,
		"is_anonymous":      false,
		"type":              "quiz",
		"correct_option_id": 2,
	}

	messageJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(messageJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send poll. Status code: %d", resp.StatusCode)
	}

	return nil
}

func (uc *telegramUseCase) SendButton(chatID int64) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", uc.telegramToken)

	keyboard := map[string]interface{}{
		"keyboard": [][]map[string]interface{}{
			{
				{"text": "Rizki"},
				{"text": "Juju"},
			},
		},
	}

	data := map[string]interface{}{
		"chat_id":      chatID,
		"text":         "What is this ?",
		"reply_markup": keyboard,
	}

	messageJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(messageJSON))
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send button. Status code: %d", resp.StatusCode)
	}

	return nil
}
