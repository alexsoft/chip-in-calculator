package notifier

import (
	"testing"

	"github.com/alexsoft/chip-in-calculator/calculator"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		share    calculator.Share
		expected string
	}{
		{calculator.Share{Name: "Spotify", Amount: 10}, "Усім привіт! 🙃\nЗа Spotify в цьому місяці – ₴10 🎫"},
		{calculator.Share{Name: "Netflix", Amount: 150}, "Здорова! А за Netflix – ₴150 🙃 🎫"},
		{calculator.Share{Name: "Default", Amount: 111}, "Default: ₴111"},
	}

	for _, tt := range tests {
		actual := Format(&tt.share)

		if actual != tt.expected {
			t.Errorf("Format(%q) = %q, want %q", tt.share.Name, actual, tt.expected)
		}
	}
}
