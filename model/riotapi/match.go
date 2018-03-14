package riotapi

import "fmt"

// MatchDto for grabbing a MatchDto from the Riot API
type MatchDto struct {
	GameID       int64  `json:"gameId"`
	PlatformID   string `json:"platformId"`
	GameCreation int64  `json:"gameCreation"`
	GameDuration int    `json:"gameDuration"`
	QueueID      int    `json:"queueId"`
	MapID        int    `json:"mapId"`
	SeasonID     int    `json:"seasonId"`
	GameVersion  string `json:"gameVersion"`
	GameMode     string `json:"gameMode"`
	GameType     string `json:"gameType"`
	Teams        []struct {
		TeamID               int    `json:"teamId"`
		Win                  string `json:"win"`
		FirstBlood           bool   `json:"firstBlood"`
		FirstTower           bool   `json:"firstTower"`
		FirstInhibitor       bool   `json:"firstInhibitor"`
		FirstBaron           bool   `json:"firstBaron"`
		FirstDragon          bool   `json:"firstDragon"`
		FirstRiftHerald      bool   `json:"firstRiftHerald"`
		TowerKills           int    `json:"towerKills"`
		InhibitorKills       int    `json:"inhibitorKills"`
		BaronKills           int    `json:"baronKills"`
		DragonKills          int    `json:"dragonKills"`
		VilemawKills         int    `json:"vilemawKills"`
		RiftHeraldKills      int    `json:"riftHeraldKills"`
		DominionVictoryScore int    `json:"dominionVictoryScore"`
		Bans                 []struct {
			ChampionID int `json:"championId"`
			PickTurn   int `json:"pickTurn"`
		} `json:"bans"`
	} `json:"teams"`
	Participants []struct {
		ParticipantID             int    `json:"participantId"`
		TeamID                    int    `json:"teamId"`
		ChampionID                int    `json:"championId"`
		Spell1ID                  int    `json:"spell1Id"`
		Spell2ID                  int    `json:"spell2Id"`
		HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier"`
		Stats                     struct {
			ParticipantID                   int  `json:"participantId"`
			Win                             bool `json:"win"`
			Item0                           int  `json:"item0"`
			Item1                           int  `json:"item1"`
			Item2                           int  `json:"item2"`
			Item3                           int  `json:"item3"`
			Item4                           int  `json:"item4"`
			Item5                           int  `json:"item5"`
			Item6                           int  `json:"item6"`
			Kills                           int  `json:"kills"`
			Deaths                          int  `json:"deaths"`
			Assists                         int  `json:"assists"`
			LargestKillingSpree             int  `json:"largestKillingSpree"`
			LargestMultiKill                int  `json:"largestMultiKill"`
			KillingSprees                   int  `json:"killingSprees"`
			LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
			DoubleKills                     int  `json:"doubleKills"`
			TripleKills                     int  `json:"tripleKills"`
			QuadraKills                     int  `json:"quadraKills"`
			PentaKills                      int  `json:"pentaKills"`
			UnrealKills                     int  `json:"unrealKills"`
			TotalDamageDealt                int  `json:"totalDamageDealt"`
			MagicDamageDealt                int  `json:"magicDamageDealt"`
			PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
			TrueDamageDealt                 int  `json:"trueDamageDealt"`
			LargestCriticalStrike           int  `json:"largestCriticalStrike"`
			TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
			MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
			PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
			TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
			TotalHeal                       int  `json:"totalHeal"`
			TotalUnitsHealed                int  `json:"totalUnitsHealed"`
			DamageSelfMitigated             int  `json:"damageSelfMitigated"`
			DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
			DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
			VisionScore                     int  `json:"visionScore"`
			TimeCCingOthers                 int  `json:"timeCCingOthers"`
			TotalDamageTaken                int  `json:"totalDamageTaken"`
			MagicalDamageTaken              int  `json:"magicalDamageTaken"`
			PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
			TrueDamageTaken                 int  `json:"trueDamageTaken"`
			GoldEarned                      int  `json:"goldEarned"`
			GoldSpent                       int  `json:"goldSpent"`
			TurretKills                     int  `json:"turretKills"`
			InhibitorKills                  int  `json:"inhibitorKills"`
			TotalMinionsKilled              int  `json:"totalMinionsKilled"`
			NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
			NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
			NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
			TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
			ChampLevel                      int  `json:"champLevel"`
			VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
			SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
			WardsPlaced                     int  `json:"wardsPlaced"`
			WardsKilled                     int  `json:"wardsKilled"`
			FirstBloodKill                  bool `json:"firstBloodKill"`
			FirstBloodAssist                bool `json:"firstBloodAssist"`
			FirstTowerKill                  bool `json:"firstTowerKill"`
			FirstTowerAssist                bool `json:"firstTowerAssist"`
			CombatPlayerScore               int  `json:"combatPlayerScore"`
			ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
			TotalPlayerScore                int  `json:"totalPlayerScore"`
			TotalScoreRank                  int  `json:"totalScoreRank"`
			PlayerScore0                    int  `json:"playerScore0"`
			PlayerScore1                    int  `json:"playerScore1"`
			PlayerScore2                    int  `json:"playerScore2"`
			PlayerScore3                    int  `json:"playerScore3"`
			PlayerScore4                    int  `json:"playerScore4"`
			PlayerScore5                    int  `json:"playerScore5"`
			PlayerScore6                    int  `json:"playerScore6"`
			PlayerScore7                    int  `json:"playerScore7"`
			PlayerScore8                    int  `json:"playerScore8"`
			PlayerScore9                    int  `json:"playerScore9"`
			Perk0                           int  `json:"perk0"`
			Perk0Var1                       int  `json:"perk0Var1"`
			Perk0Var2                       int  `json:"perk0Var2"`
			Perk0Var3                       int  `json:"perk0Var3"`
			Perk1                           int  `json:"perk1"`
			Perk1Var1                       int  `json:"perk1Var1"`
			Perk1Var2                       int  `json:"perk1Var2"`
			Perk1Var3                       int  `json:"perk1Var3"`
			Perk2                           int  `json:"perk2"`
			Perk2Var1                       int  `json:"perk2Var1"`
			Perk2Var2                       int  `json:"perk2Var2"`
			Perk2Var3                       int  `json:"perk2Var3"`
			Perk3                           int  `json:"perk3"`
			Perk3Var1                       int  `json:"perk3Var1"`
			Perk3Var2                       int  `json:"perk3Var2"`
			Perk3Var3                       int  `json:"perk3Var3"`
			Perk4                           int  `json:"perk4"`
			Perk4Var1                       int  `json:"perk4Var1"`
			Perk4Var2                       int  `json:"perk4Var2"`
			Perk4Var3                       int  `json:"perk4Var3"`
			Perk5                           int  `json:"perk5"`
			Perk5Var1                       int  `json:"perk5Var1"`
			Perk5Var2                       int  `json:"perk5Var2"`
			Perk5Var3                       int  `json:"perk5Var3"`
			PerkPrimaryStyle                int  `json:"perkPrimaryStyle"`
			PerkSubStyle                    int  `json:"perkSubStyle"`
		} `json:"stats"`
		Timeline struct {
			ParticipantID               int                `json:"participantId"`
			CreepsPerMinDeltas          map[string]float64 `json:"creepsPerMinDeltas"`
			XpPerMinDeltas              map[string]float64 `json:"xpPerMinDeltas"`
			GoldPerMinDeltas            map[string]float64 `json:"goldPerMinDeltas"`
			CsDiffPerMinDeltas          map[string]float64 `json:"csDiffPerMinDeltas"`
			XpDiffPerMinDeltas          map[string]float64 `json:"xpDiffPerMinDeltas"`
			DamageTakenPerMinDeltas     map[string]float64 `json:"damageTakenPerMinDeltas"`
			DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
			Role                        string             `json:"role"`
			Lane                        string             `json:"lane"`
		} `json:"timeline"`
	} `json:"participants"`
	ParticipantIdentities []struct {
		ParticipantID int `json:"participantId"`
		Player        struct {
			PlatformID        string `json:"platformId"`
			AccountID         int    `json:"accountId"`
			SummonerName      string `json:"summonerName"`
			SummonerID        int    `json:"summonerId"`
			CurrentPlatformID string `json:"currentPlatformId"`
			CurrentAccountID  int    `json:"currentAccountId"`
			MatchHistoryURI   string `json:"matchHistoryUri"`
			ProfileIcon       int    `json:"profileIcon"`
		} `json:"player"`
	} `json:"participantIdentities"`
}

// GetMatchDto grabs a MatchDto from the Riot API
func GetMatchDto(matchID int64) (*MatchDto, error) {
	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/match/v3/matches/%d", matchID)
	var match MatchDto
	err := GetObj(url, &match)
	return &match, err
}
