package rules

import "testing"

func TestAdvanceDoesNotModifyOriginal(t *testing.T) {
	g := NewGame(2, DefaultConfig)

	initialState := g.State

	g.Advance()

	if initialState != g.State {
		t.Errorf("Expected state %v but got %v", initialState, g.State)
	}
}

func TestNewGameFillsWithVillagers(t *testing.T) {
	numPlayers := 100
	g := NewGame(numPlayers, Config{Roles: []Role{RoleSeer, RoleTroublemaker, RoleRobber}})

	numWerewolves := 0
	numVillagers := 0
	numSeers := 0
	numTroublemakers := 0
	numRobbers := 0

	for _, p := range g.Players {
		switch p.Role {
		case RoleWerewolf:
			numWerewolves++

		case RoleVillager:
			numVillagers++

		case RoleSeer:
			numSeers++

		case RoleTroublemaker:
			numTroublemakers++

		case RoleRobber:
			numRobbers++

		default:
			t.Fatalf("Expected only default roles, but found %v", p.Role)
		}
	}

	if numWerewolves != 2 {
		t.Errorf("Expected 2 werewolves, but found %d", numWerewolves)
	}

	expectedVillagers := numPlayers - 5

	if numVillagers != expectedVillagers {
		t.Errorf("Expected %d villagers, but found %d", expectedVillagers, numVillagers)
	}

	if numSeers != 1 {
		t.Errorf("Expected 1 seer, but found %d", numSeers)
	}

	if numTroublemakers != 1 {
		t.Errorf("Expected 1 troublemaker, but found %d", numTroublemakers)
	}

	if numRobbers != 1 {
		t.Errorf("Expected 1 robber, but found %d", numRobbers)
	}
}
