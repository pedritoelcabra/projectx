package gfx

type Graphics struct {
}

func NewGraphics() *Graphics {
	aGraphics := &Graphics{}
	LoadSprites()
	SetUpLpcSpritesOffsets()
	SetUpBasicTerrainOffsets()
	SetUpHexTerrainOffsets()
	return aGraphics
}
