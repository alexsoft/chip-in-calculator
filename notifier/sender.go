package notifier

import (
	"fmt"
	"os"
	"strconv"
)

type Notifier interface {
	Send(message string) error
}

func GetNotifier() (Notifier, error) {
	apiKey := os.Getenv("TELEGRAM_BOT_API_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("TELEGRAM_BOT_API_KEY is not set")
	}

	chatIdString := os.Getenv("TELEGRAM_CHAT_ID")

	if chatIdString == "" {
		return nil, fmt.Errorf("TELEGRAM_CHAT_ID is not set")
	}

	chatId, err := strconv.ParseInt(chatIdString, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("invalid TELEGRAM_CHAT_ID: %w", err)
	}

	return &TelegramSender{
		apiKey: apiKey,
		chatId: chatId,
	}, nil
}
