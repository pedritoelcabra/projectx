package defs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type MaterialDef struct {
	Name      string
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
		dataStructure := &MaterialDef{}
		err = json.NewDecoder(file).Decode(dataStructure)
		if err != nil {
			log.Fatal(err)
		}
		materialDefs[dataStructure.Name] = dataStructure
		return walkErr
	})
	if walkErr != nil {
		log.Fatal(walkErr)
	}
}
