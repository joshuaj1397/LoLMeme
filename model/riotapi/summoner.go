package riotapi

// SummonerDto for grabbing a Summoner from the Riot API
type SummonerDto struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	SummonerLevel int32  `json:"summonerLevel"`
	RevisionDate  int32  `json:"revisionDate"`
	ID            int32  `json:"id"`
	AccountID     int32  `json:"acccountId"`
}

// GetSummoner using Riot's official API
func GetSummoner(summonerName string) (*SummonerDto, error) {
	url := "https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/" + summonerName
	var s SummonerDto
	err := GetObj(url, &s)
	return &s, err
}
