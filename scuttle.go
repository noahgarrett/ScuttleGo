package main

import (
	"fmt"
	"github.com/noahgarrett/scuttlego/api"
	"github.com/noahgarrett/scuttlego/riot/lol"
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
	scuttle := CreateClient("")
	fmt.Println(scuttle.LoL.Summoner.GetByAccountId(api.RegionNorthAmerica, "Hyy3p4ZLAzRI4Nm2etYxRExzrYEM1B-yzs99ciEiZPoz6GfFVUJsSFSm").Id)
	fmt.Println(len(scuttle.LoL.Mastery.GetAllMasteriesBySummoner(api.RegionNorthAmerica, "a5TZtZswIYxOelDN8yRvenmTmeDVNwBRUC01CZnnCP0uE3PrHlbjHbFAkQ")))

	fmt.Println(scuttle.LoL.Mastery.GetChampionMasteryBySummoner(api.RegionNorthAmerica, "a5TZtZswIYxOelDN8yRvenmTmeDVNwBRUC01CZnnCP0uE3PrHlbjHbFAkQ", 150).ChampionPoints)
}
