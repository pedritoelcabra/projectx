package defs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type ResourceDef struct {
	Name      string
	Weight    int
	StackSize int
}

func ResourceDefs() map[string]*ResourceDef {
	return resourceDefs
}

func GetResourceDef(name string) *ResourceDef {
	return resourceDefs[name]
}

var resourceDefs = make(map[string]*ResourceDef)

func LoadResourceDefs() {
	resourceDefs = make(map[string]*ResourceDef)
	directoryPath, _ := filepath.Abs(defFolder + "Resources")
	walkErr := filepath.Walk(directoryPath, func(path string, info os.FileInfo, walkErr error) error {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if filepath.Ext(path) != ".json" {
			return nil
		}
		dataStructure := &ResourceDef{}
		err = json.NewDecoder(file).Decode(dataStructure)
		if err != nil {
			log.Fatal(err)
		}
		resourceDefs[dataStructure.Name] = dataStructure
		return walkErr
	})
	if walkErr != nil {
		log.Fatal(walkErr)
	}
}
