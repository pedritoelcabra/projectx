package gfx

type Sprite interface {
	DrawSprite(*Screen, float64, float64)
}
