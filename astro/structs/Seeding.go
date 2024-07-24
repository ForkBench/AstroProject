package structs

// Seeding : Seeding details
type Seeding struct {
	SeedingPosition      uint16
	SeedingPlayer        Player
	SeedingPlayerPresent bool
}

type SeedingStage struct {
	SeedingStageID              uint8
	SeedingSize                 uint16
	SeedingSeedings             map[uint16]*Seeding
	SeedingState                State
	SeedingEnteringPlayerNumber uint16
}

func (s *SeedingStage) String() string {
	var str string

	str = string(s.SeedingStageID) + " - " + s.SeedingState.String()

	for _, seeding := range s.SeedingSeedings {
		str = str + "\n" + seeding.SeedingPlayer.String()
	}

	return str
}

func CreateSeedingStage(seedingStageID uint8, enteringPlayerNumber uint16, leavingPlayerNumber uint16) *SeedingStage {
	var s SeedingStage

	s.SeedingStageID = seedingStageID
	s.SeedingEnteringPlayerNumber = enteringPlayerNumber
	s.SeedingSeedings = make(map[uint16]*Seeding)
	s.SeedingState = IDLE
	s.SeedingSize = 0

	return &s
}

// From interface
func (s *SeedingStage) PlayerPosition(player Player) uint16 {
	return s.SeedingSeedings[player.PlayerID].SeedingPosition
}

// From interface
func (s *SeedingStage) AddPlayer(player Player) bool {
	if s.SeedingState != REGISTERING {
		return false
	}

	if s.SeedingSeedings[player.PlayerID] != nil {
		return false
	}

	s.SeedingSeedings[player.PlayerID] = &Seeding{
		SeedingPosition: s.SeedingSize,
		SeedingPlayer:   player,
	}

	s.SeedingSize++

	return true
}

// From interface
func (s *SeedingStage) RemovePlayer(player Player) bool {
	if s.SeedingState != REGISTERING {
		return false
	}

	if s.SeedingSeedings[player.PlayerID] == nil {
		return false
	}

	delete(s.SeedingSeedings, player.PlayerID)

	s.SeedingSize--

	return true
}

// From interface
func (s *SeedingStage) UpdatePlayer(player Player) bool {
	if s.SeedingState != REGISTERING {
		return false
	}

	if s.SeedingSeedings[player.PlayerID] == nil {
		return false
	}

	// TODO: Refactor
	s.SeedingSeedings[player.PlayerID].SeedingPlayer.PlayerFirstname = player.PlayerFirstname
	s.SeedingSeedings[player.PlayerID].SeedingPlayer.PlayerLastname = player.PlayerLastname
	s.SeedingSeedings[player.PlayerID].SeedingPlayer.PlayerNation = player.PlayerNation
	s.SeedingSeedings[player.PlayerID].SeedingPlayer.PlayerRegion = player.PlayerRegion
	s.SeedingSeedings[player.PlayerID].SeedingPlayer.PlayerClub = player.PlayerClub
	s.SeedingSeedings[player.PlayerID].SeedingPlayer.PlayerInitialRank = player.PlayerInitialRank

	return true
}

// From interface
func (s *SeedingStage) GetID() uint8 {
	return s.SeedingStageID
}

// From interface
func (s *SeedingStage) GetState() State {
	return s.SeedingState
}

// From interface
func (s *SeedingStage) GetKind() StageKind {
	return REGISTRATIONS
}

// From interface
func (s *SeedingStage) GetPlayers() []Player {
	players := []Player{}

	for _, seeding := range s.SeedingSeedings {
		players = append(players, seeding.SeedingPlayer)
	}

	return players
}

// From interface
func (s *SeedingStage) Register() bool {
	if s.SeedingState != IDLE {
		return false
	}

	s.SeedingState = REGISTERING

	return true
}

// From interface
func (s *SeedingStage) Start() bool {
	if s.SeedingState != REGISTERING {
		return false
	}

	if s.SeedingSize != s.SeedingEnteringPlayerNumber {
		return false
	}

	s.SeedingState = STARTED

	return true
}

// From interface
func (s *SeedingStage) End() bool {
	if s.SeedingState != STARTED {
		return false
	}

	s.SeedingState = FINISHED

	return true
}

// From interface
func (s *SeedingStage) Lock() bool {
	if s.SeedingState != FINISHED {
		return false
	}

	s.SeedingState = LOCKED

	return true
}

// From interface
func (s *SeedingStage) Build() bool {
	if s.SeedingState != REGISTERING {
		return false
	}

	if s.SeedingSize != s.SeedingEnteringPlayerNumber {
		return false
	}

	return true
}
