package lol

type LolClient struct {
	apiKey   string
	Summoner *SummonerClient
}

func CreateClient(apiKey string) *LolClient {
	return &LolClient{
		apiKey:   apiKey,
		Summoner: &SummonerClient{apiKey: apiKey},
	}
}
