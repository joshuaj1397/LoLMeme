package riotapi

import "fmt"

// SummonerDto for grabbing a Summoner from the Riot API
type SummonerDto struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	SummonerLevel int64  `json:"summonerLevel"`
	RevisionDate  int64  `json:"revisionDate"`
	ID            int64  `json:"id"`
	AccountID     int64  `json:"accountId"`
}

// GetSummoner using Riot's official API
func GetSummoner(region *string, summonerName string) (*SummonerDto, error) {
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/v3/summoners/by-name/%s", *region, summonerName)
	var s SummonerDto

	err := GetObj(url, &s)

	// Try to reach the endpoint again, but change the platform to the old NA value
	if err != nil && *region == NA1 {
		*region = NA
		return GetSummoner(region, summonerName)
	}
	return &s, err
}
