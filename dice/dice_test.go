package dice

import "testing"

func TestRollDice(t *testing.T) {
	for i := 0; i < 100; i++ {
		roll := RollDice()
		if roll < 1 || roll > 6 {
			t.Errorf("Expected dice roll to be between 1 and 6, but got %d", roll)
		}
	}
}
