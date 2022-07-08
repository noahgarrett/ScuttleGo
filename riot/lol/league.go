package lol

import (
	"encoding/json"
	"github.com/noahgarrett/scuttlego/api"
	"net/http"
)

type LeagueClient struct {
	apiKey string
}

func (l *LeagueClient) GetEntriesByQueue(region api.Region, queueType string, tier string, division string, page string) []api.LeagueEntryDTO {
	client := http.Client{}
	url := "https://" + string(region) + "." + string(api.LolBaseUrl) + "/lol/league/v4/entries/" + queueType + "/" + tier + "/" + division + "?page=" + page

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Riot-Token", l.apiKey)
	res, _ := client.Do(req)

	entryList := make([]api.LeagueEntryDTO, 0)
	json.NewDecoder(res.Body).Decode(&entryList)
	return entryList
}
