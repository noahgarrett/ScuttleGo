# ScuttleGO

ScuttleGO is a wrapper for the Riot Games API. This is written purely in Go and provides easy access to all 
supported endpoints. This library is in its BETA stage, and numerous additions are to come.

## Currently Supported
- Summoner-V4
- Champion-V3
- Champion-Mastery-V4

## Example Usage
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

	fmt.Println(summoner.Puuid)
}
```