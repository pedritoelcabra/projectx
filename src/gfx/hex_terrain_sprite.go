package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/core/world/utils"
	"image"
)

// spritesheet dimension = 672 * 736 // 21 * 23
// tile size 72*72

type HexTerrainTypes int

var hexTerrainTypeMap = make(map[HexTerrainTypes]image.Rectangle)
var terrainToHex = make(map[int]HexTerrainTypes)

const (
	HexTerrainWidth  = 7
	HexTerrainHeight = 16
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
	terrainToHex[utils.BasicWater] = 85
	terrainToHex[utils.BasicDeepWater] = 84
	terrainToHex[utils.BasicHills] = 38
	terrainToHex[utils.BasicGrass] = 51
	terrainToHex[utils.BasicMountain] = 41
	terrainToHex[utils.TundraWater] = 62
	terrainToHex[utils.TundraDeepWater] = 0
	terrainToHex[utils.TundraHills] = 98
	terrainToHex[utils.TundraGrass] = 62
	terrainToHex[utils.TundraMountain] = 96
	terrainToHex[utils.DesertWater] = 1
	terrainToHex[utils.DesertDeepWater] = 0
	terrainToHex[utils.DesertHills] = 4
	terrainToHex[utils.DesertGrass] = 18
	terrainToHex[utils.DesertMountain] = 15
	terrainToHex[utils.BorderN] = 105
	terrainToHex[utils.BorderNE] = 106
	terrainToHex[utils.BorderNW] = 107
	terrainToHex[utils.BorderS] = 108
	terrainToHex[utils.BorderSW] = 109
	terrainToHex[utils.BorderSE] = 110
	terrainToHex[utils.GreenOverlay] = 101
	terrainToHex[utils.RedOverlay] = 102
}

func DrawHexTerrain(x, y float64, terrain int, screen *Screen, op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(x, y)
	screen.DrawImage(GetImage(HexTerrain1).SubImage(hexTerrainTypeMap[terrainToHex[terrain]]).(*ebiten.Image), op)
}

func DrawHexTerrainToImage(x, y float64, terrain int, image *ebiten.Image, op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(x, y)
	image.DrawImage(GetImage(HexTerrain1).SubImage(hexTerrainTypeMap[terrainToHex[terrain]]).(*ebiten.Image), op)
}
