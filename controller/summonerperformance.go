package controller

import (
	"math"

	"github.com/joshuaj1397/LoLMemes/model/riotapi"
)

// TODO: Create some more shameful stats
// PerformanceDto is a Dto for storing the aggregate amount of shortcomings
// a summoner may have in their recent games.
// All shortcomings are averaged not cumulated so they can be compared to the
// overall average for a player of the same skill later.
type PerformanceDto struct {
	SummonerName          string  // [x] Summoner Name
	KDA                   float64 // [x] Kill-Death-Assist ratio
	WinLoss               float64 // [x] Win-Loss ratio
	CS                    int32   // Average Creep Score
	BossKillsJg           float64 // Average neutral boss kills as a jg
	VisionScoreSupp       int32   // Vision Score as a support
	SelfMitigatedDmgTank  int32   // Self Mitigated Damage as a tank
	MagicDmgMage          int32   // Magic Damage as a mage
	PhysicalDmgAdc        int32   // Physical Damage as an AD carry
	ChampLevelDifference  int     // The difference between summoner and others champ levels
	BannedButStillClapped bool    // If you constantly ban someone and you're still being clapped by someone else
}

func (perf *PerformanceDto) setKDA(totalKDA float64, numOfGames int) {
	perf.KDA = totalKDA / float64(numOfGames)
}

func (perf *PerformanceDto) setWinLoss(wins, losses int) {
	perf.WinLoss = float64(wins) / float64(wins+losses)
}

func (perf *PerformanceDto) setBossKillsDelta(bossKillsDelta float64, numOfGames int) {
	perf.BossKillsJg = bossKillsDelta / float64(numOfGames)
}

func calcBossKillDeltas(bossKills, enemyBossKills int) float64 {
	return math.Abs(float64(bossKills - enemyBossKills))
}

// TODO: Dry this function
// GetRecentPerformance gets the last 20 games and calculates the aggregate
// performance of a summoner
func GetRecentPerformance(region *string, summonerName string) (*PerformanceDto, error) {
	var matchList *riotapi.MatchListDto
	var perf PerformanceDto
	var numOfGames, wins, losses, discardedGames, bossKills, enemyBossKills int
	var totalKDA, bossKillsDelta float64

	s, summonerErr := riotapi.GetSummoner(region, summonerName)
	if summonerErr != nil {
		return nil, summonerErr
	}

	matchList, matchListErr := riotapi.GetMatchList(*region, s.AccountID)
	if matchListErr != nil {
		return nil, matchListErr
	}

	// i is the position of the match, m is the match in the list of matches
	for i, m := range matchList.Matches {

		// Only get 5v5 Summoners Rift matches
		switch m.Queue {
		case riotapi.SUMMONERS_RIFT_BLIND:
		case riotapi.SUMMONERS_RIFT_DRAFT:
		case riotapi.SUMMONERS_RIFT_RANKED:
		default:
			discardedGames++
			continue
		}

		var participantID int
		match, matchErr := riotapi.GetMatch(*region, m.GameID)
		if matchErr != nil {
			return nil, matchErr
		}

		// Find the summoner's participantID
		for _, summoner := range match.ParticipantIdentities {
			if summoner.Player.SummonerName == s.Name {
				participantID = summoner.ParticipantID
				break
			}
		}

		for _, summoner := range match.Participants {
			if summoner.ParticipantID == participantID {

				// Aggregate the KDA
				if summoner.Stats.Deaths == 0 {
					totalKDA += float64(summoner.Stats.Assists + summoner.Stats.Kills)
				} else {
					totalKDA += float64(summoner.Stats.Assists+summoner.Stats.Kills) / float64(summoner.Stats.Deaths)
				}

				// Aggregate the wins and losses
				if summoner.Stats.Win {
					wins++
				} else {
					losses++
				}

				// Aggregate jg boss secures
				for _, team := range match.Teams {
					if team.TeamID == summoner.TeamID {
						bossKills += team.BaronKills + team.DragonKills
					} else {
						enemyBossKills += team.BaronKills + team.DragonKills
					}
				}
				bossKillsDelta += calcBossKillDeltas(bossKills, enemyBossKills)

			}
		}
		numOfGames = i + 1 - discardedGames
	}
	perf.setBossKillsDelta(bossKillsDelta, numOfGames)
	perf.SummonerName = s.Name
	perf.setKDA(totalKDA, numOfGames)
	perf.setWinLoss(wins, losses)
	return &perf, nil
}
