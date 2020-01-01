package defs

import (
	"encoding/json"
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

func BuildingDefs() map[string]*BuildingDef {
	return buildingDefs
}

func SectorDefs() map[string]*SectorDef {
	return sectorDefs
}

var buildingDefs = make(map[string]*BuildingDef)
var sectorDefs = make(map[string]*SectorDef)

func InitDefs() {
	sectorDefs = make(map[string]*SectorDef)
	buildingDefs = make(map[string]*BuildingDef)
	LoadBuildingDefs()
	LoadSectorDefs()
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
