package defs

import (
	"encoding/json"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"log"
	"os"
	"path/filepath"
)

var resourceDefs = make(map[int]*ResourceDef)
var resourceTotals = make(map[string]int)
var resourceNames = make(map[string][]int)
var resourceSprites = make(map[int]gfx.Sprite)

type ResourceDef struct {
	Name           string
	MovementCost   float64
	Weight         int
	ResourceAmount int
	Resource       string
	Graphics       []string
}

func (v *ResourceDef) GetGraphic() string {
	return v.Graphics[randomizer.RandomInt(0, len(v.Graphics)-1)]
}

func AddResourceLocation(name string) {
	resourceTotals[name]++
}

func GetResourceTotals() map[string]int {
	return resourceTotals
}

func GetResourceLocationTotals(name string) int {
	return resourceTotals[name]
}

func ResourceById(id int) *ResourceDef {
	return resourceDefs[id]
}

func ResourceByName(name string) int {
	return resourceNames[name][randomizer.RandomInt(0, len(resourceNames[name])-1)]
}

func DrawResource(id int, screen *gfx.Screen, x, y float64) {
	resourceSprites[id].DrawSprite(screen, x, y)
}

func LoadResourceDefs() {
	resourceDefs = make(map[int]*ResourceDef)
	resourceSprites = make(map[int]gfx.Sprite)
	resourceNames = make(map[string][]int)
	directoryPath, _ := filepath.Abs(defFolder + "Resource")
	walkErr := filepath.Walk(directoryPath, func(path string, info os.FileInfo, walkErr error) error {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		ext := filepath.Ext(path)
		if ext != ".json" {
			return nil
		}
		dataStructure := &ResourceDef{}
		err = json.NewDecoder(file).Decode(dataStructure)
		if err != nil {
			log.Fatal(err)
		}
		for _, graphicName := range dataStructure.Graphics {
			id := len(resourceDefs)
			resourceDefs[id] = dataStructure
			resourceTotals[dataStructure.Name] = 0
			resourceSprites[id] = gfx.NewHexSprite(gfx.GetSpriteKey(graphicName))
			resourceNames[dataStructure.Name] = append(resourceNames[dataStructure.Name], id)
		}
		return walkErr
	})
	if walkErr != nil {
		log.Fatal(walkErr)
	}
}
