package file

import (
	"github.com/pedritoelcabra/projectx/defs"
	"log"
	"os"
	"path/filepath"
)

const defFolder = "defs/"

func NewDefClass(className string) defs.DefClass {
	switch className {
	case "Sectors":
		return &defs.SectorDef{}
	case "Buildings":
		return &defs.BuildingDef{}
	}
	log.Fatal("Trying to spawn non-existant Def Class")
	return &defs.SectorDef{}
}

func InitDefs() {
	defs.Definitions = make(map[string]map[string]defs.DefClass)
	LoadDefFolder("Sectors")
	LoadDefFolder("Buildings")
}

func LoadDefFolder(name string) {
	defs.Definitions[name] = LoadDefClass(NewDefClass(name), name)
}

func LoadDefClass(defClass defs.DefClass, folderName string) map[string]defs.DefClass {
	directoryPath, _ := filepath.Abs(defFolder + folderName)
	defClasses := make(map[string]defs.DefClass, 0)
	walkErr := filepath.Walk(directoryPath, func(path string, info os.FileInfo, walkErr error) error {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if filepath.Ext(path) != ".json" {
			return nil
		}
		dataStructure := defClass
		err = Decode(file, &dataStructure)
		if err != nil {
			log.Fatal(err)
		}
		defClasses[dataStructure.GetName()] = dataStructure
		return walkErr
	})
	if walkErr != nil {
		log.Fatal(walkErr)
	}
	return defClasses
}
