package game

import (
	"fmt"
	"magical-arena/dice"
	"magical-arena/player"
)

// SimulateMatch manages the turn-based logic for the match
func SimulateMatch(player1, player2 *player.Player) {
	// Lower health attacks first
	var attacker, defender *player.Player
	if player1.Health < player2.Health {
		attacker, defender = player1, player2
	} else {
		attacker, defender = player2, player1
	}

	round := 1
	for player1.IsAlive() && player2.IsAlive() {
		fmt.Printf("Round %d:\n", round)
		fmt.Printf("%s attacks %s\n", attacker.Name, defender.Name)
		attackRoll, defendRoll, damage := Attack(attacker, defender)
		fmt.Printf("Attack roll: %d, Defense roll: %d\n", attackRoll, defendRoll)
		fmt.Printf("Damage dealt: %d\n", damage)
		fmt.Printf("%s health: %d\n\n", defender.Name, defender.Health)

		if !defender.IsAlive() {
			fmt.Printf("%s has been defeated. %s wins!\n", defender.Name, attacker.Name)
			break
		}

		// Swap attacker and defender for next round
		attacker, defender = defender, attacker
		round++
	}
}

// Attack simulates an attack by the attacker on the defender
func Attack(attacker, defender *player.Player) (int, int, int) {
	attackRoll := dice.RollDice()
	defendRoll := dice.RollDice()

	attackDamage := attacker.Attack * attackRoll
	defendStrength := defender.Strength * defendRoll

	netDamage := attackDamage - defendStrength
	if netDamage < 0 {
		netDamage = 0
	}

	defender.Health -= netDamage
	if defender.Health < 0 {
		defender.Health = 0
	}

	return attackRoll, defendRoll, netDamage
}
