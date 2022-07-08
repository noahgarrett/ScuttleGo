package main

import (
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
	//scuttle := CreateClient("RGAPI-918e1b3a-c7f7-4829-9b6e-6731523fc8ac")
	//fmt.Println(scuttle.LoL.Summoner.GetByAccountId(api.RegionNorthAmerica, "Hyy3p4ZLAzRI4Nm2etYxRExzrYEM1B-yzs99ciEiZPoz6GfFVUJsSFSm").Id)
	//fmt.Println(len(scuttle.LoL.Mastery.GetAllMasteriesBySummoner(api.RegionNorthAmerica, "a5TZtZswIYxOelDN8yRvenmTmeDVNwBRUC01CZnnCP0uE3PrHlbjHbFAkQ")))
	//
	//fmt.Println(scuttle.LoL.Mastery.GetChampionMasteryBySummoner(api.RegionNorthAmerica, "a5TZtZswIYxOelDN8yRvenmTmeDVNwBRUC01CZnnCP0uE3PrHlbjHbFAkQ", 150).ChampionPoints)
	//entries := scuttle.LoL.League.GetEntriesByQueue(api.RegionNorthAmerica, "RANKED_SOLO_5x5", "BRONZE", "I", "1")
	//fmt.Println(entries)

}
