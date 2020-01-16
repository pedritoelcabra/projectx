package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"log"
	"strconv"
)

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
	graphics.lpcSpriteKeys[strconv.Itoa(int(key))] = key
	return key
}

func AddImageToKey(image *ebiten.Image, key SpriteKey) {
	graphics.images[key] = image
}

func GetImage(key SpriteKey) *ebiten.Image {
	if _, ok := graphics.images[key]; !ok {
		log.Fatal("Could not find Sprite by key: " + strconv.Itoa(int(key)))
	}
	return graphics.images[key]
}

func GetLpcKey(name string) SpriteKey {
	return graphics.lpcSpriteKeys[name]
}

func GetLpcComposite(composite []SpriteKey) SpriteKey {
	lpcKey := GetLpcKey(GetCompositeKey(composite))
	if lpcKey == 0 {
		return BuildLpcComposite(composite)
	}
	return lpcKey
}

func GetCompositeKey(composite []SpriteKey) string {
	compositeKey := ""
	for i := 0; i < len(composite); i++ {
		compositeKey += strconv.Itoa(int(composite[i]))
	}
	return compositeKey
}

func BuildLpcComposite(composite []SpriteKey) SpriteKey {
	if len(composite) == 0 {
		return 0
	}
	firstImage := GetImage(composite[0])
	width, height := firstImage.Size()
	baseImage, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	for i := 0; i < len(composite); i++ {
		op := &ebiten.DrawImageOptions{}
		baseImage.DrawImage(GetImage(composite[i]), op)
	}
	return AddLpcImage(baseImage, GetCompositeKey(composite))
}

func LpcCompositeSlotOrder() []string {
	return []string{
		"Body",
		"Torso",
		"Legs",
	}
}
