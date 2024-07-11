package services

import "math"

type Pool struct {
	PoolID      uint8 // 255 pools max
	PoolState   State
	PoolMaxSize uint16
	PoolSize    uint16
	PoolPlayers map[uint16]*Player
	PoolReferee *Referee
}

type PoolStage struct {
	PoolStageID              uint8
	PoolStageSize            uint16
	PoolPlayers              map[uint16]*Player
	PoolStagePools           map[uint8]*Pool
	PoolStageState           State
	PoolEnteringPlayerNumber uint16
}

func (p Pool) String() string {
	var s string

	s = string(p.PoolID) + " - " + p.PoolState.String()

	for _, player := range p.PoolPlayers {
		s = s + "\n" + player.String()
	}

	return s
}

func CreatePoolStage(poolStageID uint8, enteringPlayerNumber uint16, leavingPlayerNumber uint16) PoolStage {
	var s PoolStage

	s.PoolStageID = poolStageID
	s.PoolEnteringPlayerNumber = enteringPlayerNumber
	s.PoolPlayers = make(map[uint16]*Player)
	s.PoolStagePools = make(map[uint8]*Pool)
	s.PoolStageState = IDLE
	s.PoolStageSize = 0

	return s
}

// From interface
func (p *PoolStage) PlayerPosition(player *Player) uint16 {
	for _, pool := range p.PoolStagePools {
		if pool.PlayerPosition(player) != math.MaxInt16 {
			return pool.PlayerPosition(player)
		}
	}

	return math.MaxInt16
}

// From interface
func (p *PoolStage) AddPlayer(player *Player) bool {
	if p.PoolStageState != REGISTERING {
		return false
	}

	if p.PlayerPosition(player) != math.MaxInt16 {
		return false
	}

	for _, pool := range p.PoolStagePools {
		if pool.AddPlayer(player) {
			p.PoolPlayers[player.PlayerID] = player
			p.PoolStageSize++

			pool.PoolSize++
			return true
		}
	}

	// Add pool
	pool := &Pool{
		PoolID:      uint8(len(p.PoolStagePools)),
		PoolState:   REGISTERING,
		PoolMaxSize: 7,
		PoolSize:    0,
		PoolPlayers: make(map[uint16]*Player),
		PoolReferee: nil,
	}

	pool.AddPlayer(player)

	p.AddPool(pool)
	p.PoolPlayers[player.PlayerID] = player

	return true
}

// From interface
func (p *PoolStage) RemovePlayer(player *Player) bool {
	if p.PoolStageState != REGISTERING {
		return false
	}

	if p.PlayerPosition(player) == math.MaxInt16 {
		return false
	}

	for _, pool := range p.PoolStagePools {
		if pool.RemovePlayer(player) {
			delete(p.PoolPlayers, player.PlayerID)
			p.PoolStageSize--

			pool.PoolSize--

			if pool.PoolSize == 0 {
				p.RemovePool(pool)
			}

			return true
		}
	}

	return false
}

// From interface
func (p *PoolStage) GetID() uint8 {
	return p.PoolStageID
}

// From interface
func (p *PoolStage) GetState() State {
	return p.PoolStageState
}

// From interface
func (p *PoolStage) GetPlayers() []*Player {
	players := []*Player{}

	for _, player := range p.PoolPlayers {
		players = append(players, player)
	}

	return players
}

// From interface
func (p *PoolStage) Register() bool {
	if p.PoolStageState != IDLE {
		return false
	}

	p.PoolStageState = REGISTERING

	return true
}

// From interface
func (p *PoolStage) Start() bool {
	if p.PoolStageState != REGISTERING {
		return false
	}

	if p.PoolStageSize < p.PoolEnteringPlayerNumber {
		return false
	}

	p.PoolStageState = STARTED

	return true
}

// From interface
func (p *PoolStage) End() bool {
	if p.PoolStageState != STARTED {
		return false
	}

	p.PoolStageState = FINISHED

	return true
}

// From interface
func (p *PoolStage) Lock() bool {
	if p.PoolStageState != FINISHED {
		return false
	}

	p.PoolStageState = LOCKED

	return true
}

// From interface
func (p *PoolStage) Build() bool {
	// TODO: Implement
	return false
}

func (p *PoolStage) AddPool(pool *Pool) bool {
	if p.PoolStageState != REGISTERING {
		return false
	}

	if pool == nil {
		return false
	}

	p.PoolStagePools[pool.PoolID] = pool

	return true
}

func (p *PoolStage) RemovePool(pool *Pool) bool {
	if p.PoolStageState != REGISTERING {
		return false
	}

	if pool == nil {
		return false
	}

	if pool.PoolSize > 0 {
		return false
	}

	delete(p.PoolStagePools, pool.PoolID)

	return true
}

func (p *PoolStage) GetPool(poolID uint8) *Pool {
	pool, ok := p.PoolStagePools[poolID]
	if ok {
		return pool
	}

	return nil
}

// ------------------------------ Pool ------------------------------
func (p Pool) PlayerPosition(player *Player) uint16 {
	_, ok := p.PoolPlayers[player.PlayerID]
	if ok {
		return p.PoolPlayers[player.PlayerID].PlayerID
	}

	return math.MaxInt16
}

func (p *Pool) AddPlayer(player *Player) bool {
	if p.PoolState != REGISTERING {
		return false
	}

	if p.PoolSize >= p.PoolMaxSize {
		return false
	}

	if p.PlayerPosition(player) != math.MaxInt16 {
		return false
	}

	p.PoolPlayers[p.PoolSize] = player
	p.PoolSize++

	return true
}

func (p *Pool) RemovePlayer(player *Player) bool {
	if p.PoolState != REGISTERING {
		return false
	}

	if p.PlayerPosition(player) == math.MaxInt16 {
		return false
	}

	delete(p.PoolPlayers, player.PlayerID)

	return true
}

func (p *Pool) SetReferee(referee *Referee) bool {
	if p.PoolState != REGISTERING {
		return false
	}

	p.PoolReferee = referee

	return true
}
