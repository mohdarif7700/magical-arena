package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Player struct defines the attributes of a player
type Player struct {
	name     string
	health   int
	strength int
	attack   int
}

// IsAlive method checks if the player is still alive
func (p *Player) IsAlive() bool {
	return p.health > 0
}

// RollDice simulates rolling a 6-sided dice
func RollDice() int {
	return rand.Intn(6) + 1
}

// Attack simulates an attack by the attacker on the defender
func Attack(attacker, defender *Player) (int, int, int) {
	attackRoll := RollDice()
	defendRoll := RollDice()

	attackDamage := attacker.attack * attackRoll
	defendStrength := defender.strength * defendRoll

	netDamage := attackDamage - defendStrength
	if netDamage < 0 {
		netDamage = 0
	}

	defender.health -= netDamage
	if defender.health < 0 {
		defender.health = 0
	}

	return attackRoll, defendRoll, netDamage
}

// SimulateMatch manages the turn-based logic for the match
func SimulateMatch(player1, player2 *Player) {
	// Lower health attacks first
	var attacker, defender *Player
	if player1.health < player2.health {
		attacker, defender = player1, player2
	} else {
		attacker, defender = player2, player1
	}

	round := 1
	for player1.IsAlive() && player2.IsAlive() {
		fmt.Printf("Round %d:\n", round)
		fmt.Printf("%s attacks %s\n", attacker.name, defender.name)
		attackRoll, defendRoll, damage := Attack(attacker, defender)
		fmt.Printf("Attack roll: %d, Defense roll: %d\n", attackRoll, defendRoll)
		fmt.Printf("Damage dealt: %d\n", damage)
		fmt.Printf("%s health: %d\n\n", defender.name, defender.health)

		if !defender.IsAlive() {
			fmt.Printf("%s has been defeated. %s wins!\n", defender.name, attacker.name)
			break
		}

		// Swap attacker and defender for next round
		attacker, defender = defender, attacker
		round++
	}
}

func abc() {
	rand.Seed(time.Now().UnixNano())

	// Initialize players
	playerA := &Player{name: "Player A", health: 50, strength: 5, attack: 10}
	playerB := &Player{name: "Player B", health: 100, strength: 10, attack: 5}

	// Simulate match
	SimulateMatch(playerA, playerB)
}
