package notifier

import (
	"fmt"
	"os"
	"strconv"
)

type Notifier interface {
	Send(message string) error
}

func GetNotifier() Notifier {
	apiKey := os.Getenv("TELEGRAM_BOT_API_KEY")

	if apiKey == "" {
		return &NilSender{}
	}

	chatIdString := os.Getenv("TELEGRAM_CHAT_ID")

	if chatIdString == "" {
		return &NilSender{}
	}

	chatId, err := strconv.ParseInt(chatIdString, 10, 64)

	if err != nil {
		fmt.Println("Error parsing chat ID:", err)

		return &NilSender{}
	}

	return &TelegramSender{
		apiKey: apiKey,
		chatId: chatId,
	}
}
