package main

import (
	"flag"
	"fmt"
)

func main() {
	task := flag.String("task", "", "Команда")
	file := flag.String("file", "", "Команда")
	command := flag.String("command", "", "Команда для создания|чтения|удаления файла")
	mixQuestions := flag.Bool("mix", false, "Перемешать порядок вопросов в квизке")

	flag.Parse()

	switch *task {
	case "files":
		filesCli(*file, *command)
	case "quiz":
		quiz(*file, *mixQuestions)
	default:
		fmt.Println("Задание не найдено")
	}
}
