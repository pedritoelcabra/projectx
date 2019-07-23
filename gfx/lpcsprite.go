package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"image"
)

type lpcAnimation int
type lpcFacing int

const (
	spriteWidth      = 64
	spriteWidthHalf  = 32
	spriteHeight     = 64
	spriteHeightHalf = 32
)

const (
	FaceUp lpcFacing = iota
	FaceLeft
	FaceDown
	FaceRight
)

const (
	Casting lpcAnimation = iota
	Thrusting
	Walking
	Slashing
	Shooting
	Dying
)

var animationLength = map[lpcAnimation]int{
	Casting:   7,
	Thrusting: 8,
	Walking:   9,
	Slashing:  6,
	Shooting:  13,
	Dying:     6,
}

var offsetMap = make(map[lpcAnimation]map[lpcFacing]map[int]image.Rectangle)

type LpcSprite struct {
	key       spriteKey
	facing    lpcFacing
	animation lpcAnimation
	frame     int
}

func SetUpLpcSpritesOffsets() {
	for i := Casting; i <= Dying; i++ {
		offsetMap[i] = make(map[lpcFacing]map[int]image.Rectangle)
		for k := FaceUp; k <= FaceRight; k++ {
			offsetMap[i][k] = make(map[int]image.Rectangle)
			for f := 0; f < animationLength[Shooting]; f++ {
				x := f * 64
				y := int(i*4*64) + int(k*64)
				rect := image.Rect(x, y, x+spriteWidth, y+spriteHeight)
				offsetMap[i][k][f] = rect
			}
		}
	}
}

func NewLpcSprite(key spriteKey) *LpcSprite {
	aSprite := &LpcSprite{}
	aSprite.key = key
	aSprite.animation = Walking
	aSprite.facing = FaceDown
	aSprite.frame = 0
	return aSprite
}

func (s *LpcSprite) DrawSprite(screen *Screen, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = ebiten.TranslateGeo(x-spriteWidthHalf, y-spriteHeightHalf)
	screen.DrawImage(s.getFrame(), op)
}

func (s *LpcSprite) getFrame() *ebiten.Image {
	return GetSprite(s.key).SubImage(s.getRect()).(*ebiten.Image)
}

func (s *LpcSprite) getRect() image.Rectangle {
	rect := offsetMap[s.animation][s.facing][s.frame]
	return rect
}
