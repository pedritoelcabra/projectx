package world

import (
	"github.com/pedritoelcabra/projectx/gfx"
)

type PlayerDirection int

const (
	PLAYERUP PlayerDirection = iota
	PLAYERLEFT
	PLAYERDOWN
	PLAYERRIGHT
)

type Player struct {
	unit        *Unit
	movingUp    bool
	movingLeft  bool
	movingDown  bool
	movingRight bool
}

func NewPlayer() *Player {
	aPlayer := &Player{}
	aPlayer.unit = NewUnit()
	aPlayer.unit.SetSpeed(5)
	return aPlayer
}

func (p *Player) DrawSprite(screen *gfx.Screen) {
	p.unit.DrawSprite(screen)
}

func (p *Player) SetPosition(x, y float64) {
	p.unit.SetPosition(x, y)
}

func (p *Player) Update(tick int) {
	p.UpdateDestination()
	p.unit.Update(tick)
}

func (p *Player) UpdateDestination() {
	destX := p.unit.x
	destY := p.unit.y
	displacement := float64(1000)
	if p.movingUp && !p.movingDown {
		destY -= displacement
	}
	if p.movingDown && !p.movingUp {
		destY += displacement
	}
	if p.movingLeft && !p.movingRight {
		destX -= displacement
	}
	if p.movingRight && !p.movingLeft {
		destX += displacement
	}
	p.unit.SetDestination(destX, destY)
}

func (p *Player) GetPos() (x, y float64) {
	return p.unit.GetPos()
}

func (p *Player) SetMovement(direction PlayerDirection, value bool) {
	switch direction {
	case PLAYERUP:
		p.movingUp = value
	case PLAYERLEFT:
		p.movingLeft = value
	case PLAYERDOWN:
		p.movingDown = value
	case PLAYERRIGHT:
		p.movingRight = value
	}
}
