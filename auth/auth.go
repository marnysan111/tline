package auth

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func AuthLine() (*linebot.Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	Secret := os.Getenv("Secret")
	Token := os.Getenv("Token")
	bot, err := linebot.New(Secret, Token)
	if err != nil {
		return nil, err
	}
	return bot, nil
}
