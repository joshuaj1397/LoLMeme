package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	RIOT_API_KEY string
	client       = http.Client{
		Timeout: time.Second * 5,
	}
)

type Summoner struct {
	ProfileIconID int    `json:"id"`
	Name          string `json:"name"`
	SummonerLevel int32  `json:"summonerLevel"`
	RevisionDate  int32  `json:"revisionDate"`
	ID            int32  `json:"id"`
	AccountID     int32  `json:"acccountId"`
}

func init() {
	RIOT_API_KEY = os.Getenv("RIOT_API_KEY")
	if RIOT_API_KEY == "" {
		panic("RIOT_API_KEY NOT FOUND IN ENV")
	}
}

// Constructs a Summoner using Riot's official API
func GetSummoner(summonerName string) (*Summoner, error) {

	url := "https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/" + summonerName
	var s Summoner

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Riot-Token", RIOT_API_KEY)

	res, getErr := client.Do(req)
	if getErr != nil {
		return nil, err
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, err
	}

	jsonErr := json.Unmarshal(body, &s)
	if jsonErr != nil {
		return nil, err
	}

	return &s, nil
}
