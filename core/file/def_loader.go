package file

import (
	"log"
	"os"
	"path/filepath"
)

const defFolder = "defs/"

type DefClass interface {
	GetName() string
}

type SectorDef struct {
	Name          string
	CenterGraphic string
	Weight        int
	Land          int
}

func (s *SectorDef) GetName() string {
	return s.Name
}

func NewDefClass(className string) DefClass {
	switch className {
	case "Sector":
		return &SectorDef{}

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
	defs["Sectors"] = LoadDefClass(NewDefClass("Sector"), "sectors")
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
