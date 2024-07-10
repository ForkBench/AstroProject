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

func (r *Nation) Load() []Nation {
	return LoadData[Nation]("../resources/nations.json")
}

func (r *Region) Load() []Region {
	return LoadData[Region]("../resources/regions.json")
}

func (c *Club) Load() []Club {
	return LoadData[Club]("../resources/clubs.json")
}
