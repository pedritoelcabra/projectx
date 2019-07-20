package gfx

type Graphics struct {
}

func NewGraphics() *Graphics {
	aGraphics := &Graphics{}
	LoadLPCSprites()
	return aGraphics
}
