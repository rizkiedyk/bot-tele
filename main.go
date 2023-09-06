package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var telegramToken string

func telParseMessage(message map[string]interface{}) (int64, string) {
	chat := message["message"].(map[string]interface{})["chat"].(map[string]interface{})
	chatID := chat["id"].(int64)
	text := message["message"].(map[string]interface{})["text"].(string)
	return chatID, text
}

func telSendMessage(chatID int64, text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramToken)
	payload := gin.H{
		"chat_id": chatID,
		"text":    text,
	}

	// koncersi pesan ke json
	messageJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Kirim permintaan HTTP POST
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(messageJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to send. Status code: %d", resp.StatusCode)
	}

	return nil
}

func telSendImage(chatID int64) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", telegramToken)
	data := map[string]interface{}{
		"chat_id": chatID,
		"photo":   "https://i.pinimg.com/originals/c2/cf/70/c2cf70612bcdcafbd892b8aaca092221.jpg",
		"caption": "This is a sample image",
	}

	// Buat permintaan multipart/fom-data
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	for key, val := range data {
		_ = writer.WriteField(key, fmt.Sprintf("%v", val))
	}

	// Kirim permintaan POST dengan gambar
	writer.Close()
	resp, err := http.Post(url, writer.FormDataContentType(), &requestBody)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to send image. Status code: %d", resp.StatusCode)
	}

	return nil
}

func telSendAudio(chatID int64) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendAudio", telegramToken)

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
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to send audio. Status code: %d", resp.StatusCode)
	}

	return nil
}

func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	telegramToken = os.Getenv("TOKEN_API")

	r.POST("/", func(c *gin.Context) {
		var msg map[string]interface{}
		if err := c.ShouldBindJSON(&msg); err != nil {
			fmt.Println("Error parsing JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		chatID, txt := telParseMessage(msg)
		switch txt {
		case "hi":
			err := telSendMessage(chatID, "Hello, world!")
			if err != nil {
				fmt.Println("Error sending message:", err)
			}
		case "image":
			err := telSendImage(chatID)
			if err != nil {
				fmt.Println("Error sending image:", err)
			}
		case "audio":
			err := telSendAudio(chatID)
			if err != nil {
				fmt.Println("Error sending audio : ", err)
			}
		default:
			err := telSendMessage(chatID, "Are you okay ?")
			if err != nil {
				fmt.Println("Error sending message:", err)
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome!")
	})

	// Ganti port sesuai keinginan Anda, misalnya ":8080"
	r.Run(":8080")
}
