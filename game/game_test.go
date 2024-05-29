package game

import (
	"magical-arena/player"
	"math/rand"
	"testing"
)

func TestSimulateMatch(t *testing.T) {
	playerA := player.Player{Name: "Player A", Health: 50, Strength: 5, Attack: 10}
	playerB := player.Player{Name: "Player B", Health: 100, Strength: 10, Attack: 5}

	// Mock the dice rolls for a controlled test
	rand.Seed(1) // Set seed for reproducibility

	SimulateMatch(&playerA, &playerB)

	// Expected results based on the fixed seed
	if playerA.Health != 0 {
		t.Errorf("Expected Player A to be dead, but got health %d", playerA.Health)
	}
	if playerB.Health <= 0 {
		t.Errorf("Expected Player B to be alive, but got health %d", playerB.Health)
	}
}
