package main

import (
	"testing"
)

func TestCalculateShare(t *testing.T) {
	tests := []struct {
		price        float64
		membersCount int
		rate         float64
		expected     int64
	}{
		{100, 10, 1.2, 12},
		{200, 5, 1.5, 60},
		{21.99, 3, 45, 330},
	}

	for _, tt := range tests {
		actual := CalculateShare(tt.price, tt.membersCount, tt.rate)
		if actual != tt.expected {
			t.Errorf("CalculateShare(%f, %d, %f) = %d; want %d", tt.price, tt.membersCount, tt.rate, actual, tt.expected)
		}
	}
}
