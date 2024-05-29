package dice

import "math/rand"

// RollDice simulates rolling a 6-sided dice
func RollDice() int {
	return rand.Intn(6) + 1
}
