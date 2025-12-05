package main

import (
	"math"
)

func CalculateShare(
	price float64,
	membersCount int,
	rate float64,
) int64 {
	return int64(math.Round(price / float64(membersCount) * rate))
}
