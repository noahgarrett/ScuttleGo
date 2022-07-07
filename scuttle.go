package main

import (
	"ScuttleGo/riot/lol"
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

//func main() {
//	scuttle := CreateClient("<API Key Here>")
//	fmt.Println(scuttle.LoL.Summoner.GetByAccountId(api.RegionNorthAmerica, "Hyy3p4ZLAzRI4Nm2etYxRExzrYEM1B-yzs99ciEiZPoz6GfFVUJsSFSm").Name)
//}
