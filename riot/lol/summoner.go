package lol

import (
	"encoding/json"
	"github.com/noahgarrett/scuttlego/api"
	"github.com/noahgarrett/scuttlego/internal"
)

type SummonerClient struct {
	apiKey string
}

// GetByAccountId Gets a summoner's information via account id
func (s *SummonerClient) GetByAccountId(region api.Region, accountId string) *api.SummonerDTO {
	res := internal.GetRequest(region, api.LolBaseUrl, "lol/summoner/v4/summoners/by-account", accountId, s.apiKey)

	summ := new(api.SummonerDTO)
	json.NewDecoder(res.Body).Decode(summ)
	return summ
}

// GetByName Gets a summoner's information via summoner name
func (s *SummonerClient) GetByName(region api.Region, summonerName string) *api.SummonerDTO {
	res := internal.GetRequest(region, api.LolBaseUrl, "lol/summoner/v4/summoners/by-name", summonerName, s.apiKey)

	summ := new(api.SummonerDTO)
	json.NewDecoder(res.Body).Decode(summ)
	return summ
}

// GetByPuuid Gets a summoner's information via puuid
func (s *SummonerClient) GetByPuuid(region api.Region, puuid string) *api.SummonerDTO {
	res := internal.GetRequest(region, api.LolBaseUrl, "lol/summoner/v4/summoners/by-puuid", puuid, s.apiKey)

	summ := new(api.SummonerDTO)
	json.NewDecoder(res.Body).Decode(summ)
	return summ
}

// GetBySummonerId Gets a summoner's information via puuid
func (s *SummonerClient) GetBySummonerId(region api.Region, summonerId string) *api.SummonerDTO {
	res := internal.GetRequest(region, api.LolBaseUrl, "lol/summoner/v4/summoners", summonerId, s.apiKey)

	summ := new(api.SummonerDTO)
	json.NewDecoder(res.Body).Decode(summ)
	return summ
}
