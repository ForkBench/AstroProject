package services

const MinStageSize = 3

// Session : Session details
type Session struct {
	Competitions      map[uint8]*Competition
	CompetitionNumber uint8
}

func (s *Session) AddCompetition(name string, category string, weapon string) *Competition {

	id := uint8(0)

	for {
		// Check if the competition already exists
		if _, ok := s.Competitions[id]; !ok {
			break
		}
		id++
	}

	competition := CreateCompetition(
		id,
		name,
		CreateCategory(category),
		CreateWeapon(weapon),
		MinStageSize)

	s.Competitions[competition.CompetitionID] = competition
	s.CompetitionNumber++

	return competition
}

func (s *Session) RemoveCompetition(competitionID uint8) bool {
	if _, ok := s.Competitions[competitionID]; !ok {
		return false
	}

	delete(s.Competitions, competitionID)

	s.CompetitionNumber--

	return true
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
