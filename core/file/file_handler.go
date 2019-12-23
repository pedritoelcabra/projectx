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

func SaveToFile(data SaveGameData, fileName string) {
	absolutePath, err := filepath.Abs(saveGameBasePath)
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		os.Mkdir(absolutePath, os.ModeDir|os.ModePerm)
	}
	fullPath := saveGameBasePath + fileName
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(fullPath)
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
	defer file.Close()
}

func Marshal(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func SaveGameExists(fileName string) bool {
	saveGameFilePath := saveGameBasePath + "/" + fileName
	absolutePath, err := filepath.Abs(saveGameFilePath)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		return false
	}
	return true
}
