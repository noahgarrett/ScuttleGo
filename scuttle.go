package main

import (
	"ScuttleGo/api"
	"ScuttleGo/riot/lol"
	"fmt"
)

type Client struct {
	apiKey string
	LoL    *lol.LolClient
}

// CreateClient returns a created client for library usage
func CreateClient(apiKey string) *Client {
	client := &Client{
		apiKey: apiKey,
	}

	client.LoL = lol.CreateClient(apiKey)

	return client
}

func main() {
	scuttle := CreateClient("<Api Key Here>")
	fmt.Println(scuttle.LoL.Summoner.GetByName(api.RegionNorthAmerica, "C9Beemo").SummonerLevel)
}
