package calculator

import (
	"math"

	"github.com/alexsoft/chip-in-calculator/config"
)

type Share struct {
	Name   string
	Amount int64
}

func NewShare(name string, amount int64) *Share {
	return &Share{
		Name:   name,
		Amount: amount,
	}
}

type Calculator struct {
	subscriptions []config.Subscription
}

func (c *Calculator) Calculate(rate float64) []*Share {
	var shares []*Share

	for _, sub := range c.subscriptions {
		price := sub.Price
		membersCount := sub.MembersCount
		name := sub.Name

		share := calculateShare(price, membersCount, rate)

		shares = append(shares, NewShare(name, share))
	}

	return shares
}

func NewCalculator(subscriptions []config.Subscription) *Calculator {
	return &Calculator{
		subscriptions: subscriptions,
	}
}

func calculateShare(
	price float64,
	membersCount int,
	rate float64,
) int64 {
	return int64(math.Round(price / float64(membersCount) * rate))
}
