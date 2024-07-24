package structs

// Stage : Stage details
type Stage interface {
	PlayerPosition(player Player) uint16
	AddPlayer(player Player) bool
	RemovePlayer(player Player) bool
	GetID() uint8
	GetState() State
	GetKind() StageKind // TODO: Implement in pools
	GetPlayers() []Player
	UpdatePlayer(player Player) bool
	Register() bool
	Start() bool
	End() bool
	Lock() bool
	Build() bool
}

// Stage Kind
type StageKind uint8

const (
	REGISTRATIONS     StageKind = iota // 0
	POOLS                              // 1
	POOLSRESULTS                       // 2
	DIRECTELIMINATION                  // 3
	FINALRANKING                       // 4
	UNKNOWN                            // 5
)

func (sk StageKind) String() string {
	return [...]string{"REGISTRATIONS", "POOLS", "POOLSRESULTS", "DIRECTELIMINATION", "FINALRANKING"}[sk] // TODO: Do the same with Util?
}
