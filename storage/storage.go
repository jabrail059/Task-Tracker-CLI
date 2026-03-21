package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/stxa005/Task-Tracker-CLI/models"
)

type Storage interface {
	Add(description string) (models.Task, error)
	GetAll() []models.Task
	Delete(Id int) error
	Done(Id int) error
}

type FileStorage struct {
	filePath string
	tasks    []models.Task
	nextID   int
}

func (f *FileStorage) Add(description string) (models.Task, error) {
	task := models.Task{
		ID:          f.nextID,
		Description: description,
		Completed:   false,
	}
	f.nextID++
	f.tasks = append(f.tasks, task)
	err := f.save()
	if err != nil {
		f.nextID--
		f.tasks = f.tasks[:len(f.tasks)-1]
		return models.Task{}, err
	}
	return task, nil
}

func (f *FileStorage) GetAll() []models.Task {
	return f.tasks
}

func (f *FileStorage) Delete(Id int) error {
	for index := range f.tasks {
		if f.tasks[index].ID == Id {
			deletedTask := f.tasks[index]
			f.tasks = append(f.tasks[:index], f.tasks[index+1:]...)
			err := f.save()
			if err != nil {
				f.tasks = append(f.tasks[:index], append([]models.Task{deletedTask}, f.tasks[index:]...)...)
				return fmt.Errorf("Не удалось сохранить изменения. Задача не удалена.")
			}
			return nil
		}
	}
	return fmt.Errorf("Задача с Id: %d не найдена!\n", Id)
}

func (f *FileStorage) Done(Id int) error {
	for index, task := range f.tasks {
		if task.ID == Id {
			f.tasks[index].Completed = true
			err := f.save()
			if err != nil {
				f.tasks[index].Completed = false
				return fmt.Errorf("Не удалось сохранить изменения. Статус задачи не изменён.")
			}
			return nil
		}
	}
	return fmt.Errorf("Задача с Id: %d не найдена.", Id)
}

func (f *FileStorage) save() error {
	file, err := os.Create(f.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	JsonData, err := json.MarshalIndent(f.tasks, "", "\t")
	if err != nil {
		return err
	}
	_, err = file.Write(JsonData)
	return err
}

func NewFileStorage(filePath string) (*FileStorage, error) {
	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		return &FileStorage{
			filePath: filePath,
			tasks:    []models.Task{},
			nextID:   1,
		}, nil
	} else if err != nil {
		return nil, err
	} else {
		defer file.Close()
	}
	tasks := []models.Task{}
	err = json.NewDecoder(file).Decode(&tasks)
	if err == io.EOF {
		return &FileStorage{
			filePath: filePath,
			tasks:    []models.Task{},
			nextID:   1,
		}, nil
	} else if err != nil {
		return nil, err
	}
	maxId := 0

	if len(tasks) > 0 {
		for _, task := range tasks {
			if task.ID > maxId {
				maxId = task.ID
			}
		}
	}
	return &FileStorage{
		filePath: filePath,
		tasks:    tasks,
		nextID:   maxId + 1,
	}, nil
}
