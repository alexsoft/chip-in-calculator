package sender

import (
	"os"
)

type Sender interface {
	Send(message string) error
}

func GetSender() Sender {
	if apiKey := os.Getenv("TELEGRAM_BOT_API_KEY"); apiKey != "" {
		return &TelegramSender{
			apiKey: apiKey,
		}
	}

	return &NilSender{}
}
