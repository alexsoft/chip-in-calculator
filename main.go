package main

import (
	"fmt"
	"os"

	"github.com/alexsoft/chip-in-calculator/calculator"
	"github.com/alexsoft/chip-in-calculator/config"
	"github.com/alexsoft/chip-in-calculator/notifier"
	"github.com/umputun/go-flags"
)

var opts struct {
	ExchangeRate float64 `short:"r" long:"rate" description:"current EUR -> UAH rate"`
}

func main() {
	fmt.Printf("chip-in-go\n")

	cfg, err := config.Load("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	flags.Parse(&opts) // nolint

	if opts.ExchangeRate <= 0 {
		fmt.Println("Valid exchange rate must be provided")
		os.Exit(1)
	}
	fmt.Printf("Exchange rate: %v\n", opts.ExchangeRate)

	calc := calculator.NewCalculator(cfg.Subscriptions)

	notify := notifier.GetNotifier()

	for _, share := range calc.Calculate(opts.ExchangeRate) {
		if err := notify.Send(notifier.Format(share)); err != nil {
			fmt.Printf("Failed to send notification: %v, %v, %T\n", err, share, notify)
		}
	}
}
