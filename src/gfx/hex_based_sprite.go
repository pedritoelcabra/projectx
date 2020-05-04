package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var spriteToKeyMap = make(map[string]SpriteKey)
var spriteKeyCount = 10000

func LoadGfxFolder(folderName string) {
	directoryPath := GetAbsoluteGfxPath(folderName)
	walkErr := filepath.Walk(directoryPath, func(path string, info os.FileInfo, walkErr error) error {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if filepath.Ext(path) != ".png" {
			return nil
		}
		_, fileName := filepath.Split(path)
		fileName = strings.Replace(fileName, filepath.Ext(path), "", -1)
		img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		key := AddImage(img)
		spriteToKeyMap[fileName] = key
		return walkErr
	})
	if walkErr != nil {
		log.Fatal(walkErr)
	}
}

func GetSpriteKey(name string) SpriteKey {
	return spriteToKeyMap[name]
}

type HexSprite struct {
	key SpriteKey
}

func NewHexSprite(key SpriteKey) *HexSprite {
	aSprite := &HexSprite{}
	aSprite.key = key
	return aSprite
}

func (h *HexSprite) DrawSprite(screen *Screen, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	h.DrawSpriteWithOptions(screen, x, y, op)
}

func (h *HexSprite) DrawSpriteWithOptions(screen *Screen, x, y float64, op *ebiten.DrawImageOptions) {
	op.GeoM = ebiten.TranslateGeo(x, y)
	screen.DrawImage(GetImage(h.key), op)
}

func (h *HexSprite) DrawSpriteSubImage(screen *Screen, x, y float64, percentage float64) {

	op := &ebiten.DrawImageOptions{}
	drawImage := GetImage(h.key)
	if percentage < 0 {
		percentage = 0.0
	}
	if percentage > 100 {
		percentage = 100.0
	}
	width, height := drawImage.Size()
	wantedHeight := (float64(height) / 100.0) * percentage
	heightOffset := height - int(wantedHeight)
	if heightOffset < 0 {
		heightOffset = 0
	}
	if heightOffset > height {
		heightOffset = height
	}
	rect := image.Rect(0, heightOffset, width, height)
	op.GeoM = ebiten.TranslateGeo(x, y+float64(heightOffset))
	screen.DrawImage(drawImage.SubImage(rect).(*ebiten.Image), op)
}

func (h *HexSprite) SetFacing(direction spriteFacing) {

}

func (h *HexSprite) QueueAttackAnimation(x, y float64, speed int) {

}
