package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

var spriteMap = make(map[SpriteKey]*ebiten.Image)

type SpriteLoader struct {
	sprites map[SpriteKey]*ebiten.Image
}

func GetSprite(key SpriteKey) *ebiten.Image {
	return spriteMap[key]
}

func LoadSprites() {
	spriteMap = make(map[SpriteKey]*ebiten.Image)
	for key, path := range SpritePaths() {
		img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		spriteMap[key] = img
	}
}

func SpritePaths() map[SpriteKey]string {
	return map[SpriteKey]string{
		BodyMaleLight: "resources/Universal-LPC-spritesheet/body/male/light.png",
		BasicTerrain:  "resources/tiles/terrain.png",
	}
}
