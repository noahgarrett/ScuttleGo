package internal

import (
	"fmt"
	"github.com/noahgarrett/scuttlego/api"
	"net/http"
)

func GetRequest(region api.Region, baseUrl api.BaseUrl, subPath string, data string, apiKey string) *http.Response {
	client := http.Client{}

	url := "https://" + string(region) + "." + string(baseUrl) + "/" + subPath + "/" + data
	fmt.Println(url)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Riot-Token", apiKey)

	res, _ := client.Do(req)

	return res
}
