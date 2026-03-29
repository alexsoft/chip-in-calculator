package notifier

import (
	"testing"
)

func TestGetNotifierTelegram(t *testing.T) {
	t.Setenv("TELEGRAM_BOT_API_KEY", "1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	t.Setenv("TELEGRAM_CHAT_ID", "123456")

	n, err := GetNotifier()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := n.(*TelegramSender); !ok {
		t.Errorf("expected *TelegramSender, got %T", n)
	}
}

func TestGetNotifierError(t *testing.T) {
	tests := []struct {
		name    string
		envVars map[string]string
	}{
		{
			name:    "No env variables",
			envVars: map[string]string{},
		},
		{
			name:    "No telegram bot api key but with chat ID",
			envVars: map[string]string{"TELEGRAM_CHAT_ID": "123456"},
		},
		{
			name:    "No chat ID",
			envVars: map[string]string{"TELEGRAM_BOT_API_KEY": "123456"},
		},
		{
			name:    "Cannot parse chat ID",
			envVars: map[string]string{"TELEGRAM_BOT_API_KEY": "123456", "TELEGRAM_CHAT_ID": "invalid"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.envVars {
				t.Setenv(k, v)
			}

			n, err := GetNotifier()

			if err == nil {
				t.Errorf("expected error, got notifier %T", n)
			}
		})
	}
}
