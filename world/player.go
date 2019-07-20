package world

import "github.com/hajimehoshi/ebiten"

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
	return aPlayer
}

func (p *Player) DrawSprite(screen *ebiten.Image) {
	p.unit.DrawSprite(screen)
}

func (p *Player) SetPosition(x, y float64) {
	p.unit.SetPosition(x, y)
}

func (p *Player) Update(tick int) {
	p.unit.Update(tick)
}

func (p *Player) MovingUp(value bool) {
	p.movingUp = value
}

func (p *Player) MovingLeft(value bool) {
	p.movingLeft = value
}

func (p *Player) MovingDown(value bool) {
	p.movingDown = value
}

func (p *Player) MovingRight(value bool) {
	p.movingRight = value
}
