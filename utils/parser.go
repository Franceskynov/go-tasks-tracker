package utils

import (
	"encoding/json"
	"errors"
	"os"
)

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Tasks struct {
	Tasks [] Task `json:"tasks"`
}

func JSONWriter(fileName string, data Tasks) error {
	file, err := os.Create(fileName)

	if err != nil {
		return errors.New("error creating or reading the file")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	
	if err = encoder.Encode(data); err != nil {
		return errors.New("error writing data into the file")
	} else {
		return nil
	}
}

func JSONReader(fileName string) (Tasks, error) {
	var tasks Tasks
	file, err := os.Open(fileName)
	if err != nil {
		return tasks, errors.New("error at reading existing file")
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	if err = decoder.Decode(&tasks); err != nil {
		return tasks, errors.New("error decoding data from file")
	} else {
		return tasks, nil
	}
}

func FindIndex(list[] Task, id int) int {
	idx := -1
	for i, v := range list {
		if v.ID == id {
			idx = i
			break
		}
	}
	return idx
}
