package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
	"os"
	"path/filepath"
)

var spriteMap = make(map[SpriteKey]*ebiten.Image)
var lpcSprites = make(map[string]SpriteKey)

type SpriteLoader struct {
	sprites map[SpriteKey]*ebiten.Image
}

func GetSprite(key SpriteKey) *ebiten.Image {
	return spriteMap[key]
}

func LoadSprites() {
	spriteMap = make(map[SpriteKey]*ebiten.Image)
	for _, path := range SpritePaths() {
		img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		key := len(spriteMap)
		spriteMap[SpriteKey(key)] = img
	}
	lpcSprites = make(map[string]SpriteKey)
	for folderName, folderPath := range LPCSpriteFolders() {
		directoryPath, _ := filepath.Abs(folderPath)
		walkErr := filepath.Walk(directoryPath, func(path string, info os.FileInfo, walkErr error) error {
			pathStruct, _ := filepath.Split(path)
			if pathStruct != directoryPath+string(filepath.Separator) {
				return nil
			}
			if info.IsDir() {
				return nil
			}
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			if filepath.Ext(path) != ".png" {
				return nil
			}
			img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
			_, fileName := filepath.Split(path)
			key := SpriteKey(len(spriteMap))
			spriteMap[key] = img
			lpcSprites[(folderName + fileName)] = key
			return walkErr
		})
		if walkErr != nil {
			log.Fatal(walkErr)
		}
	}
}

func SpritePaths() map[SpriteKey]string {
	return map[SpriteKey]string{
		BodyMaleLight:  "resources/Universal-LPC-spritesheet/body/male/light.png",
		BodyMaleTanned: "resources/Universal-LPC-spritesheet/body/male/tanned.png",
		HexTerrain1:    "resources/tiles/wesnoth1.png",
	}
}

func LPCSpriteFolders() map[string]string {
	baseLpcFolder := "resources/Universal-LPC-spritesheet/"
	lpc := make(map[string]string)
	folders := []string{
		"body/male/",
	}
	for _, folderName := range folders {
		lpc[folderName] = baseLpcFolder + folderName
	}
	return lpc
}
