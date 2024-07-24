package services

import "astroproject/astro/structs"

/*
For now, competitionManager has a UserSession field, which is a pointer to a Session struct.
I guess it's better for now to keep it like this insted of having a Manager in Session struct.
*/
type CompetitionManager struct {
	UserSession *Session
}

func (cm *CompetitionManager) AddPlayerToCompetition(competitionID uint8, player *structs.Player) bool {
	competition := cm.UserSession.GetCompetition(competitionID)
	if competition == nil {
		return false
	}

	if player == nil {
		return false
	}

	// If player id is UINT16MAX, set it
	var id uint16 = 0
	for {
		// Check if id is already taken
		_, ok := competition.CompetitionPlayers[id]
		if !ok {
			break
		}
		id++
	}

	player.PlayerID = id

	return competition.AddPlayer(player)
}

func (cm *CompetitionManager) RemovePlayerFromCompetition(competitionID uint8, player *structs.Player) bool {
	competition := cm.UserSession.GetCompetition(competitionID)
	if competition == nil {
		return false
	}

	return competition.RemovePlayer(player)
}

func (cm *CompetitionManager) GetAllPlayersFromCompetition(competitionID uint8) []*structs.Player {
	competition := cm.UserSession.GetCompetition(competitionID)
	if competition == nil {
		return []*structs.Player{}
	}

	players := []*structs.Player{}
	for _, player := range competition.CompetitionPlayers {
		players = append(players, player)
	}
	return players
}

func (cm *CompetitionManager) UpdateCompetitionPlayer(competitionID uint8, player *structs.Player) bool {
	competition := cm.UserSession.GetCompetition(competitionID)
	if competition == nil {
		return false
	}

	// TODO: Check if player is in competition

	return competition.UpdatePlayer(player)
}
