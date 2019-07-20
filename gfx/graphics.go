package gfx

type Graphics struct {
}

func NewGraphics() *Graphics {
	aGraphics := &Graphics{}
	LoadLPCSprites()
	SetUpLpcSpritesOffsets()
	return aGraphics
}
