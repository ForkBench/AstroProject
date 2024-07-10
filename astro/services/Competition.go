package services

import "math"

// Competition : Competition details
type Competition struct {
	CompetitionID             uint8 // 255 competitions max
	CompetitionName           string
	CompetitionCategory       Category
	CompetitionWeapon         Weapon
	CompetitionState          State
	CompetitionMaxStageNumber uint8
	CompetitionStages         []Stage
	CompetitionPlayers        []Player
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
	c.CompetitionStages = []Stage{}
	c.CompetitionPlayers = []Player{}

	return c
}

func (c *Competition) AddStage(stage Stage) bool {
	if c.CompetitionState != IDLE {
		return false
	}

	if len(c.CompetitionStages) >= int(c.CompetitionMaxStageNumber) {
		return false
	}

	c.CompetitionStages = append(c.CompetitionStages, stage)

	return true
}

func (c *Competition) StagePosition(stage Stage) uint16 {
	for i, competitionStage := range c.CompetitionStages {
		if competitionStage == stage {
			return uint16(i)
		}
	}

	return math.MaxInt16
}

func (c *Competition) RemoveStage(stage Stage) bool {
	if c.CompetitionState != IDLE {
		return false
	}

	if c.StagePosition(stage) == math.MaxInt16 {
		return false
	}

	c.CompetitionStages = append(c.CompetitionStages[:c.StagePosition(stage)], c.CompetitionStages[c.StagePosition(stage)+1:]...)

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

func (c *Competition) AddPlayer(player Player) bool {
	if c.CompetitionState != REGISTERING {
		return false
	}

	c.CompetitionPlayers = append(c.CompetitionPlayers, player)

	return true
}

func (c *Competition) PlayerPosition(player Player) uint16 {
	for i, competitionPlayer := range c.CompetitionPlayers {
		if competitionPlayer == player {
			return uint16(i)
		}
	}

	return math.MaxInt16
}

func (c *Competition) RemovePlayer(player Player) bool {
	if c.CompetitionState != REGISTERING {
		return false
	}

	pos := c.PlayerPosition(player)

	if pos == math.MaxInt16 {
		return false
	}

	c.CompetitionPlayers = append(c.CompetitionPlayers[:pos], c.CompetitionPlayers[pos+1:]...)

	return true
}

func (c *Competition) AddPlayerToStage(player Player, stage Stage) bool {
	return stage.AddPlayer(player)
}

func (c *Competition) RemovePlayerFromStage(player Player, stage Stage) bool {
	return stage.RemovePlayer(player)
}
