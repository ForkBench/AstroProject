package services

// Category : Age category
type Category uint8

const (
	U7      Category = iota // 0
	U9                      // 1
	U11                     // 2
	U13                     // 3
	U15                     // 4
	U17                     // 5
	U20                     // 6
	SENIOR                  // 7
	VETERAN                 // 8
)

func (c Category) String() string {
	switch c {
	case U7:
		return "U7"
	case U9:
		return "U9"
	case U11:
		return "U11"
	case U13:
		return "U13"
	case U15:
		return "U15"
	case U17:
		return "U17"
	case U20:
		return "U20"
	case SENIOR:
		return "Senior"
	case VETERAN:
		return "Veteran"
	default:
		return "Unknown"
	}
}

func CreateCategory(name string) Category {
	switch name {
	case "U7":
		return U7
	case "U9":
		return U9
	case "U11":
		return U11
	case "U13":
		return U13
	case "U15":
		return U15
	case "U17":
		return U17
	case "U20":
		return U20
	case "Senior":
		return SENIOR
	case "Veteran":
		return VETERAN
	default:
		return SENIOR
	}
}

// Weapon : Weapon used in competition
type Weapon uint8

const (
	FOIL  Weapon = iota // 0
	EPEE                // 1
	SABRE               // 2
)

func (w Weapon) String() string {
	switch w {
	case FOIL:
		return "Foil"
	case EPEE:
		return "Epee"
	case SABRE:
		return "Sabre"
	default:
		return "Unknown"
	}
}

func CreateWeapon(name string) Weapon {
	switch name {
	case "Foil":
		return FOIL
	case "Epee":
		return EPEE
	case "Sabre":
		return SABRE
	default:
		return FOIL
	}
}

// State : Competition state
type State uint8

const (
	REGISTERING State = iota // 0
	STARTED                  // 1
	FINISHED                 // 2
	IDLE                     // 3
)

func (s State) String() string {
	switch s {
	case REGISTERING:
		return "Registering"
	case STARTED:
		return "Started"
	case FINISHED:
		return "Finished"
	case IDLE:
		return "Idle"
	default:
		return "Unknown"
	}
}
