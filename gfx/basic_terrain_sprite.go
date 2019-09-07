package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"image"
)

// spritesheet dimension = 672 * 736 // 21 * 23

type BasicTerrainTypes int

var terrainTypeMap = make(map[BasicTerrainTypes]image.Rectangle)

const (
	BasicTerrainWidth  = 21
	BasicTerrainHeight = 23
	TileSize           = 32
)

func SetUpBasicTerrainOffsets() {
	terrainTypeMap = make(map[BasicTerrainTypes]image.Rectangle)
	for x := 0; x < BasicTerrainWidth; x++ {
		for y := 0; y < BasicTerrainHeight; y++ {
			key := BasicTerrainTypes((y * BasicTerrainWidth) + x)
			keyX := TileSize * x
			keyY := TileSize * y
			imageRect := image.Rect(keyX, keyY, keyX+TileSize, keyY+TileSize)
			terrainTypeMap[key] = imageRect
		}
	}
}

func DrawBasicTerrain(x, y float64, terrainType BasicTerrainTypes, screen *Screen) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = ebiten.TranslateGeo(x, y)
	screen.DrawImage(GetSprite(BasicTerrain).SubImage(terrainTypeMap[terrainType]).(*ebiten.Image), op)
}
