package services

// Stage : Stage details
type Stage interface {
	CreateStage(stageID uint8, stageMaxSize uint16) Stage
	PlayerPosition(player Player) uint16
	AddPlayer(player Player) bool
	RemovePlayer(player Player) bool
	GetID() uint8
}
