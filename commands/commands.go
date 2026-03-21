package commands

import (
	"github.com/stxa005/Task-Tracker-CLI/models"
	"github.com/stxa005/Task-Tracker-CLI/storage"
)

func Add(s storage.Storage, description string) (*models.Task, error) {
	task, err := s.Add(description)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func GetAll(s storage.Storage) []models.Task {
	return s.GetAll()
}

func Delete(s storage.Storage, Id int) error {
	return s.Delete(Id)
}

func Done(s storage.Storage, Id int) error {
	return s.Done(Id)
}
