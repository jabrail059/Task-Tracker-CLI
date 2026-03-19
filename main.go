package main

import (
	"fmt"
	"os"

	"github.com/stxa005/Task-Tracker-CLI/storage"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Пожалуйста, передайте хотя бы один аргумент")
		return
	}

	command := args[1]
	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Ошибка: Сообщение не передано!")
			return
		}
		fmt.Printf("Задача добавлена с ID: %d\n", storage.Add(args[2]).ID)
	case "list":
		for _, task := range storage.GetAll() {
			fmt.Printf("ID: %d; Задача: %s; Статус: %v\n", task.ID, task.Description, task.Completed)
		}
	default:
		fmt.Println("Ошибка! Неизвестная команда!")
	}
}
