package gfx

import "github.com/hajimehoshi/ebiten"

type Graphics struct {
	images        map[SpriteKey]*ebiten.Image
	lpcSpriteKeys map[string]SpriteKey
	imageKeyCount SpriteKey
}

var graphics = &Graphics{}

func NewGraphics() *Graphics {
	graphics = &Graphics{}
	graphics.images = make(map[SpriteKey]*ebiten.Image)
	graphics.lpcSpriteKeys = make(map[string]SpriteKey)
	graphics.imageKeyCount = 10000
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

func AddLpcImage(image *ebiten.Image, name string) SpriteKey {
	key := AddImage(image)
	graphics.lpcSpriteKeys[name] = key
	return key
}

func AddImageToKey(image *ebiten.Image, key SpriteKey) {
	graphics.images[key] = image
}

func GetImage(key SpriteKey) *ebiten.Image {
	return graphics.images[key]
}

func GetLpcKey(name string) SpriteKey {
	return graphics.lpcSpriteKeys[name]
}

func GetLpcComposite(composite []SpriteKey) SpriteKey {
	return composite[0]
}
