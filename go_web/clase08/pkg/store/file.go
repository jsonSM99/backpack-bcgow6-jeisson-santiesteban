package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

const (
	FileType  Type = "file"
	MongoType Type = "mongo"
)

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &fileStore{fileName}
	}
	return nil
}

type fileStore struct {
	FilePath string
}

func (f *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(f.FilePath, fileData, 0644)
}

func (f *fileStore) Read(data interface{}) error {
	// Open our jsonFile
	jsonFile, err := os.Open(f.FilePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &data)
	return err
}
