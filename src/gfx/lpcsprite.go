package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/core/world/utils"
	"image"
	"log"
	"strconv"
)

type lpcAnimation int
type lpcTween int

const (
	spriteWidth          = 64
	spriteWidthHalf      = 32
	spriteHeight         = 64
	spriteHeightHalf     = 32
	centerXAxis          = 32
	centerYAxis          = 64
	DefaultCollisionSize = 32
)

const (
	AttackAnimationTween lpcTween = iota
)

const (
	AttackAnimationForward  float64 = 3.0
	AttackAnimationBackward         = 2.0
)

const (
	Casting lpcAnimation = iota
	Thrusting
	Walking
	Slashing
	Shooting
	Dying
)

type Tween struct {
	TweenType lpcTween
	X         float64
	Y         float64
	Progress  int
	Speed     int
}

func NewAttackAnimation(x, y float64, speed int) *Tween {
	aTween := &Tween{}
	aTween.TweenType = AttackAnimationTween
	aTween.X = x
	aTween.Y = y
	aTween.Speed = speed
	aTween.Progress = 0
	return aTween
}

func (t *Tween) apply() (x, y float64, ended bool) {
	if t.TweenType == AttackAnimationTween {
		return t.ApplyAttackAnimation()
	}
	log.Fatal("Executing invalid tween: " + strconv.Itoa(int(t.TweenType)))
	return 0.0, 0.0, true
}

func (t *Tween) ApplyAttackAnimation() (x, y float64, ended bool) {
	t.Progress++
	ended = false
	x, y = 0.0, 0.0
	endFrame := t.Speed / 2
	goFrameEnd := t.Speed / 6
	if goFrameEnd == 0 || endFrame == 0 {
		return 0.0, 0.0, true
	}
	if t.Progress >= endFrame {
		return 0.0, 0.0, true
	}
	if t.Progress < goFrameEnd {
		progress := AttackAnimationForward * float64(t.Progress)
		x, y = utils.AdvanceAlongLine(0, 0, t.X, t.Y, progress)
	} else {
		progress := AttackAnimationBackward * float64(t.Progress-goFrameEnd)
		x, y = utils.AdvanceAlongLine(t.X, t.Y, 0, 0, progress)
	}
	return x, y, false
}

var animationLength = map[lpcAnimation]int{
	Casting:   7,
	Thrusting: 8,
	Walking:   9,
	Slashing:  6,
	Shooting:  13,
	Dying:     6,
}

var offsetMap = make(map[lpcAnimation]map[spriteFacing]map[int]image.Rectangle)

type LpcSprite struct {
	key       SpriteKey
	facing    spriteFacing
	animation lpcAnimation
	frame     int
	tween     *Tween
}

func SetUpLpcSpritesOffsets() {
	for i := Casting; i <= Dying; i++ {
		offsetMap[i] = make(map[spriteFacing]map[int]image.Rectangle)
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

func NewLpcSprite(key SpriteKey) *LpcSprite {
	aSprite := &LpcSprite{}
	aSprite.key = key
	aSprite.animation = Walking
	aSprite.facing = FaceDown
	aSprite.frame = 0
	return aSprite
}

func (s *LpcSprite) DrawSprite(screen *Screen, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	s.DrawSpriteWithOptions(screen, x, y, op)
}

func (s *LpcSprite) DrawSpriteWithOptions(screen *Screen, x, y float64, op *ebiten.DrawImageOptions) {
	tweenX, tweenY := 0.0, 0.0
	if s.tween != nil {
		ended := false
		tweenX, tweenY, ended = s.tween.apply()
		if ended {
			s.tween = nil
		}
	}
	op.GeoM = ebiten.TranslateGeo(x-centerXAxis+tweenX, y-centerYAxis+tweenY)
	screen.DrawImage(s.getFrame(), op)
}

func (s *LpcSprite) DrawSpriteSubImage(screen *Screen, x, y float64, percentage float64) {
	s.DrawSprite(screen, x, y)
}

func (s *LpcSprite) getFrame() *ebiten.Image {
	return GetImage(s.key).SubImage(s.getRect()).(*ebiten.Image)
}

func (s *LpcSprite) getRect() image.Rectangle {
	rect := offsetMap[s.animation][s.facing][s.frame]
	return rect
}

func (s *LpcSprite) SetFacing(direction spriteFacing) {
	s.facing = direction
}

func (s *LpcSprite) QueueAttackAnimation(x, y float64, speed int) {
	s.tween = NewAttackAnimation(x, y, speed)
}
