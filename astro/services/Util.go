package services

import "math/rand"

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
	LOCKED                   // 3
	IDLE                     // 4
)

func (s State) String() string {
	switch s {
	case REGISTERING:
		return "Registering"
	case STARTED:
		return "Started"
	case FINISHED:
		return "Finished"
	case LOCKED:
		return "Locked"
	case IDLE:
		return "Idle"
	default:
		return "Unknown"
	}
}

// RGB : RGB color
type RGB struct {
	R uint8
	G uint8
	B uint8
}

func NewRGB(r, g, b uint8) RGB {
	return RGB{R: r, G: g, B: b}
}

func (c RGB) String() string {
	return string(c.R) + ", " + string(c.G) + ", " + string(c.B)
}

// export function randomColor(brightness: number): string {
//   function randomChannel(brightness: number): string {
//     var r = 255-brightness;
//     var n = 0|((Math.random() * r) + brightness);
//     var s = n.toString(16);
//     return (s.length==1) ? '0'+s : s;
//   }
//   return '#' + randomChannel(brightness) + randomChannel(brightness) + randomChannel(brightness);
// }

func RandomChannel(brightness uint8) string {
	r := 255 - brightness
	n := uint8(rand.Uint32())%r + brightness
	s := string(n)
	if len(s) == 1 {
		return "0" + s
	}
	return s
}

// RandomColor : Generate a random color
func RandomColor(brightness uint8) string {
	return "#" + RandomChannel(brightness) + RandomChannel(brightness) + RandomChannel(brightness)
}
