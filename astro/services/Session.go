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
