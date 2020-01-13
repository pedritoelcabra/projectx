package world

import (
	"github.com/pedritoelcabra/projectx/src/gfx"
)

type PlayerDirection int

const (
	PLAYERUP PlayerDirection = iota
	PLAYERLEFT
	PLAYERDOWN
	PLAYERRIGHT
)

type Player struct {
	ClassName   string
	unit        *Unit
	MovingUp    bool
	MovingLeft  bool
	MovingDown  bool
	MovingRight bool
}

func NewPlayer() *Player {
	aPlayer := &Player{}
	aPlayer.Init()
	aPlayer.unit = NewUnit()
	aPlayer.unit.SetSpeed(10)
	return aPlayer
}

func (p *Player) Init() {
	p.unit = theWorld.GetUnit(0)
	p.ClassName = "Player"
}

func (p *Player) DrawSprite(screen *gfx.Screen) {
	p.unit.DrawSprite(screen)
}

func (p *Player) SetPosition(x, y float64) {
	p.unit.SetPosition(x, y)
}

func (p *Player) Update(tick int, grid *Grid) {
	p.UpdateDestination()
}

func (p *Player) UpdateDestination() {
	destX := p.unit.X
	destY := p.unit.Y
	displacement := float64(1000)
	if p.MovingUp && !p.MovingDown {
		destY -= displacement
	}
	if p.MovingDown && !p.MovingUp {
		destY += displacement
	}
	if p.MovingLeft && !p.MovingRight {
		destX -= displacement
	}
	if p.MovingRight && !p.MovingLeft {
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
		p.MovingUp = value
	case PLAYERLEFT:
		p.MovingLeft = value
	case PLAYERDOWN:
		p.MovingDown = value
	case PLAYERRIGHT:
		p.MovingRight = value
	}
}

func (p *Player) GetClassName() string {
	return p.ClassName
}

func (p *Player) Get(key int) int {
	return p.unit.Get(key)
}

func (p *Player) GetF(key int) float64 {
	return p.unit.GetF(key)
}

func (p *Player) Set(key, value int) {
	p.unit.Set(key, value)
}

func (p *Player) SetF(key int, value float64) {
	p.unit.SetF(key, value)
}
