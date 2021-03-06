package riotapi

import (
	"fmt"
)

// MatchListDto for grabbing a MatchList from the Riot API
type MatchListDto struct {
	Matches []struct {
		PlatformID string `json:"platformId"`
		GameID     int64  `json:"gameId"`
		Champion   int    `json:"champion"`
		Queue      int    `json:"queue"`
		Season     int    `json:"season"`
		Timestamp  int64  `json:"timestamp"`
		Role       string `json:"role"`
		Lane       string `json:"lane"`
	} `json:"matches"`
	StartIndex int `json:"startIndex"`
	EndIndex   int `json:"endIndex"`
	TotalGames int `json:"totalGames"`
}

// GetMatchList constructs a new MatchListDto using the AccountID from the accountID
// TODO: Make this function configurable
func GetMatchList(region string, accountID int64) (*MatchListDto, error) {
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/match/v3/matchlists/by-account/%d/recent", region, accountID)
	var matchList MatchListDto
	err := GetObj(url, &matchList)
	return &matchList, err
}
