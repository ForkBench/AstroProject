package services

// Stage : Stage details
type Stage interface {
	PlayerPosition(player *Player) uint16
	AddPlayer(player *Player) bool
	RemovePlayer(player *Player) bool
	GetID() uint8
	GetState() State
	GetPlayers() []*Player
	Register() bool
	Start() bool
	End() bool
	Lock() bool
	Build() bool
}
