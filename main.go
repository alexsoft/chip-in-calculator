package main

import (
	"fmt"
	"math"
	"os"

	"github.com/umputun/go-flags"
)

const netflixMembersCount = 3
const netflixPrice = 21.99

const spotifyMembersCount = 5
const spotifyPrice = 21.99

var opts struct {
	ExchangeRate float64 `short:"r" long:"rate" description:"current EUR -> UAH rate"`
}

type Share struct {
	name   string
	amount int64
}

func NewShare(name string, amount int64) *Share {
	return &Share{
		name:   name,
		amount: amount,
	}
}

func main() {
	fmt.Printf("chip-in-go\n")

	flags.Parse(&opts)

	if opts.ExchangeRate <= 0 {
		fmt.Println("Valid exchange rate must be provided")
		os.Exit(1)
	}

	fmt.Printf("Exchange rate: %v\n", opts.ExchangeRate)

	for _, share := range calculateWithExchangeRate(opts.ExchangeRate) {
		fmt.Printf("%s: %v UAH\n", share.name, share.amount)
	}
}

func calculateWithExchangeRate(rate float64) []*Share {
	return []*Share{
		NewShare("Netflix", int64(math.Round(netflixPrice/netflixMembersCount*rate))),
		NewShare("Spotify", int64(math.Round(spotifyPrice/spotifyMembersCount*rate))),
	}
}
