package defs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type MaterialDef struct {
	Name      string
	ID        int
	Weight    int
	StackSize int
}

func GetMaterialDefs() map[string]*MaterialDef {
	return materialDefs
}

func GetMaterialDef(name string) *MaterialDef {
	return materialDefs[name]
}

func GetMaterialKeyByName(name string) int {
	return materialDefs[name].ID
}

func GetMaterialDefByKey(key int) *MaterialDef {
	for _, def := range materialDefs {
		if def.ID == key {
			return def
		}
	}
	log.Fatal("Invalid material def: " + strconv.Itoa(key))
	return nil
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
