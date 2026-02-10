package notifier

import (
	"slices"
	"strings"
	"testing"

	"github.com/alexsoft/chip-in-calculator/calculator"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		share          calculator.Share
		mentions       string
		expectedPrefix string
	}{
		{calculator.Share{Name: "Spotify", Amount: 10}, "@user1 @user2", "Усім привіт! @user1 @user2\nЗа Spotify в цьому місяці – ₴10 "},
		{calculator.Share{Name: "Spotify", Amount: 10}, "", "Усім привіт! \nЗа Spotify в цьому місяці – ₴10 "},
		{calculator.Share{Name: "Netflix", Amount: 150}, "", "Здорова! А за Netflix – ₴150 "},
	}

	for _, tt := range tests {
		actual := Format(&tt.share, tt.mentions)

		if !strings.HasPrefix(actual, tt.expectedPrefix) {
			t.Errorf("Format(%q) = %q, want prefix %q", tt.share.Name, actual, tt.expectedPrefix)
		}

		emoji := strings.TrimPrefix(actual, tt.expectedPrefix)
		if !slices.Contains(emojis, emoji) {
			t.Errorf("Format(%q) emoji = %q, want one of the predefined emojis", tt.share.Name, emoji)
		}
	}
}

func TestFormatDefault(t *testing.T) {
	share := calculator.Share{Name: "Default", Amount: 111}
	actual := Format(&share, "")
	expected := "Default: ₴111"

	if actual != expected {
		t.Errorf("Format(%q) = %q, want %q", share.Name, actual, expected)
	}
}

func TestRandomEmoji(t *testing.T) {
	emoji := randomEmoji()
	if !slices.Contains(emojis, emoji) {
		t.Errorf("randomEmoji() = %q, want one of the predefined emojis", emoji)
	}
}
