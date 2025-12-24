package notifier

import (
	"fmt"

	"github.com/alexsoft/chip-in-calculator/calculator"
)

var messageTemplates = map[string]string{
	"Spotify": "Усім привіт! 🙃\nЗа Spotify в цьому місяці – ₴%v %s",
	"Netflix": "Здорова! А за Netflix – ₴%v 🙃 %s",
}

func Format(share *calculator.Share) string {
	if template, ok := messageTemplates[share.Name]; ok {
		return fmt.Sprintf(template, share.Amount, "🎫")
	}

	return fmt.Sprintf("%s: ₴%v", share.Name, share.Amount)
}
