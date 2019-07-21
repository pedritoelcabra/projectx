package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gfx"
)

type Unit struct {
	sprite gfx.Sprite
	x      float64
	y      float64
	destX  float64
	destY  float64
	moving bool
	speed  float64
}

func (u *Unit) DrawSprite(screen *ebiten.Image) {
	u.sprite.DrawSprite(screen, u.x, u.y)
}

func (u *Unit) SetPosition(x, y float64) {
	u.x = x
	u.y = y
	u.destX = x
	u.destY = y
	u.CheckIfMoving()
}

func (u *Unit) SetDestination(x, y float64) {
	u.destX = x
	u.destY = y
	u.CheckIfMoving()
}

func NewUnit() *Unit {
	aUnit := &Unit{}
	aUnit.sprite = gfx.NewLpcSprite(gfx.BodyMaleLight)
	aUnit.speed = 100
	return aUnit
}

func (u *Unit) Update(tick int) {
	if u.moving {
		newX, newY := AdvanceAlongLine(u.x, u.y, u.destX, u.destY, u.speed)
		u.SetPosition(newX, newY)
	}
}

func (u *Unit) CheckIfMoving() {
	if u.destY != u.y || u.destX != u.x {
		u.moving = true
		return
	}
	u.moving = false
}

func (u *Unit) GetPos() (x, y float64) {
	return u.x, u.y
}
