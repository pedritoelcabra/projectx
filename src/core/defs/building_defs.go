package defs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type BuildingDef struct {
	Name                string
	Description         string
	Graphic             string
	ConstructionGraphic string
	Unit                string
	UnitLimit           int
	UnitTimer           int
	Influence           int
	ConstructionWork    int
}

func (def *BuildingDef) GetName() string {
	return def.Name
}

func (def *BuildingDef) GetStats() string {
	return def.Description
}

func BuildingDefs() map[string]*BuildingDef {
	return buildingDefs
}

func GetBuildingDef(name string) *BuildingDef {
	return buildingDefs[name]
}

var buildingDefs = make(map[string]*BuildingDef)

func LoadBuildingDefs() {
	buildingDefs = make(map[string]*BuildingDef)
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
