package services

import "math"

type Pool struct {
	PoolID      uint8 // 255 pools max
	PoolState   State
	PoolMaxSize uint16
	PoolPlayers []Player
	PoolReferee Referee
}

func (p Pool) String() string {
	var s string

	s = string(p.PoolID) + " - " + p.PoolState.String()

	for _, player := range p.PoolPlayers {
		s = s + "\n" + player.String()
	}

	return s
}

func CreateStage(poolID uint8, poolMaxSize uint16) Pool {
	var p Pool

	p.PoolID = poolID
	p.PoolState = IDLE
	p.PoolMaxSize = poolMaxSize
	p.PoolPlayers = make([]Player, poolMaxSize)
	p.PoolReferee = Referee{}

	return p
}

func (p Pool) PlayerPosition(player Player) uint16 {
	for i, poolPlayer := range p.PoolPlayers {
		if poolPlayer.PlayerID == player.PlayerID {
			return uint16(i)
		}
	}

	return math.MaxInt16
}

func (p *Pool) AddPlayer(player Player) bool {
	if p.PoolState != IDLE {
		return false
	}

	if len(p.PoolPlayers) >= int(p.PoolMaxSize) {
		return false
	}

	if p.PlayerPosition(player) == math.MaxInt16 {
		p.PoolPlayers = append(p.PoolPlayers, player)
		return true
	}

	return false
}

func (p *Pool) RemovePlayer(player Player) bool {
	if p.PoolState != IDLE {
		return false
	}

	pos := p.PlayerPosition(player)

	if pos != math.MaxInt16 {
		p.PoolPlayers = append(p.PoolPlayers[:pos], p.PoolPlayers[pos+1:]...)
		return true
	}

	return false
}

func (p *Pool) SetReferee(referee Referee) {
	p.PoolReferee = referee
}
