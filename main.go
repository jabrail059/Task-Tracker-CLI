package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/stxa005/Task-Tracker-CLI/commands"
	"github.com/stxa005/Task-Tracker-CLI/storage"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Пожалуйста, передайте хотя бы один аргумент")
		return
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Не удалось получить доступ к папке")
		return
	}
	configDir := filepath.Join(homeDir, ".task-tracker")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		fmt.Println("Не удалось получить доступ к папке")
		return
	}
	filePath := filepath.Join(configDir, "tasks.json")
	store, err := (storage.NewFileStorage(filePath))
	if err != nil {
		fmt.Println("Не удалось создать хранилище")
		return
	}

	command := args[1]
	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Ошибка: Сообщение не передано!")
			return
		}
		task, err := commands.Add(store, args[2])
		if err != nil {
			fmt.Println("Ошибка! Не удалось получить задачу!")
			return
		}
		fmt.Printf("Задача добавлена с ID: %d\n", task.ID)
	case "list":
		listOfTasks := commands.GetAll(store)
		if len(listOfTasks) == 0 {
			fmt.Println("Задач нет")
			return
		}
		for _, task := range listOfTasks {
			fmt.Printf("ID: %d; Задача: %s; Статус: %v\n", task.ID, task.Description, task.Completed)
		}
	case "delete":
		if len(args) < 3 {
			fmt.Println("Ошибка! Укажите Id задачи!")
			return
		}
		Id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Ошикба! Не удалось получить Id задачи!")
			return
		}
		err = commands.Delete(store, Id)
		if err != nil {
			fmt.Println("Ошибка! Не удалось удалить задачу! " + err.Error())
			return
		}
		fmt.Printf("Id: %d. Задача удалена\n", Id)
	case "done":
		if len(args) < 3 {
			fmt.Println("Ошибка! Укажите Id задачи!")
			return
		}
		Id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Ошибка! не удалось получить Id задачи!")
			return
		}
		err = commands.Done(store, Id)
		if err != nil {
			fmt.Println("Ошибка! Не удалось изменить статус задачи! " + err.Error())
			return
		}
		fmt.Printf("Id: %d. Статус задачи изменён.\n", Id)

	default:
		fmt.Println("Ошибка! Неизвестная команда!")
	}
}
