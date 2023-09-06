package handlers

import (
	"bot-tele/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TelegramHandler interface {
	HandleTelegramMessage(c *gin.Context)
}

type telegramHandler struct {
	uc usecase.TelegramUseCase
}

func NewTelegramHandler(uc usecase.TelegramUseCase) TelegramHandler {
	return &telegramHandler{uc}
}

func (h *telegramHandler) HandleTelegramMessage(c *gin.Context) {
	var msg map[string]interface{}
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	chatID, txt := h.uc.ParseMessage(msg)
	switch txt {
	case "hi":
		if err := h.uc.SendMessage(chatID, "Hello, world!"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	case "image":
		if err := h.uc.SendImage(chatID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	case "audio":
		if err := h.uc.SendAudio(chatID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	case "video":
		if err := h.uc.SendVideo(chatID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	case "poll":
		if err := h.uc.SendPoll(chatID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	case "button":
		if err := h.uc.SendButton(chatID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	default:
		if err := h.uc.SendMessage(chatID, "Are you okay ?"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
