package structs

// Player : Person details
/*
PlayerID is 16 bits, 4 are for the competitionID and 12 are for the playerID (4095 players max)
*/
type Player struct {
	PlayerID          uint16 // More than 255 players
	PlayerFirstname   string
	PlayerLastname    string
	PlayerNation      Nation
	PlayerRegion      Region
	PlayerClub        Club
	PlayerInitialRank uint16
}

func (p Player) String() string {
	return "Player : " + p.PlayerFirstname + " " + p.PlayerLastname
}

func CreatePlayer(playerID uint16, playerFirstname string, playerLastname string) *Player {
	return &Player{
		PlayerID:        playerID,
		PlayerFirstname: playerFirstname,
		PlayerLastname:  playerLastname,
	}
}

// Referee : Person details
type Referee struct {
	RefereeID        uint16 // More than 255 referees
	RefereeFirstname string
	RefereeLastname  string
	RefereeNationID  uint16
	RefereeRegionID  uint16
	RefereeClubID    uint16
}

func (r Referee) String() string {
	return "Referee : " + r.RefereeFirstname + " " + r.RefereeLastname
}
