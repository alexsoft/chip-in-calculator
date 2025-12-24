package notifier

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TelegramSender struct {
	apiKey string
	chatId int64
}

func (s *TelegramSender) Send(message string) error {
	if s.apiKey == "" {
		return errors.New("missing Telegram Bot API key")
	}
	if s.chatId == 0 {
		return errors.New("missing Telegram chat ID")
	}

	// Implementation for sending message via Telegram API
	fmt.Println("Sending message via Telegram:", message)

	body, err := json.Marshal(map[string]any{
		"chat_id": s.chatId,
		"text":    message,
	})

	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	c := &http.Client{Timeout: 3 * time.Second}

	resp, err := c.Post(
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", s.apiKey),
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		fmt.Println("Error sending message to Telegram:", err)

		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Telegram API returned status code %d\n", resp.StatusCode)

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Error reading response body:", err)

			return err
		}

		fmt.Printf("Response: %s", string(body))

		return fmt.Errorf("failed to send message to Telegram")
	}

	return nil
}
