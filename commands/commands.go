package commands

import (
	"github.com/jabrail059/Task-Tracker-CLI/models"
	"github.com/jabrail059/Task-Tracker-CLI/storage"
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

func InProgress(s storage.Storage, Id int) error {
	return s.InProgress(Id)
}

func GetDoneTasks(s storage.Storage) []models.Task {
	return s.GetDoneTasks()
}

func GetInProgressTasks(s storage.Storage) []models.Task {
	return s.GetInProgressTasks()
}

func Update(s storage.Storage, Id int, NewDescription string) error {
	return s.Update(Id, NewDescription)
}
