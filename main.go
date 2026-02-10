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
	ConfigPath   string  `short:"c" long:"config" default:"config.json" description:"path to config file"`
	Mentions     string  `long:"mentions" description:"Mentions to put into message after greeting"`
}

func main() {
	fmt.Printf("chip-in-go\n")

	if _, err := flags.Parse(&opts); err != nil {
		fmt.Println("Error parsing flags:", err)
		os.Exit(1)
	}

	cfg, err := config.Load(opts.ConfigPath)
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	if opts.ExchangeRate <= 0 {
		fmt.Println("Valid exchange rate must be provided")
		os.Exit(1)
	}
	fmt.Printf("Exchange rate: %v\n", opts.ExchangeRate)

	calc := calculator.NewCalculator(cfg.Subscriptions)

	notify := notifier.GetNotifier()

	for _, share := range calc.Calculate(opts.ExchangeRate) {
		if err := notify.Send(notifier.Format(share, opts.Mentions)); err != nil {
			fmt.Printf("Failed to send notification: %v, %v, %T\n", err, share, notify)
		}
	}
}
