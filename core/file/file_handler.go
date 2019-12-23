package file

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

type SaveGameData struct {
	Seed int
}

func SaveToFile(data SaveGameData, fileName string) error {
	basePath := "save_games/"
	absolutePath, err := filepath.Abs(basePath)
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		os.Mkdir(absolutePath, os.ModeDir|os.ModePerm)
	}
	fullPath := basePath + fileName
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(fullPath)
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := Marshal(data)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, jsonData)
	defer file.Close()
	return err
}

var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
