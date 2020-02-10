package world

import (
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"github.com/pedritoelcabra/projectx/src/world/utils"
	"math"
)

func (u *Unit) SetPosition(x, y float64) {
	u.X = x
	u.Y = y
	u.CheckIfMoving()
}

func (u *Unit) SetDestination(x, y float64) {
	u.DestX = x
	u.DestY = y
	u.CheckIfMoving()
	if u.Moving {
		u.OrientateTowardsPoint(u.DestX, u.DestY)
	}
}

func (u *Unit) OrientateTowardsPoint(x, y float64) {
	if math.Abs(u.X-x)+1 > math.Abs(u.Y-y) {
		if u.X > x {
			u.Sprite.SetFacing(gfx.FaceLeft)
			return
		}
		u.Sprite.SetFacing(gfx.FaceRight)
		return
	}
	if u.Y > y {
		u.Sprite.SetFacing(gfx.FaceUp)
		return
	}
	u.Sprite.SetFacing(gfx.FaceDown)
}

func (u *Unit) CheckIfMoving() {
	if u.DestY != u.Y || u.DestX != u.X {
		u.Moving = true
		return
	}
	u.Moving = false
}

func (u *Unit) StopMovement() {
	u.DestY = u.X
	u.DestY = u.Y
	u.Moving = false
}

func (u *Unit) GetPos() (x, y float64) {
	return u.X, u.Y
}

func (u *Unit) GetTileCoord() tiling.Coord {
	return tiling.PixelFToTileC(u.GetPos())
}

func (u *Unit) DistanceToUnit(t *Unit) int {
	return u.DistanceToPoint(t.GetPos())
}

func (u *Unit) DistanceToPoint(x, y float64) int {
	return tiling.NewCoordF(u.GetPos()).ChebyshevDist(tiling.NewCoordF(x, y))
}

func (u *Unit) DistanceWithinVision(distance int) bool {
	return distance < int(u.GetF(Vision))
}

func (u *Unit) DistanceWithinAttackRange(distance int) bool {
	attackRange := int(u.GetF(AttackRange))
	return distance <= attackRange
}

func (u *Unit) CollidesWith(x, y float64) bool {
	return utils.CalculateDistance(u.X, u.Y, x, y) < u.GetF(Size)
}
