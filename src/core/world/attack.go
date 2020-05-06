package world

type Attack struct {
	Damage   float64
	Attacker *Unit
	Defender *Unit
}

func NewAttack() *Attack {
	anAttack := &Attack{}
	return anAttack
}
