package notifier

import (
	"fmt"

	"github.com/alexsoft/chip-in-calculator/calculator"
)

type Formatter interface {
	Format(share *calculator.Share) string
}

func GetFormatter(s *calculator.Share) Formatter {
	switch s.Name {
	case "Spotify":
		return &SpotifyFormatter{}
	case "Netflix":
		return &NetflixFormatter{}
	default:
		return &DefaultFormatter{}
	}
}

type DefaultFormatter struct{}

func (f *DefaultFormatter) Format(share *calculator.Share) string {
	return fmt.Sprintf("%s: %v UAH", share.Name, share.Amount)
}

type SpotifyFormatter struct{}

func (f *SpotifyFormatter) Format(share *calculator.Share) string {
	return fmt.Sprintf("Усім привіт! 🙃\nЗа Spotify в цьому місяці – ₴%v %s", share.Amount, "🎫")
}

type NetflixFormatter struct{}

func (f *NetflixFormatter) Format(share *calculator.Share) string {
	return fmt.Sprintf("Здорова! А за Netflix – ₴%v 🙃 %s", share.Amount, "🎫")
}
