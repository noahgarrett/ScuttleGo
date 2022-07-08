# ScuttleGO

ScuttleGO is a wrapper for the Riot Games API. This is written purely in Go and provides easy access to all 
supported endpoints. This library is in its BETA stage, and numerous additions are to come.

## Currently Supported
- Summoner-V4
- Champion-V3
- Champion-Mastery-V4
- League-V4

## Example Usage
`go get github.com/noahgarrett/scuttlego`

```go
package ScuttleExample

import (
	"fmt"
	"github.com/noahgarrett/scuttlego"
	"github.com/noahgarrett/scuttlego/api"
)

func main() {
	scuttle := scuttlego.CreateClient("API_KEY")

	summoner := scuttle.LoL.Summoner.GetByName(api.RegionNorthAmerica, "C9Beemo")
	rotation := scuttle.LoL.Champion.GetRotation(api.RegionNorthAmerica)
	masteries := scuttle.LoL.Mastery.GetAllMasteriesBySummoner(api.RegionNorthAmerica, "a5TZtZswIYxOelDN8yRvenmTmeDVNwBRUC01CZnnCP0uE3PrHlbjHbFAkQ")
	entries := scuttle.LoL.League.GetEntriesByQueue(api.RegionNorthAmerica, "RANKED_SOLO_5x5", "BRONZE", "I", "1")
	
	fmt.Println(summoner.Puuid)
}
```
