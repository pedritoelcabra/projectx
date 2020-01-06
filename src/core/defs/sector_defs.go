package defs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type SectorDef struct {
	Name           string
	CenterBuilding string
	Weight         int
	Size           int
}

func SectorDefs() map[string]*SectorDef {
	return sectorDefs
}

var sectorDefs = make(map[string]*SectorDef)

func LoadSectorDefs() {
	sectorDefs = make(map[string]*SectorDef)
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
