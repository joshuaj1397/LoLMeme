package riotapi

import (
	"fmt"
)

// SummonerDto for grabbing a Summoner from the Riot API
type SummonerDto struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	SummonerLevel int64  `json:"summonerLevel"`
	RevisionDate  int64  `json:"revisionDate"`
	ID            int64  `json:"id"`
	AccountID     int64  `json:"acccountId"`
}

// GetSummoner using Riot's official API
func GetSummoner(summonerName string) (*SummonerDto, error) {
	url := "https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/" + summonerName
	fmt.Println(url)
	var s SummonerDto
	err := GetObj(url, &s)
	fmt.Println(s)
	return &s, err
}
