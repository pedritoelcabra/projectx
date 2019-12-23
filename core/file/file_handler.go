package file

import (
	"bytes"
	"encoding/json"
	"github.com/pedritoelcabra/projectx/world"
	"io"
	"log"
	"os"
	"path/filepath"
)

var saveGameBasePath = "save_games/"
var DefaultSaveGameName = "save.pxs"

type SaveGameData struct {
	Seed   int
	Tick   int
	Player world.Player
}

func getSaveGameFullPath(fileName string) string {
	absolutePath, err := filepath.Abs(saveGameBasePath)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		os.Mkdir(absolutePath, os.ModeDir|os.ModePerm)
	}
	return absolutePath + "/" + fileName
}

func SaveToFile(data SaveGameData, fileName string) {
	fullPath := getSaveGameFullPath(fileName)
	file, err := os.Create(fullPath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := Encode(data)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(file, jsonData)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadFromFile(fileName string) SaveGameData {
	fullPath := getSaveGameFullPath(fileName)
	file, err := os.Open(fullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	dataStructure := SaveGameData{}
	err = Decode(file, &dataStructure)
	if err != nil {
		log.Fatal(err)
	}
	return dataStructure
}

func Decode(reader io.Reader, structure interface{}) error {
	return json.NewDecoder(reader).Decode(structure)
}

func Encode(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func SaveGameExists(fileName string) bool {
	fullPath := getSaveGameFullPath(fileName)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return false
	}
	return true
}
