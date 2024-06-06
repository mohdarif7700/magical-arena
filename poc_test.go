package main

import (
	"math/rand"
	"testing"
)

func TestPlayerIsAlive(t *testing.T) {
	player := Player{name: "Test Player", health: 10, strength: 5, attack: 5}
	if !player.IsAlive() {
		t.Errorf("Expected player to be alive, but got false")
	}
	player.health = 0
	if player.IsAlive() {
		t.Errorf("Expected player to be dead, but got true")
	}
}

func TestRollDice(t *testing.T) {
	for i := 0; i < 100; i++ {
		roll := RollDice()
		if roll < 1 || roll > 6 {
			t.Errorf("Expected dice roll to be between 1 and 6, but got %d", roll)
		}
	}
}

func TestAttack(t *testing.T) {
	attacker := Player{name: "Attacker", health: 100, strength: 10, attack: 5}
	defender := Player{name: "Defender", health: 100, strength: 5, attack: 3}

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
	if defender.health != 100-expectedDamage {
		t.Errorf("Expected defender health %d, but got %d", 100-expectedDamage, defender.health)
	}
}

func TestSimulateMatch(t *testing.T) {
	playerA := Player{name: "Player A", health: 50, strength: 5, attack: 10}
	playerB := Player{name: "Player B", health: 100, strength: 10, attack: 5}

	// Mock the dice rolls for a controlled test
	rand.Seed(1) // Set seed for reproducibility

	SimulateMatch(&playerA, &playerB)

	// Expected results based on the fixed seed
	if playerA.health != 0 {
		t.Errorf("Expected Player A to be dead, but got health %d", playerA.health)
	}
	if playerB.health <= 0 {
		t.Errorf("Expected Player B to be alive, but got health %d", playerB.health)
	}
}
