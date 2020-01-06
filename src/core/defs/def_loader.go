package defs

import (
	"encoding/json"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"log"
	"os"
	"path/filepath"
)

const defFolder = "defs/"

type SectorDef struct {
	Name           string
	CenterBuilding string
	Weight         int
	Size           int
}

type BuildingDef struct {
	Name    string
	Graphic string
}

type VegetationDef struct {
	Name         string
	MovementCost float64
	Weight       int
	Graphics     []string
}

func (v *VegetationDef) GetGraphic() string {
	return v.Graphics[randomizer.RandomInt(0, len(v.Graphics)-1)]
}

func VegetationDefs() map[int]*VegetationDef {
	return vegetationDefs
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

func BuildingDefs() map[string]*BuildingDef {
	return buildingDefs
}

func SectorDefs() map[string]*SectorDef {
	return sectorDefs
}

var buildingDefs = make(map[string]*BuildingDef)
var sectorDefs = make(map[string]*SectorDef)
var vegetationDefs = make(map[int]*VegetationDef)
var vegetationNames = make(map[string][]int)
var vegetationSprites = make(map[int]gfx.Sprite)

func InitDefs() {
	sectorDefs = make(map[string]*SectorDef)
	buildingDefs = make(map[string]*BuildingDef)
	vegetationDefs = make(map[int]*VegetationDef)
	vegetationSprites = make(map[int]gfx.Sprite)
	vegetationNames = make(map[string][]int)
	LoadBuildingDefs()
	LoadSectorDefs()
	LoadVegetationDefs()
}

func LoadVegetationDefs() {
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

func LoadBuildingDefs() {
	directoryPath, _ := filepath.Abs(defFolder + "Buildings")
	walkErr := filepath.Walk(directoryPath, func(path string, info os.FileInfo, walkErr error) error {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if filepath.Ext(path) != ".json" {
			return nil
		}
		dataStructure := &BuildingDef{}
		err = json.NewDecoder(file).Decode(dataStructure)
		if err != nil {
			log.Fatal(err)
		}
		buildingDefs[dataStructure.Name] = dataStructure
		return walkErr
	})
	if walkErr != nil {
		log.Fatal(walkErr)
	}
}

func LoadSectorDefs() {
	directoryPath, _ := filepath.Abs(defFolder + "Sectors")
	walkErr := filepath.Walk(directoryPath, func(path string, info os.FileInfo, walkErr error) error {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if filepath.Ext(path) != ".json" {
			return nil
		}
		dataStructure := &SectorDef{}
		err = json.NewDecoder(file).Decode(dataStructure)
		if err != nil {
			log.Fatal(err)
		}
		sectorDefs[dataStructure.Name] = dataStructure
		return walkErr
	})
	if walkErr != nil {
		log.Fatal(walkErr)
	}
}
