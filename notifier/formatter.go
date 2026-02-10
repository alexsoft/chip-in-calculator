package notifier

import (
	"fmt"
	"math/rand"

	"github.com/alexsoft/chip-in-calculator/calculator"
)

var messageTemplates = map[string]string{
	"Spotify": "Усім привіт! %[3]s\nЗа Spotify в цьому місяці – ₴%[1]v %[2]s",
	"Netflix": "Здорова! А за Netflix – ₴%[1]v %[2]s",
}

var emojis = []string{
	"🎸", "🎺", "🎷", "🎹", "🎻", "🥁", "🎵", "🎶", "🎤", "🎧",
	"🎮", "🕹️", "🎲", "🎯", "🎳", "🏓", "🏸", "🎾", "⚽", "🏀",
	"🚀", "🛸", "✈️", "🚁", "🏎️", "🚂", "⛵", "🚤", "🛶", "🎢",
	"🌈", "🌟", "⭐", "🌙", "☀️", "🔥", "❄️", "🌊", "🌸", "🌻",
	"🍕", "🍔", "🍟", "🌮", "🍣", "🍩", "🍪", "🎂", "🍰", "🧁",
	"🐶", "🐱", "🐼", "🦊", "🦁", "🐸", "🦋", "🐝", "🦄", "🐧",
	"💎", "👑", "🏆", "🎖️", "🥇", "🎁", "🎀", "🎈", "🎉", "🎊",
	"🤖", "👾", "👻", "💀", "🎃", "🦸", "🧙", "🧛", "🧟", "🦹",
	"📚", "🔮", "🧲", "💡", "🔑", "🗝️", "🧩", "🎭", "🎨", "🖌️",
	"🌍", "🗻", "🏝️", "🌋", "🏰", "🗼", "🎪", "🎠", "🎡", "⛺",
}

func randomEmoji() string {
	return emojis[rand.Intn(len(emojis))]
}

func Format(share *calculator.Share, mentions string) string {
	if template, ok := messageTemplates[share.Name]; ok {
		return fmt.Sprintf(template, share.Amount, randomEmoji(), mentions)
	}

	return fmt.Sprintf("%s: ₴%v", share.Name, share.Amount)
}
