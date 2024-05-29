package player

import "testing"

func TestPlayerIsAlive(t *testing.T) {
	player := Player{Name: "Test Player", Health: 10, Strength: 5, Attack: 5}
	if !player.IsAlive() {
		t.Errorf("Expected player to be alive, but got false")
	}
	player.Health = 0
	if player.IsAlive() {
		t.Errorf("Expected player to be dead, but got true")
	}
}
