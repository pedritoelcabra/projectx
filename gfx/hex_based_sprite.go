package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var spriteToKeyMap = make(map[string]SpriteKey)
var spriteKeyCount = 10000

const gfxFolder = "resources/"

func LoadGfxFolder(folderName string) {
	spriteToKeyMap = make(map[string]SpriteKey)
	spriteKeyCount = 10000
	directoryPath, _ := filepath.Abs(gfxFolder + folderName)
	walkErr := filepath.Walk(directoryPath, func(path string, info os.FileInfo, walkErr error) error {
		spriteKeyCount++
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
		spriteToKeyMap[fileName] = SpriteKey(spriteKeyCount)
		img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		spriteMap[SpriteKey(spriteKeyCount)] = img
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
	op.GeoM = ebiten.TranslateGeo(x, y)
	screen.DrawImage(GetSprite(h.key), op)
}

func (h *HexSprite) SetFacing(direction spriteFacing) {

}
