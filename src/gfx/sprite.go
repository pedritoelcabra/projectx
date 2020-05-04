package gfx

import "github.com/hajimehoshi/ebiten"

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
	DrawSpriteWithOptions(screen *Screen, x, y float64, op *ebiten.DrawImageOptions)
	DrawSpriteSubImage(screen *Screen, x, y float64, percentage float64)
	SetFacing(direction spriteFacing)
	QueueAttackAnimation(x, y float64, speed int)
}
