package storage

import . "github.com/stxa005/Task-Tracker-CLI/models"

var repository []Task
var idCounter int

func Add(description string) Task {
	idCounter++
	task := Task{
		ID:          idCounter,
		Description: description,
		Completed:   false,
	}
	repository = append(repository, task)
	return task
}

func GetAll() []Task {
	return repository
}
