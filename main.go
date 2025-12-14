package main

import (
	"fmt"
	"os"

	"github.com/alexsoft/chip-in-calculator/calculator"
	"github.com/alexsoft/chip-in-calculator/config"
	"github.com/alexsoft/chip-in-calculator/sender"
	"github.com/umputun/go-flags"
)

const netflixMembersCount = 3
const netflixPrice = 21.99

const spotifyMembersCount = 5
const spotifyPrice = 21.99

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

	flags.Parse(&opts)

	if opts.ExchangeRate <= 0 {
		fmt.Println("Valid exchange rate must be provided")
		os.Exit(1)
	}
	fmt.Printf("Exchange rate: %v\n", opts.ExchangeRate)

	calculator := calculator.NewCalculator(cfg.Subscriptions)

	shares := calculator.Calculate(opts.ExchangeRate)

	sender := sender.GetSender()

	for _, share := range shares {
		sender.Send(fmt.Sprintf("%s: %v UAH", share.Name, share.Amount))
	}
}
