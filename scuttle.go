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
	scuttle := CreateClient("RGAPI-918e1b3a-c7f7-4829-9b6e-6731523fc8ac")
	fmt.Println(scuttle.LoL.Summoner.GetByName(api.RegionNorthAmerica, "C9Beemo").SummonerLevel)
}
