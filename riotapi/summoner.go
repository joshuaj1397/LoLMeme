package riotapi

import (
	"net/http"
	"os"
	"time"
)

var (
	// RIOT_API_KEY env var denoting my secret API key
	RIOT_API_KEY string

	// Client would instantiate a new http.Client with a 5 sec timeout on requests
	Client = http.Client{
		Timeout: time.Second * 5,
	}
)

// SummonerDto for grabbing a Summoner from the Riot API
type SummonerDto struct {
	ProfileIconID int    `json:"id"`
	Name          string `json:"name"`
	SummonerLevel int32  `json:"summonerLevel"`
	RevisionDate  int32  `json:"revisionDate"`
	ID            int32  `json:"id"`
	AccountID     int32  `json:"acccountId"`
}

// Get the RIOT_API_KEY from env vars
func init() {
	RIOT_API_KEY = os.Getenv("RIOT_API_KEY")
	if RIOT_API_KEY == "" {
		panic("RIOT_API_KEY NOT FOUND IN ENV")
	}
}

// GetSummoner using Riot's official API
func GetSummoner(summonerName string) (*SummonerDto, error) {
	url := "https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/" + summonerName
	var s SummonerDto
	err := getObj(url, &s)
	return &s, err
}
