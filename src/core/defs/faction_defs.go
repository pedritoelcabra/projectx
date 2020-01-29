package defs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type FactionDef struct {
	Name            string
	DefaultRelation int
}

func FactionDefs() map[string]*FactionDef {
	return factionDefs
}

func GetFactionDef(name string) *FactionDef {
	return factionDefs[name]
}

var factionDefs = make(map[string]*FactionDef)

func LoadFactionDefs() {
	factionDefs = make(map[string]*FactionDef)
	directoryPath, _ := filepath.Abs(defFolder + "Factions")
	walkErr := filepath.Walk(directoryPath, func(path string, info os.FileInfo, walkErr error) error {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if filepath.Ext(path) != ".json" {
			return nil
		}
		dataStructure := &FactionDef{}
		err = json.NewDecoder(file).Decode(dataStructure)
		if err != nil {
			log.Fatal(err)
		}
		factionDefs[dataStructure.Name] = dataStructure
		return walkErr
	})
	if walkErr != nil {
		log.Fatal(walkErr)
	}
}
