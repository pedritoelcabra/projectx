package file

import (
	"log"
	"os"
	"path/filepath"
)

const defFolder = "defs/"

func NewDefClass(className string) DefClass {
	switch className {
	case "Sectors":
		return &SectorDef{}
	case "Buildings":
		return &BuildingDef{}
	}
	log.Fatal("Trying to spawn non-existant Def Class")
	return &SectorDef{}
}

var defs = make(map[string]map[string]DefClass)

func GetDefs(defType string) map[string]DefClass {
	return defs[defType]
}

func InitDefs() {
	defs = make(map[string]map[string]DefClass)
	LoadDefFolder("Sectors")
	LoadDefFolder("Buildings")
}

func LoadDefFolder(name string) {
	defs[name] = LoadDefClass(NewDefClass(name), name)
}

func LoadDefClass(defClass DefClass, folderName string) map[string]DefClass {
	directoryPath, _ := filepath.Abs(defFolder + folderName)
	defClasses := make(map[string]DefClass, 0)
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
