package api

type SummonerDTO struct {
	AccountId     string
	ProfileIconId string
	RevisionDate  int64
	Name          string
	Id            string
	Puuid         string
	SummonerLevel int64
}

type ChampionRotationDTO struct {
	MaxNewPlayerLevel            int
	FreeChampionIdsForNewPlayers []int
	FreeChampionIds              []int
}

type ChampionMasteryDTO struct {
	ChampionPointsUntilNextLevel int64
	ChestGranted                 bool
	ChampionId                   int64
	ChampionLevel                int
	SummonerId                   string
	ChampionPoints               int
	ChampionPointsSinceLastLevel int64
	TokensEarned                 int
}

type LeagueEntryDTO struct {
	LeagueId     string
	SummonerId   string
	SummonerName string
	QueueType    string
	Tier         string
	Rank         string
	LeaguePoints int
	Wins         int
	Losses       int
	HotStreak    bool
	Veteran      bool
	FreshBlood   bool
	Inactive     bool
}
