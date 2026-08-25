package exchangerate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	monobankBaseURL = "https://api.monobank.ua/bank/currency"
	eurCurrencyCode = 978
	uahCurrencyCode = 980
)

type MonobankClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewMonobankClient() *MonobankClient {
	return &MonobankClient{
		baseURL:    monobankBaseURL,
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}
}

type currencyRate struct {
	CurrencyCodeA int     `json:"currencyCodeA"`
	CurrencyCodeB int     `json:"currencyCodeB"`
	RateSell      float64 `json:"rateSell"`
}

// EURUAHRate returns the rate at which Monobank sells EUR for UAH,
// i.e. how much UAH it costs to buy back the EUR that was spent.
func (c *MonobankClient) EURUAHRate() (float64, error) {
	resp, err := c.httpClient.Get(c.baseURL)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch rates from Monobank: %w", err)
	}
	defer resp.Body.Close() // nolint

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("monobank API returned status code %d", resp.StatusCode)
	}

	var rates []currencyRate
	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		return 0, fmt.Errorf("failed to parse Monobank response: %w", err)
	}

	for _, r := range rates {
		if r.CurrencyCodeA == eurCurrencyCode && r.CurrencyCodeB == uahCurrencyCode {
			if r.RateSell <= 0 {
				return 0, fmt.Errorf("monobank did not return a EUR/UAH sell rate")
			}
			return r.RateSell, nil
		}
	}

	return 0, fmt.Errorf("EUR/UAH rate not found in Monobank response")
}
