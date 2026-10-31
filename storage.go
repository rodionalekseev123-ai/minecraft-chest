package main

import (
	"encoding/json"
	"os"
)

type FileRepository struct {
	FileName string
}

func (fr FileRepository) LoadChest() (Chest, error) {
	data, err := os.ReadFile(fr.FileName)
	if err != nil {
		println("error while reading file")
		return Chest{}, err
	}

	var chest Chest //объявлем переменную куда будем сохранять json

	err = json.Unmarshal(data, &chest)
	if err != nil {
		println("error while unmarshalling")
		return Chest{}, err
	}
	return chest, nil
}
