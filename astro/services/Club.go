package services

type Nation struct {
	NationID   uint16 `json:"nation_id"`
	NationName string `json:"nation_name"`
	NationCode string `json:"nation_code"`
}

type Region struct {
	RegionID   uint16 `json:"region_id"`
	RegionName string `json:"region_name"`
}

type Club struct {
	ClubID   uint16 `json:"club_id"`
	ClubName string `json:"club_name"`
}

func (n *Nation) GetFlagPath() string {
	return "https://flagsapi.com/" + n.NationCode + "/flat/64.png"
}
