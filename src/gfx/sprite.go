package gfx

type spriteFacing int

const (
	FaceUp spriteFacing = iota
	FaceLeft
	FaceDown
	FaceRight
)

type Sprite interface {
	DrawSprite(*Screen, float64, float64)
	SetFacing(direction spriteFacing)
}
