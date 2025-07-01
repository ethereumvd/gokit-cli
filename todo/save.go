package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Storage[T any] struct{
	FileName string
}

func NewStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{FileName: filename}
}

func (storage * Storage[T]) SaveStorage(data T) error {

	fileData , err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return fmt.Errorf("Error Saving Data")
	}

	return os.WriteFile(storage.FileName, fileData, 0666)
}

func (storage *Storage[T]) LoadData(data *T) error {
	fileData, err := os.ReadFile(storage.FileName)

	if err != nil {
		return fmt.Errorf("Error loading from File")
	}

	return json.Unmarshal(fileData, data)
}