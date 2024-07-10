package services

const MinStageSize = 3

// Session : Session details
type Session struct {
	Competitions      []Competition
	CompetitionNumber uint8
}

func (s *Session) AddCompetition(name string, category string, weapon string) {
	competition := CreateCompetition(
		s.CompetitionNumber,
		name,
		CreateCategory(category),
		CreateWeapon(weapon),
		MinStageSize)

	s.Competitions = append(s.Competitions, competition)
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
	return s.Competitions
}

func (s *Session) GetCompetition(competitionID uint8) Competition {
	for _, competition := range s.Competitions {
		if competition.CompetitionID == competitionID {
			return competition
		}
	}
	return Competition{}
}

func (s *Session) AddPlayerToCompetition(competitionID uint8, player Player) bool {
	competition := s.GetCompetition(competitionID)
	if competition.CompetitionName == "" {
		return false
	}

	return competition.AddPlayer(player)
}

func (s *Session) RemovePlayerFromCompetition(competitionID uint8, player Player) bool {
	competition := s.GetCompetition(competitionID)
	if competition.CompetitionName == "" {
		return false
	}

	return competition.RemovePlayer(player)
}
