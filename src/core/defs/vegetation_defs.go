package defs

import (
	"encoding/json"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"log"
	"os"
	"path/filepath"
)

type VegetationDef struct {
	Name         string
	MovementCost float64
	Weight       int
	Resource     string
	Graphics     []string
}

func (v *VegetationDef) GetGraphic() string {
	return v.Graphics[randomizer.RandomInt(0, len(v.Graphics)-1)]
}

func VegetationById(id int) *VegetationDef {
	return vegetationDefs[id]
}

func VegetationByName(name string) int {
	return vegetationNames[name][randomizer.RandomInt(0, len(vegetationNames[name])-1)]
}

func DrawVegetation(id int, screen *gfx.Screen, x, y float64) {
	vegetationSprites[id].DrawSprite(screen, x, y)
}

var vegetationDefs = make(map[int]*VegetationDef)
var vegetationNames = make(map[string][]int)
var vegetationSprites = make(map[int]gfx.Sprite)

func LoadVegetationDefs() {
	vegetationDefs = make(map[int]*VegetationDef)
	vegetationSprites = make(map[int]gfx.Sprite)
	vegetationNames = make(map[string][]int)
	directoryPath, _ := filepath.Abs(defFolder + "Vegetation")
	walkErr := filepath.Walk(directoryPath, func(path string, info os.FileInfo, walkErr error) error {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if filepath.Ext(path) != ".json" {
			return nil
		}
		dataStructure := &VegetationDef{}
		err = json.NewDecoder(file).Decode(dataStructure)
		if err != nil {
			log.Fatal(err)
		}
		for _, graphicName := range dataStructure.Graphics {
			id := len(vegetationDefs)
			vegetationDefs[id] = dataStructure
			vegetationSprites[id] = gfx.NewHexSprite(gfx.GetSpriteKey(graphicName))
			vegetationNames[dataStructure.Name] = append(vegetationNames[dataStructure.Name], id)
		}
		return walkErr
	})
	if walkErr != nil {
		log.Fatal(walkErr)
	}
}
