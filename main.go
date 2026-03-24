package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/stxa005/Task-Tracker-CLI/commands"
	"github.com/stxa005/Task-Tracker-CLI/models"
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
		switch {
		case len(args) == 3 && args[2] == "done":
			DoneTasks := commands.GetDoneTasks(store)
			if len(DoneTasks) == 0 {
				fmt.Println("Выполненных задач нет.")
				return
			}
			PrintTasks(DoneTasks)
		case len(args) == 3 && args[2] == "in-progress":
			InProgressTasks := commands.GetInProgressTasks(store)
			if len(InProgressTasks) == 0 {
				fmt.Println("Не выполненных задач нет.")
				return
			}
			PrintTasks(InProgressTasks)
		default:
			ListOfTasks := commands.GetAll(store)
			if len(ListOfTasks) == 0 {
				fmt.Println("Задач нет")
				return
			}
			if len(args) > 2 {
				fmt.Printf("Команда не распознана: %s\n", args[3])
			}
			PrintTasks(ListOfTasks)
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
			fmt.Println("Ошибка! Недостаточно аргументов!")
			return
		}
		Id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Ошибка! Не удалось получить Id задачи!")
			return
		}
		err = commands.Done(store, Id)
		if err != nil {
			fmt.Println("Ошибка! Не удалось изменить статус задачи! " + err.Error())
			return
		}
		fmt.Printf("Id: %d. Статус задачи изменён на \"Done\".\n", Id)
	case "in-progress":
		if len(args) < 3 {
			fmt.Println("Ошибка! Укажите Id задачи!")
			return
		}
		Id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Ошибка! Не удалось получить Id задачи!")
			return
		}
		err = commands.InProgress(store, Id)
		if err != nil {
			fmt.Println("Ошибка! Не удалось изменить статус задачи! " + err.Error())
			return
		}
		fmt.Printf("Id: %d. Статус задачи изменён на \"In-Progress\".\n", Id)
	case "update":
		if len(args) < 4 {
			fmt.Println("Ошибка! Недостаточно аргументов!")
			return
		}
		Id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Ошибка! Не удалось получить Id задачи!")
			return
		}
		if len(args[3]) == 0 {
			fmt.Println("Ошибка! Укажите описание новой задачи!")
			return
		}
		err = commands.Update(store, Id, args[3])
		if err != nil {
			fmt.Println("Ошибка! Не удалось изменить описание задачи! " + err.Error())
			return
		}
		fmt.Printf("Id: %d. Описание задачи изменено.\n", Id)
	case "help":
		for _, function := range Help() {
			fmt.Println(function)
		}
	default:
		fmt.Println("Ошибка! Неизвестная команда!")
	}
}

func PrintTasks(tasks []models.Task) {
	for _, task := range tasks {
		fmt.Printf("ID: %d;\n Задача: %s; Статус: %v;\n Дата создания: %v; Дата Последнего изменения: %v\n\n", task.ID, task.Description, task.Status, task.CreatedAt.Local().Format("02.01.2006 15:04"), task.UpdatedAt.Local().Format("02.01.2006 15:04"))
	}
}

func Help() []string {
	return []string{
		"help - Вывод всех команд",
		"add - Добавление новой задачи",
		"delete - Удаление задачи",
		"list - Вывод всех задач",
		"list done - Вывод выполненных задач",
		"list in-progress - Вывод недовыполненных задач",
		"update - Обновление описания задачи",
		"done - Изменение статуса задачи на \"Выполнено\"",
		"in-progress - Изменение статуса задачи на \"В работе\"",
	}
}
