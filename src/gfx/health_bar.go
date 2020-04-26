package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	"log"
	"strconv"
)

const (
	HealthBarWidth     float64 = 32.0
	HealthBarHalfWidth float64 = 16.0
	HealthBarHeight    float64 = 4.0
)

var healthBarImg = &ebiten.Image{}

type HealthBarOwner interface {
	GetHealth() float64
	GetMaxHealth() float64
	GetX() float64
	GetY() float64
}

func HealthString(owner HealthBarOwner) string {
	return strconv.Itoa(int(owner.GetHealth())) + "/" + strconv.Itoa(int(owner.GetMaxHealth()))
}

func InitHealthBar() {
	healthBarImgPath := "healthbar.png"
	loadedHealthBarImg, _, err := ebitenutil.NewImageFromFile(GetAbsoluteGfxPath(healthBarImgPath), ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	healthBarImg = loadedHealthBarImg
}

func DrawHealthBar(owner HealthBarOwner, screen *Screen) {
	op := &ebiten.DrawImageOptions{}
	drawX := owner.GetX() - HealthBarHalfWidth
	drawY := owner.GetY() - HexTileSize
	op.GeoM = ebiten.TranslateGeo(drawX, drawY)
	width := (HealthBarWidth / owner.GetMaxHealth()) * owner.GetHealth()
	drawRect := image.Rect(0, 0, int(width), int(HealthBarHeight))
	screen.DrawImage(healthBarImg.SubImage(drawRect).(*ebiten.Image), op)
}
