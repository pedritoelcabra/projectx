package gfx

type Graphics struct {
}

func NewGraphics() *Graphics {
	aGraphics := &Graphics{}
	LoadSprites()
	SetUpLpcSpritesOffsets()
	SetUpBasicTerrainOffsets()
	return aGraphics
}
