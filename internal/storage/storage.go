package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jabrail059/Task-Tracker-CLI/models"
)

type Storage interface {
	Add(description string) (models.Task, error)
	GetAll() []models.Task
	Delete(Id int) error
	Done(Id int) error
	Update(Id int, NewDescription string) error
	InProgress(Id int) error
	GetDoneTasks() []models.Task
	GetInProgressTasks() []models.Task
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
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
	for index := range f.tasks {
		if f.tasks[index].ID == Id {
			OldTime := f.tasks[index].UpdatedAt
			OldStatus := f.tasks[index].Status
			f.tasks[index].Status = "done"
			f.tasks[index].UpdatedAt = time.Now()
			err := f.save()
			if err != nil {
				f.tasks[index].Status = OldStatus
				f.tasks[index].UpdatedAt = OldTime
				return fmt.Errorf("Не удалось сохранить изменения. Статус задачи не изменён.")
			}
			return nil
		}
	}
	return fmt.Errorf("Задача с Id: %d не найдена.", Id)
}

func (f *FileStorage) InProgress(Id int) error {
	for index := range f.tasks {
		if f.tasks[index].ID == Id {
			OldTime := f.tasks[index].UpdatedAt
			OldStatus := f.tasks[index].Status
			f.tasks[index].Status = "in-progress"
			f.tasks[index].UpdatedAt = time.Now()
			err := f.save()
			if err != nil {
				f.tasks[index].Status = OldStatus
				f.tasks[index].UpdatedAt = OldTime
				return fmt.Errorf("Не удалось сохранить изменения. Статус задачи не изменён.")
			}

			return nil
		}
	}
	return fmt.Errorf("Задача с Id: %d не найдена.", Id)
}

func (f *FileStorage) Update(Id int, NewDescription string) error {
	for index := range f.tasks {
		if f.tasks[index].ID == Id {
			OldDescription := f.tasks[index].Description
			OldTime := f.tasks[index].UpdatedAt
			f.tasks[index].Description = NewDescription
			f.tasks[index].UpdatedAt = time.Now()
			err := f.save()
			if err != nil {
				f.tasks[index].Description = OldDescription
				f.tasks[index].UpdatedAt = OldTime
				return fmt.Errorf("Не удалось сохранить изменения. Описание задачи не изменено.")
			}
			return nil
		}
	}
	return fmt.Errorf("Задача с Id: %d не найдена.", Id)
}

func (f *FileStorage) GetDoneTasks() []models.Task {
	DoneTasks := []models.Task{}
	for index := range f.tasks {
		if f.tasks[index].Status == "done" {
			DoneTasks = append(DoneTasks, f.tasks[index])
		}
	}
	return DoneTasks
}

func (f *FileStorage) GetInProgressTasks() []models.Task {
	InProgressTasks := []models.Task{}
	for index := range f.tasks {
		if f.tasks[index].Status == "in-progress" {
			InProgressTasks = append(InProgressTasks, f.tasks[index])
		}
	}
	return InProgressTasks
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
