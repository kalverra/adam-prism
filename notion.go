package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// WriteFitbitData writes Fitbit data to Notion
func WriteFitbitData(client *resty.Client) {
	client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", "secret"))
}

// Add to Pomofocus Database

// Add to TODOist Database
