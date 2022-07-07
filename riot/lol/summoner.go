package lol

import (
	"ScuttleGo/api"
	"encoding/json"
	"net/http"
)

type SummonerClient struct {
	apiKey string
}

// GetByName Gets a summoner's information via summoner name
func (s *SummonerClient) GetByName(region api.Region, summonerName string) *api.SummonerDTO {
	client := http.Client{}

	url := "https://" + string(region) + ".api.riotgames.com/lol/summoner/v4/summoners/by-name/" + summonerName

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Riot-Token", s.apiKey)

	res, _ := client.Do(req)

	summ := new(api.SummonerDTO)
	json.NewDecoder(res.Body).Decode(summ)
	return summ
}
