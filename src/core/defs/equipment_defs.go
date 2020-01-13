package defs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type EquipmentDef struct {
	Name    string
	Graphic string
}

func EquipmentDefs() map[string]*EquipmentDef {
	return equipmentDefs
}

var equipmentDefs = make(map[string]*EquipmentDef)

func LoadEquipmentDefs() {
	equipmentDefs = make(map[string]*EquipmentDef)
	directoryPath, _ := filepath.Abs(defFolder + "Equipments")
	walkErr := filepath.Walk(directoryPath, func(path string, info os.FileInfo, walkErr error) error {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if filepath.Ext(path) != ".json" {
			return nil
		}
		dataStructure := &EquipmentDef{}
		err = json.NewDecoder(file).Decode(dataStructure)
		if err != nil {
			log.Fatal(err)
		}
		equipmentDefs[dataStructure.Name] = dataStructure
		return walkErr
	})
	if walkErr != nil {
		log.Fatal(walkErr)
	}
}
