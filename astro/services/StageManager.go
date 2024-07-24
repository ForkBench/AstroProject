package services

import "astroproject/astro/structs"

// StageManager : StageManager details
type StageManager struct {
	UserSession *Session
}

func (sm *StageManager) GetStageKind(competitionID uint8, stageID uint8) structs.StageKind {
	competition := sm.UserSession.GetCompetition(competitionID)
	if competition == nil {
		return structs.UNKNOWN
	}

	stage := competition.GetStage(stageID)
	if stage == nil {
		return structs.UNKNOWN
	}

	return (*stage).GetKind()
}

func (sm *StageManager) AddPlayerToCompetitionStage(competitionID uint8, stageID uint8, player *structs.Player) bool {
	competition := sm.UserSession.GetCompetition(competitionID)
	if competition == nil {
		return false
	}

	stage := competition.GetStage(stageID)
	if stage == nil {
		return false
	}

	return (*stage).AddPlayer(player)
}

func (sm *StageManager) RemovePlayerFromCompetitionStage(competitionID uint8, stageID uint8, player *structs.Player) bool {
	competition := sm.UserSession.GetCompetition(competitionID)
	if competition == nil {
		return false
	}

	stage := competition.GetStage(stageID)
	if stage == nil {
		return false
	}

	return (*stage).RemovePlayer(player)
}

func (sm *StageManager) GetPlayersFromCompetitionStage(competitionID uint8, stageID uint8) []*structs.Player {
	competition := sm.UserSession.GetCompetition(competitionID)
	if competition == nil {
		return []*structs.Player{}
	}

	stage := competition.GetStage(stageID)
	if stage == nil {
		return []*structs.Player{}
	}

	return (*stage).GetPlayers()
}
