package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

type SpriteLoader struct {
	sprites map[spriteKey]*ebiten.Image
}

func NewSpriteLoader() *SpriteLoader {
	aLoader := &SpriteLoader{}
	return aLoader
}

func (l *SpriteLoader) GetSprite(key spriteKey) *ebiten.Image {
	return l.sprites[key]
}

func (l *SpriteLoader) LoadLPCSprites() {
	l.sprites = make(map[spriteKey]*ebiten.Image)
	for key, path := range SpritePaths() {
		img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		l.sprites[key] = img
	}
}

func SpritePaths() map[spriteKey]string {
	return map[spriteKey]string{
		bodyMaleLight: "resources/body/male/bodyMaleLight.png",
	}
}
