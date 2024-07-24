package services

import "astroproject/astro/structs"

const MinStageSize = 3
const MaxCompetitions = 16 // 4 bits

// Session : Session details
type Session struct {
	Competitions      [](*structs.Competition)
	CompetitionNumber uint8
}

func (s *Session) AddCompetition(name string, category string, weapon string) {

	if s.CompetitionNumber >= MaxCompetitions {
		return
	}

	competition := structs.CreateCompetition(
		s.CompetitionNumber,
		name,
		structs.CreateCategory(category),
		structs.CreateWeapon(weapon),
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

func (s *Session) GetCompetitions() []structs.Competition {
	comps := []structs.Competition{}
	for _, competition := range s.Competitions {
		comps = append(comps, *competition)
	}
	return comps
}

func (s *Session) GetCompetition(competitionID uint8) *structs.Competition {

	if competitionID >= s.CompetitionNumber {
		return nil
	}

	for _, competition := range s.Competitions {
		if competition.CompetitionID == competitionID {
			return competition
		}
	}
	return nil
}
