package services

const MinStageSize = 3

// Session : Session details
type Session struct {
	Competitions      [](*Competition)
	CompetitionNumber uint8
}

func (s *Session) AddCompetition(name string, category string, weapon string) {
	competition := CreateCompetition(
		s.CompetitionNumber,
		name,
		CreateCategory(category),
		CreateWeapon(weapon),
		MinStageSize)

	s.Competitions = append(s.Competitions, &competition)
	s.CompetitionNumber++
}

func (s *Session) RemoveCompetition(competitionID uint8) {
	for i, competition := range s.Competitions {
		if competition.CompetitionID == competitionID {
			s.Competitions = append(s.Competitions[:i], s.Competitions[i+1:]...)
		}
	}
	s.CompetitionNumber--
}

func (s *Session) GetCompetitions() []Competition {
	comps := []Competition{}
	for _, competition := range s.Competitions {
		comps = append(comps, *competition)
	}
	return comps
}

func (s *Session) GetCompetition(competitionID uint8) *Competition {
	for _, competition := range s.Competitions {
		if competition.CompetitionID == competitionID {
			return competition
		}
	}
	return nil
}

func (s *Session) AddPlayerToCompetition(competitionID uint8, player *Player) bool {
	competition := s.GetCompetition(competitionID)
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

func (s *Session) RemovePlayerFromCompetition(competitionID uint8, player *Player) bool {
	competition := s.GetCompetition(competitionID)
	if competition.CompetitionName == "" {
		return false
	}

	return competition.RemovePlayer(player)
}

func (s *Session) GetAllPlayersFromCompetition(competitionID uint8) []*Player {
	competition := s.GetCompetition(competitionID)
	if competition.CompetitionName == "" {
		return []*Player{}
	}

	players := []*Player{}
	for _, player := range competition.CompetitionPlayers {
		players = append(players, player)
	}
	return players
}

func (s *Session) UpdateCompetitionPlayer(competitionID uint8, player *Player) bool {
	competition := s.GetCompetition(competitionID)
	if competition.CompetitionName == "" {
		return false
	}

	return competition.UpdatePlayer(player)
}
