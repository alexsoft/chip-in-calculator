package notifier

import (
	"testing"
)

func TestGetNotifierTelegram(t *testing.T) {
	t.Setenv("TELEGRAM_BOT_API_KEY", "1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	t.Setenv("TELEGRAM_CHAT_ID", "123456")

	notifier := GetNotifier()

	if _, ok := notifier.(*TelegramSender); !ok {
		t.Errorf("expected *TelegramSender, got %T", notifier)
	}
}

func TestGetNotifierNil(t *testing.T) {
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

			notifier := GetNotifier()

			if _, ok := notifier.(*NilSender); !ok {
				t.Errorf("expected *NilSender, got %T", notifier)
			}
		})
	}
}
