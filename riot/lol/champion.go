package lol

import (
	"encoding/json"
	"github.com/noahgarrett/scuttlego/api"
	"github.com/noahgarrett/scuttlego/internal"
)

type ChampionClient struct {
	apiKey string
}

// GetRotation returns champion rotations, including free-to-play and low-level free-to-play rotations
func (c *ChampionClient) GetRotation(region api.Region) *api.ChampionRotationDTO {
	res := internal.GetRequest(region, api.LolBaseUrl, "lol/platform/v3/champion-rotations", "", c.apiKey)

	cRotation := new(api.ChampionRotationDTO)
	json.NewDecoder(res.Body).Decode(cRotation)
	return cRotation
}
