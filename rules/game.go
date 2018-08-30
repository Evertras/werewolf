package rules

import (
	"math/rand"
)

type State int

const (
	StateSetup State = iota
	StateNight
	StateDay
	StateReveal
)

type Game struct {
	Players []Player `json:"players"`
	State   State    `json:"state"`
	Roles   []Role   `json:"roles"`
}

type Config struct {
	Roles []Role `json:"roles"`
}

var DefaultRoles = []Role{RoleSeer, RoleTroublemaker, RoleRobber}

var DefaultConfig = Config{
	Roles: DefaultRoles,
}

func NewGame(numPlayers int, config Config) Game {
	if config.Roles == nil {
		config.Roles = DefaultRoles
	}

	g := Game{
		Players: make([]Player, numPlayers),
		State:   StateSetup,
		Roles:   config.Roles,
	}

	selectableRoles := append([]Role{RoleWerewolf, RoleWerewolf}, config.Roles...)

	for len(selectableRoles) < numPlayers {
		selectableRoles = append(selectableRoles, RoleVillager)
	}

	for i := range g.Players {
		numSelectable := len(selectableRoles)
		r := rand.Intn(numSelectable)
		selectedRole := selectableRoles[r]
		selectableRoles[r] = selectableRoles[numSelectable-1]
		selectableRoles = selectableRoles[:numSelectable-1]

		g.Players[i].Role = selectedRole
	}

	return g
}

func (g Game) Advance() Game {
	switch g.State {
	case StateSetup:
		g.State = StateNight

	case StateNight:
		g.State = StateDay

	case StateDay:
		g.State = StateReveal
	}

	return g
}
