package gfx

type Graphics struct {
}

func NewGraphics() *Graphics {
	aGraphics := &Graphics{}
	LoadSprites()
	LoadGfxFolder("buildings")
	SetUpLpcSpritesOffsets()
	SetUpHexTerrainOffsets()
	return aGraphics
}
