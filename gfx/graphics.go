package gfx

type Graphics struct {
}

func NewGraphics() *Graphics {
	aGraphics := &Graphics{}
	LoadSprites()
	SetUpLpcSpritesOffsets()
	SetUpHexTerrainOffsets()
	return aGraphics
}
