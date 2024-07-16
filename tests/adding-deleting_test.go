package tests

import (
	"testing"

	"astroproject/astro/services"
)

const CompetitionNumber = 100

func SessionTesting(t *testing.T) {
	session := services.Session{
		CompetitionNumber: 0,
		Competitions:      map[uint8]*services.Competition{},
	}

	for i := 0; i < CompetitionNumber; i++ {
		session.AddCompetition("Test", "U17", "Foil")
	}

	if session.CompetitionNumber != CompetitionNumber {
		t.Errorf("Expected 5 competitions, got %d", session.CompetitionNumber)
	}

	for _, competition := range session.Competitions {
		session.RemoveCompetition(competition.CompetitionID)
	}

	if session.CompetitionNumber != 0 {
		t.Errorf("Expected 0 competitions, got %d", session.CompetitionNumber)
	}
}

func CompetitionTesting(t *testing.T) {
	competition := services.CreateCompetition(0, "Test", services.CreateCategory("U17"), services.CreateWeapon("Foil"), services.MinStageSize)

	if competition.CompetitionID != 0 {
		t.Errorf("Expected competition ID to be 0, got %d", competition.CompetitionID)
	}

	if competition.CompetitionName != "Test" {
		t.Errorf("Expected competition name to be Test, got %s", competition.CompetitionName)
	}

	if competition.CompetitionCategory.String() != "U17" {
		t.Errorf("Expected competition category to be U17, got %s", competition.CompetitionCategory.String())
	}

	if competition.CompetitionWeapon.String() != "Foil" {
		t.Errorf("Expected competition weapon to be Foil, got %s", competition.CompetitionWeapon.String())
	}

	if len(competition.CompetitionStages) != services.MinStageSize {
		t.Errorf("Expected competition to have %d stages, got %d", services.MinStageSize, len(competition.CompetitionStages))
	}
}
