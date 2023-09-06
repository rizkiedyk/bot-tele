package main

import (
	"bot-tele/config"
	"bot-tele/handlers"
	"bot-tele/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.NewConfig()

	r := gin.Default()

	// Initial use case and handler
	telegramUseCase := usecase.NewTelegramUseCase(cfg.TelegramToken)
	telegramHandler := handlers.NewTelegramHandler(telegramUseCase)

	r.POST("/", telegramHandler.HandleTelegramMessage)
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome !")
	})

	r.Run(":" + cfg.Port)
}
