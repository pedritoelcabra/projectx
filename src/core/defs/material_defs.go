package defs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type MaterialDef struct {
	Name      string
	ID        int
	Weight    int
	StackSize int
}

func MaterialDefs() map[string]*MaterialDef {
	return materialDefs
}

func GetMaterialDef(name string) *MaterialDef {
	return materialDefs[name]
}

var materialDefs = make(map[string]*MaterialDef)

func LoadMaterialDefs() {
	id := 0
	materialDefs = make(map[string]*MaterialDef)
	directoryPath, _ := filepath.Abs(defFolder + "Materials")
	walkErr := filepath.Walk(directoryPath, func(path string, info os.FileInfo, walkErr error) error {

		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if filepath.Ext(path) != ".json" {
			return nil
		}
		id++
		dataStructure := &MaterialDef{}
		err = json.NewDecoder(file).Decode(dataStructure)
		if err != nil {
			log.Fatal(err)
		}
		dataStructure.ID = id
		materialDefs[dataStructure.Name] = dataStructure
		return walkErr
	})
	if walkErr != nil {
		log.Fatal(walkErr)
	}
}
