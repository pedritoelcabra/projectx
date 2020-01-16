package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var lpcSprites = make(map[string]SpriteKey)

func LoadSprites() {
	for key, path := range SpritePaths() {
		img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		AddImageToKey(img, key)
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
			key := AddImage(img)
			lpcSprites[(folderName + strings.Replace(fileName, ".png", "", -1))] = key
			return walkErr
		})
		if walkErr != nil {
			log.Fatal(walkErr)
		}
	}
}

func SpritePaths() map[SpriteKey]string {
	return map[SpriteKey]string{
		HexTerrain1: "resources/tiles/wesnoth1.png",
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
