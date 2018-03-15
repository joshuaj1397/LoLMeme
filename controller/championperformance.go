package main

import (
	"github.com/joshuaj1397/LoLMemes/model/riotapi"
)

// TODO: Create some more shameful stats
// PerformanceDto is a Dto for storing the aggregate amount of shortcomings
// a summoner may have in their recent games.
// All shortcomings are averaged not cumulated so they can be compared to the
// overall average for a player of the same skill later.
type PerformanceDto struct {
	KDA                   float32 // Kill-Death-Assist ratio
	WinLoss               float32 // Win-Loss ratio
	WinLossWithPremades   float32 // Win-Loss with a particular premade ratio
	CS                    int32   // Creep Score
	BossKillsJg           int32   // Neutral Boss kills as a jungler
	VisionScoreSupp       int32   // Vision Score as a support
	SelfMitigatedDmgTank  int32   // Self Mitigated Damage as a tank
	MagicDmgMage          int32   // Magic Damage as a mage
	PhysicalDmgAdc        int32   // Physical Damage as an AD carry
	ChampLevelDifference  int     // The difference between summoner and others champ levels
	BannedButStillClapped bool    // If you constantly ban someone and you're still being clapped by someone else
}

// TODO: Dry this function
// GetRecentPerformance gets the last 20 games and calculates the aggregate
// performance of a summoner
func GetRecentPerformance(s riotapi.SummonerDto) (*PerformanceDto, error) {
	var matchList *riotapi.MatchListDto
	var perf *PerformanceDto

	matchList, err := riotapi.GetMatchList(s.AccountID)
	if err != nil {
		return nil, err
	}

	matches := matchList.Matches
	for i, m := range matches {
		match, err := riotapi.GetMatchDto(m.GameID)
		var participantID int
		if err != nil {
			return nil, err
		}

		summoners := match.ParticipantIdentities
		for _, summoner := range summoners {

			// Find the user summoner
			if summoner.Player.SummonerName == s.Name {
				participantID = summoner.ParticipantID
			}
		}
	}

	return perf, err
}
