package gfx

import "github.com/hajimehoshi/ebiten"

type Graphics struct {
	images        map[SpriteKey]*ebiten.Image
	imageKeyCount SpriteKey
}

var graphics = &Graphics{}

func NewGraphics() *Graphics {
	graphics = &Graphics{}
	graphics.images = make(map[SpriteKey]*ebiten.Image)
	graphics.imageKeyCount = 10000
	spriteToKeyMap = make(map[string]SpriteKey)
	LoadSprites()
	LoadGfxFolder("buildings")
	LoadGfxFolder("vegetation")
	SetUpLpcSpritesOffsets()
	SetUpHexTerrainOffsets()
	return graphics
}

func AddImage(image *ebiten.Image) SpriteKey {
	graphics.imageKeyCount++
	graphics.images[graphics.imageKeyCount] = image
	return graphics.imageKeyCount
}

func AddImageToKey(image *ebiten.Image, key SpriteKey) {
	graphics.images[key] = image
}

func GetImage(key SpriteKey) *ebiten.Image {
	return graphics.images[key]
}

func GetLpcKey(name string) SpriteKey {
	return lpcSprites[name]
}

func GetLpcComposite(composite []SpriteKey) SpriteKey {
	return composite[0]
}
