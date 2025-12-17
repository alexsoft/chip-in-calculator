package calculator

import (
	"testing"

	"github.com/alexsoft/chip-in-calculator/config"
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
		actual := calculateShare(tt.price, tt.membersCount, tt.rate)
		if actual != tt.expected {
			t.Errorf("CalculateShare(%f, %d, %f) = %d; want %d", tt.price, tt.membersCount, tt.rate, actual, tt.expected)
		}
	}
}

func TestCalculate(t *testing.T) {
	sut := NewCalculator(
		[]config.Subscription{
			config.Subscription{
				Name:         "Subscription 1",
				Price:        100,
				MembersCount: 2,
			},
			config.Subscription{
				Name:         "Subscription 2",
				Price:        22.99,
				MembersCount: 3,
			},
		},
	)

	actual := sut.Calculate(11)

	if len(actual) != 2 {
		t.Errorf("Calculate(11) = %d; want %d", len(actual), 2)
	}

	if actual[0].Name != "Subscription 1" {
		t.Errorf("Calculate(11) = %s; want %s", actual[0].Name, "Subscription 1")
	}
	if actual[0].Amount != 550 {
		t.Errorf("Calculate(11) = %d; want %s", actual[0].Amount, "50 (amount)")
	}
}
