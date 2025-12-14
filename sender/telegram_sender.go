package sender

import (
	"errors"
	"fmt"
)

type TelegramSender struct {
	apiKey string
}

func (s *TelegramSender) Send(message string) error {
	if s.apiKey == "" {
		return errors.New("missing Telegram Bot API key")
	}

	// Implementation for sending message via Telegram API
	fmt.Println("Sending message via Telegram:", message)
	return nil
}

// func telegram(share *Share) {
// 	body, err := json.Marshal(map[string]any{
// 		"chat_id": 123,
// 		"text":    fmt.Sprintf("%s: %v UAH", share.name, share.amount),
// 	})

// 	if err != nil {
// 		fmt.Println("Error marshaling JSON:", err)
// 		return
// 	}

// 	botKey := os.Getenv("BOT_KEY")

// 	if botKey == "" {
// 		fmt.Println("BOT_KEY environment variable is not set, not sending to Telegram")
// 		return
// 	} else {
// 		resp, err := http.Post(
// 			fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botKey),
// 			"application/json",
// 			bytes.NewBuffer(body),
// 		)

// 		if err != nil {
// 			fmt.Println("Error sending message to Telegram:", err)
// 			return
// 		}

// 		defer resp.Body.Close()

// 		if resp.StatusCode != http.StatusOK {
// 			fmt.Printf("Telegram API returned status code %d\n", resp.StatusCode)

// 			body, err := io.ReadAll(resp.Body)

// 			if err != nil {
// 				fmt.Println("Error reading response body:", err)
// 				return
// 			}

// 			fmt.Printf("Response: %s", body)
// 		}
// 	}
// }
