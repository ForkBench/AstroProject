package services

// Player : Person details
type Player struct {
	PlayerID        uint16 // More than 255 players
	PlayerFirstname string
	PlayerLastname  string
	PlayerNationID  uint16
	PlayerRegionID  uint16
	PlayerClubID    uint16
}

func (p Player) String() string {
	return "Player : " + p.PlayerFirstname + " " + p.PlayerLastname
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
