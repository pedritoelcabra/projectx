package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func LoadSprites() {
	for key, path := range SpritePaths() {
		img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		AddImageToKey(img, key)
	}
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
			name := folderName + strings.Replace(fileName, ".png", "", -1)
			AddLpcImage(img, name)
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
		"legs/pants/male/",
		"torso/shirts/longsleeve/male/",
	}
	for _, folderName := range folders {
		lpc[folderName] = baseLpcFolder + folderName
	}
	return lpc
}
