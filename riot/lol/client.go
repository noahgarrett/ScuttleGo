package lol

type LolClient struct {
	apiKey   string
	Summoner *SummonerClient
	Champion *ChampionClient
	Mastery  *MasteryClient
	League   *LeagueClient
}

func CreateClient(apiKey string) *LolClient {
	return &LolClient{
		apiKey:   apiKey,
		Summoner: &SummonerClient{apiKey: apiKey},
		Champion: &ChampionClient{apiKey: apiKey},
		Mastery:  &MasteryClient{apiKey: apiKey},
		League:   &LeagueClient{apiKey: apiKey},
	}
}
