package lol

import (
	"encoding/json"
	"github.com/noahgarrett/scuttlego/api"
	"github.com/noahgarrett/scuttlego/internal"
	"net/http"
)

type MasteryClient struct {
	apiKey string
}

// GetAllMasteriesBySummoner returns an array of ChampionMastery objects
func (m *MasteryClient) GetAllMasteriesBySummoner(region api.Region, summonerId string) []api.ChampionMasteryDTO {
	res := internal.GetRequest(region, api.LolBaseUrl, "lol/champion-mastery/v4/champion-masteries/by-summoner", summonerId, m.apiKey)

	masteryList := make([]api.ChampionMasteryDTO, 0)
	json.NewDecoder(res.Body).Decode(&masteryList)

	return masteryList

}

// GetChampionMasteryBySummoner returns a ChampionMastery object for specific champion requested from a summoner
func (m *MasteryClient) GetChampionMasteryBySummoner(region api.Region, summonerId string, championId int64) *api.ChampionMasteryDTO {
	client := http.Client{}
	url := "https://" + string(region) + "." + string(api.LolBaseUrl) +
		"/lol/champion-mastery/v4/champion-masteries/by-summoner/" + summonerId + "/by-champion/" + string(championId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Riot-Token", m.apiKey)
	res, _ := client.Do(req)

	mastery := new(api.ChampionMasteryDTO)
	json.NewDecoder(res.Body).Decode(mastery)
	return mastery
}
