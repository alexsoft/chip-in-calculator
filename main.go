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

func main() {
	fmt.Printf("chip-in-go\n")

	flags.Parse(&opts)

	if opts.ExchangeRate <= 0 {
		fmt.Println("Valid exchange rate must be provided")
		os.Exit(1)
	}

	fmt.Printf("Exchange rate: %v\n", opts.ExchangeRate)

	calculateWithExchangeRate(opts.ExchangeRate)
}

func calculateWithExchangeRate(rate float64) {
	fmt.Println("Netflix:", math.Round(netflixPrice/netflixMembersCount*rate), "UAH")
	fmt.Println("Spotify:", math.Round(spotifyPrice/spotifyMembersCount*rate), "UAH")
}
