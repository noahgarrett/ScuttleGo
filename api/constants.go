package api

// BaseUrl represents the base url for each api
type BaseUrl string

const (
	LolBaseUrl BaseUrl = "api.riotgames.com"
)

// Region represents a server region
type Region string

const (
	RegionBrazil            Region = "br1"
	RegionEuropeNorthEast   Region = "eun1"
	RegionEuropeWest        Region = "euw1"
	RegionJapan             Region = "jp1"
	RegionKorea             Region = "kr"
	RegionLatinAmericaNorth Region = "la1"
	RegionLatinAmericaSouth Region = "la2"
	RegionNorthAmerica      Region = "na1"
	RegionOceania           Region = "oc1"
	RegionTurkey            Region = "tr1"
	RegionRussia            Region = "ru"
	RegionPBE               Region = "pbe1"
)

// Route represents a server region's route
type Route string

const (
	RouteAmericas Route = "americas"
	RouteAsia     Route = "asia"
	RouteEurope   Route = "europe"
)

var (
	// Regions is a list of all available regions
	Regions = []Region{
		RegionBrazil,
		RegionEuropeNorthEast,
		RegionEuropeWest,
		RegionJapan,
		RegionKorea,
		RegionLatinAmericaNorth,
		RegionLatinAmericaSouth,
		RegionNorthAmerica,
		RegionOceania,
		RegionTurkey,
		RegionRussia,
		RegionPBE,
	}

	// RegionToRoute maps each region to its route
	RegionToRoute = map[Region]Route{
		RegionBrazil:            RouteAmericas,
		RegionEuropeNorthEast:   RouteEurope,
		RegionEuropeWest:        RouteEurope,
		RegionJapan:             RouteAsia,
		RegionKorea:             RouteAsia,
		RegionLatinAmericaNorth: RouteAmericas,
		RegionLatinAmericaSouth: RouteAmericas,
		RegionNorthAmerica:      RouteAmericas,
		RegionOceania:           RouteAmericas,
		RegionTurkey:            RouteEurope,
		RegionRussia:            RouteEurope,
	}
)
