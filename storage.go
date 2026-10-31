package main

import (
	"encoding/json"
	"fmt"
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

	if len(chest.Slots) != 27 {
		return Chest{}, fmt.Errorf("invalid_storage")
	}

	return chest, nil
}

func (fr FileRepository) SaveChest(c Chest) error {

	if len(c.Slots) != 27 {
		return fmt.Errorf("invalid_storage")
	}

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(fr.FileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
