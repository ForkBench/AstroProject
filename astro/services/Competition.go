package services

// Competition : Competition details
type Competition struct {
	CompetitionID             uint8 // 255 competitions max
	CompetitionName           string
	CompetitionCategory       Category
	CompetitionWeapon         Weapon
	CompetitionState          State
	CompetitionMaxStageNumber uint8
	CompetitionStages         map[uint8]*Stage
	CompetitionPlayers        map[uint16]*Player
}

func (c *Competition) String() string {
	return "Competition : " + c.CompetitionName + " - " + c.CompetitionCategory.String() + " - " + c.CompetitionWeapon.String() + " - " + c.CompetitionState.String()
}

func CreateCompetition(competitionID uint8, competitionName string, competitionCategory Category, competitionWeapon Weapon, competitionMaxStageNumber uint8) Competition {
	var c Competition

	c.CompetitionID = competitionID
	c.CompetitionName = competitionName
	c.CompetitionCategory = competitionCategory
	c.CompetitionWeapon = competitionWeapon
	c.CompetitionState = REGISTERING
	c.CompetitionMaxStageNumber = competitionMaxStageNumber
	c.CompetitionStages = map[uint8]*Stage{}
	c.CompetitionPlayers = map[uint16]*Player{}

	c.InitCompetition()

	return c
}

func (c *Competition) InitCompetition() bool {
	if c.CompetitionState != REGISTERING {
		return false
	}

	// Add new stage
	var poolStage = CreateSeedingStage(0, 300, 300)
	// TODO: Max player number should be dynamic

	c.AddStage(poolStage)

	return true
}

func (c *Competition) StartCompetition() bool {
	if c.CompetitionState != REGISTERING {
		return false
	}

	c.CompetitionState = STARTED

	return true
}

func (c *Competition) FinishCompetition() bool {
	if c.CompetitionState != STARTED {
		return false
	}

	// TODO: Implement competition results

	c.CompetitionState = FINISHED

	return true
}

func (c *Competition) AddPlayer(player *Player) bool {
	if c.CompetitionState != REGISTERING {
		return false
	}

	c.CompetitionPlayers[player.PlayerID] = player

	return true
}

func (c *Competition) RemovePlayer(player *Player) bool {
	if c.CompetitionState != REGISTERING {
		return false
	}

	if _, ok := c.CompetitionPlayers[player.PlayerID]; !ok {
		return false
	}

	delete(c.CompetitionPlayers, player.PlayerID)

	return true
}

func (c *Competition) AddPlayerToStage(player Player, stage Stage) bool {
	return stage.AddPlayer(&player)
}

func (c *Competition) RemovePlayerFromStage(player Player, stage Stage) bool {
	return stage.RemovePlayer(&player)
}

func (c *Competition) UpdatePlayer(player *Player) bool {
	// Check if player is in competition
	if _, ok := c.CompetitionPlayers[player.PlayerID]; !ok {
		return false
	}

	c.CompetitionPlayers[player.PlayerID] = player

	return true
}

func (c *Competition) AddStage(stage Stage) bool {
	if c.CompetitionState != REGISTERING {
		return false
	}

	c.CompetitionStages[stage.GetID()] = &stage

	return true
}
