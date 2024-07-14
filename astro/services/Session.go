package services

import "fmt"

const MinStageSize = 3

// Session : Session details
type Session struct {
	Competitions      [](*Competition)
	CompetitionNumber uint8
}

func (s *Session) AddCompetition(name string, category string, weapon string) bool {
	fmt.Println("Adding competition")
	competition := CreateCompetition(
		s.CompetitionNumber,
		name,
		CreateCategory(category),
		CreateWeapon(weapon),
		MinStageSize)

	s.Competitions = append(s.Competitions, &competition)
	s.CompetitionNumber++

	return true
}

func (s *Session) RemoveCompetition(competitionID uint8) bool {
	for i, competition := range s.Competitions {
		if competition.CompetitionID == competitionID {
			s.Competitions = append(s.Competitions[:i], s.Competitions[i+1:]...)
		}
	}
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

func (s *Session) GetLastCompetition() *Competition {
	if len(s.Competitions) > 0 {
		return s.Competitions[len(s.Competitions)-1]
	}
	return nil
}
