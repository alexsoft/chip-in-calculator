package config

import (
	"encoding/json"
	"os"
)

type Subscription struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	MembersCount int     `json:"membersCount"`
}

type Config struct { 
	Subscriptions []Subscription `json:"subscriptions"`
}

func Load(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
