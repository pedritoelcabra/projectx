package gfx

type spriteFacing int
type SpriteKey int

const (
	_ SpriteKey = iota
	HexTerrain1
)

const (
	FaceUp spriteFacing = iota
	FaceLeft
	FaceDown
	FaceRight
)

type Sprite interface {
	DrawSprite(*Screen, float64, float64)
	SetFacing(direction spriteFacing)
	QueueAttackAnimation(x, y float64, speed int)
}
