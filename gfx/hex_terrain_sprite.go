package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/world/defs"
	"image"
)

// spritesheet dimension = 672 * 736 // 21 * 23
// tile size 72*72

type HexTerrainTypes int

var hexTerrainTypeMap = make(map[HexTerrainTypes]image.Rectangle)
var terrainToHex = make(map[int]HexTerrainTypes)

const (
	HexTerrainWidth  = 7
	HexTerrainHeight = 15
	HexTileSize      = 72
)

func SetUpHexTerrainOffsets() {
	hexTerrainTypeMap = make(map[HexTerrainTypes]image.Rectangle)
	terrainToHex = make(map[int]HexTerrainTypes)
	for x := 0; x < HexTerrainWidth; x++ {
		for y := 0; y < HexTerrainHeight; y++ {
			key := HexTerrainTypes((y * HexTerrainWidth) + x)
			keyX := HexTileSize * x
			keyY := HexTileSize * y
			imageRect := image.Rect(keyX, keyY, keyX+HexTileSize, keyY+HexTileSize)
			hexTerrainTypeMap[key] = imageRect
		}
	}
	terrainToHex[defs.BasicWater] = 2
	terrainToHex[defs.BasicDesert] = 24
	terrainToHex[defs.BasicGrass] = 51
	terrainToHex[defs.BasicMountain] = 60
}

func DrawHexTerrain(x, y float64, terrain int, screen *Screen, op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(x, y)
	screen.DrawImage(GetSprite(HexTerrain1).SubImage(hexTerrainTypeMap[terrainToHex[terrain]]).(*ebiten.Image), op)
}
