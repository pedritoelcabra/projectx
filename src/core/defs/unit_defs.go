package defs

import (
	"encoding/json"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"log"
	"os"
	"path/filepath"
)

type UnitDef struct {
	Name       string
	Equipments []EquipmentItemDef
}

type EquipmentItemDef struct {
	Chance     int
	Slot       string
	Graphic    string
	GraphicKey gfx.SpriteKey
}

func ResolveGraphicChance(equipments []EquipmentItemDef, slot string) string {
	totalPoints := 0
	for _, def := range equipments {
		if def.Slot != slot {
			continue
		}
		totalPoints += def.Chance
	}
	selected := randomizer.RandomInt(0, totalPoints)
	for _, def := range equipments {
		if def.Slot != slot {
			continue
		}
		if selected < def.Chance {
			return def.Graphic
		}
		selected -= def.Chance
	}
	log.Fatal("Unable to select random Equipment for slot " + slot)
	return ""
}

func UnitDefs() map[string]*UnitDef {
	return unitDefs
}

var unitDefs = make(map[string]*UnitDef)

func LoadUnitDefs() {
	unitDefs = make(map[string]*UnitDef)
	directoryPath, _ := filepath.Abs(defFolder + "Units")
	walkErr := filepath.Walk(directoryPath, func(path string, info os.FileInfo, walkErr error) error {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if filepath.Ext(path) != ".json" {
			return nil
		}
		dataStructure := &UnitDef{}
		err = json.NewDecoder(file).Decode(dataStructure)
		if err != nil {
			log.Fatal(err)
		}
		unitDefs[dataStructure.Name] = dataStructure
		return walkErr
	})
	if walkErr != nil {
		log.Fatal(walkErr)
	}
}
