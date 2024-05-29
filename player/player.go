package player

// Player struct defines the attributes of a player
type Player struct {
	Name     string
	Health   int
	Strength int
	Attack   int
}

// IsAlive method checks if the player is still alive
func (p *Player) IsAlive() bool {
	return p.Health > 0
}
