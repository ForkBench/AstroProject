package services

import (
	"math/rand/v2"

	"github.com/go-faker/faker/v4"

	"astroproject/astro/structs"
)

// PersonManager : PersonManager details
type PersonManager struct {
	UserSession *Session
}

func (pm *PersonManager) GenerateRandomID(competitionID uint8) uint16 {
	// PlayerID is 16 bits, 4 are for the competitionID and 12 are for the playerID (4095 players max)
	playerID := uint16(competitionID) << 12

	// Pick a number between 0 and 4095
	playerID += uint16(rand.Uint32()) % 4095

	return playerID
}

func (pm *PersonManager) GenerateRandomPlayer(competitionID uint8) structs.Player {

	return structs.Player{
		PlayerID:          pm.GenerateRandomID(competitionID),
		PlayerFirstname:   faker.FirstName(),
		PlayerLastname:    faker.LastName(),
		PlayerClub:        structs.GenerateClub(),
		PlayerNation:      structs.GenerateNation(),
		PlayerRegion:      structs.GenerateRegion(),
		PlayerInitialRank: uint16(rand.Uint32()/2) % 100,
	}
}
