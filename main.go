package main

import (
	"magical-arena/game"
	"magical-arena/player"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initialize players
	playerA := &player.Player{Name: "Player A", Health: 50, Strength: 5, Attack: 10}
	playerB := &player.Player{Name: "Player B", Health: 100, Strength: 10, Attack: 5}

	// Simulate match
	game.SimulateMatch(playerA, playerB)
}
