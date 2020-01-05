package gfx

type Graphics struct {
}

func NewGraphics() *Graphics {
	aGraphics := &Graphics{}
	spriteKeyCount = 10000
	spriteToKeyMap = make(map[string]SpriteKey)
	LoadSprites()
	LoadGfxFolder("buildings")
	LoadGfxFolder("vegetation")
	SetUpLpcSpritesOffsets()
	SetUpHexTerrainOffsets()
	return aGraphics
}
