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

func (pm *PersonManager) GenerateRandomPlayer() *structs.Player {
	return &structs.Player{
		// PlayerID:          65535, Make it random
		PlayerID:          uint16(rand.Uint32()),
		PlayerFirstname:   faker.FirstName(),
		PlayerLastname:    faker.LastName(),
		PlayerClub:        structs.GenerateClub(),
		PlayerNation:      structs.GenerateNation(),
		PlayerRegion:      structs.GenerateRegion(),
		PlayerInitialRank: uint16(rand.Uint32()/2) % 100,
	}
}
