package file

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

var saveGameBasePath = "save_games/"
var DefaultSaveGameName = "save.pxs"

type SaveGameData struct {
	Seed int
}

func getSaveGameFullPath(fileName string) string {
	absolutePath, err := filepath.Abs(saveGameBasePath)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		os.Mkdir(absolutePath, os.ModeDir|os.ModePerm)
	}
	return saveGameBasePath + fileName
}

func SaveToFile(data SaveGameData, fileName string) {
	fullPath := getSaveGameFullPath(fileName)
	file, err := os.Create(fullPath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(file, jsonData)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadFromFile(fileName string) SaveGameData {
	file, err := os.Open(getSaveGameFullPath(fileName))
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	dataStructure := SaveGameData{}
	json.NewDecoder(file).Decode(dataStructure)
	return dataStructure
}

func Marshal(v interface{}) (io.Reader, error) {
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
