package gfx

type Graphics struct {
	SpriteLoader *SpriteLoader
}

func NewGraphics() *Graphics {
	aGraphics := &Graphics{}
	aGraphics.SpriteLoader = NewSpriteLoader()
	aGraphics.SpriteLoader.LoadLPCSprites()
	return aGraphics
}
