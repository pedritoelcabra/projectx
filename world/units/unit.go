package units

import (
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/world/movement"
	"math"
)

type Unit struct {
	Sprite     gfx.Sprite `json:"-"`
	SpriteName gfx.SpriteKey
	X          float64
	Y          float64
	DestX      float64
	DestY      float64
	Moving     bool
	Speed      float64
}

func (u *Unit) DrawSprite(screen *gfx.Screen) {
	u.Sprite.DrawSprite(screen, u.X, u.Y)
}

func (u *Unit) SetPosition(x, y float64) {
	u.X = x
	u.Y = y
	u.DestX = x
	u.DestY = y
	u.CheckIfMoving()
}

func (u *Unit) SetDestination(x, y float64) {
	u.DestX = x
	u.DestY = y
	u.CheckIfMoving()
	u.CheckOrientation()
}

func (u *Unit) CheckOrientation() {
	if !u.Moving {
		return
	}
	if math.Abs(u.X-u.DestX)+1 > math.Abs(u.Y-u.DestY) {
		if u.X > u.DestX {
			u.Sprite.SetFacing(gfx.FaceLeft)
			return
		}
		u.Sprite.SetFacing(gfx.FaceRight)
		return
	}
	if u.Y > u.DestY {
		u.Sprite.SetFacing(gfx.FaceUp)
		return
	}
	u.Sprite.SetFacing(gfx.FaceDown)
}

func (u *Unit) SetSpeed(speed float64) {
	u.Speed = speed
}

func NewUnit() *Unit {
	aUnit := &Unit{}
	aUnit.SpriteName = gfx.BodyMaleLight
	aUnit.InitObjects()
	aUnit.Speed = 100
	return aUnit
}

func (u *Unit) InitObjects() {
	u.Sprite = gfx.NewLpcSprite(u.SpriteName)
}

func (u *Unit) Update(tick int) {
	if u.Moving {
		newX, newY := movement.AdvanceAlongLine(u.X, u.Y, u.DestX, u.DestY, u.Speed)
		u.SetPosition(newX, newY)
	}
}

func (u *Unit) CheckIfMoving() {
	if u.DestY != u.Y || u.DestX != u.X {
		u.Moving = true
		return
	}
	u.Moving = false
}

func (u *Unit) GetPos() (x, y float64) {
	return u.X, u.Y
}
