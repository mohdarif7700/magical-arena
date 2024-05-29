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

func TestAttack(t *testing.T) {
	attacker := player.Player{Name: "Attacker", Health: 100, Strength: 10, Attack: 5}
	defender := player.Player{Name: "Defender", Health: 100, Strength: 5, Attack: 3}

	// Mock the dice rolls for a controlled test
	rand.Seed(1) // Set seed for reproducibility

	attackRoll, defendRoll, damage := Attack(&attacker, &defender)

	expectedAttackRoll := 1
	expectedDefendRoll := 6
	expectedDamage := 5*expectedAttackRoll - 5*expectedDefendRoll // attack damage - defense

	if attackRoll != expectedAttackRoll {
		t.Errorf("Expected attack roll %d, but got %d", expectedAttackRoll, attackRoll)
	}
	if defendRoll != expectedDefendRoll {
		t.Errorf("Expected defend roll %d, but got %d", expectedDefendRoll, defendRoll)
	}
	if damage != expectedDamage {
		t.Errorf("Expected damage %d, but got %d", expectedDamage, damage)
	}
	if defender.Health != 100-expectedDamage {
		t.Errorf("Expected defender health %d, but got %d", 100-expectedDamage, defender.Health)
	}
}
