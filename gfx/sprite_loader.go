package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

var spriteMap = make(map[spriteKey]*ebiten.Image)

type SpriteLoader struct {
	sprites map[spriteKey]*ebiten.Image
}

func GetSprite(key spriteKey) *ebiten.Image {
	return spriteMap[key]
}

func LoadSprites() {
	spriteMap = make(map[spriteKey]*ebiten.Image)
	for key, path := range SpritePaths() {
		img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		spriteMap[key] = img
	}
}

func SpritePaths() map[spriteKey]string {
	return map[spriteKey]string{
		BodyMaleLight: "resources/Universal-LPC-spritesheet/body/male/light.png",
		BasicTerrain:  "resources/tiles/terrain.png",
		HexTerrain1:   "resources/tiles/wesnoth1.png",
	}
}
