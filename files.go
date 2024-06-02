package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func filesCli(fileName, command string) {
	flag.Parse()

	if "" == fileName {
		panic("Ошибка: необходимо указать имя файла.")
	}

	switch command {
	case "create":
		createFile(fileName)
	case "delete":
		deleteFile(fileName)
	case "read":
		readFile(fileName)
	default:
		fmt.Println("Команда не определена")
	}
}

func createFile(path string) {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("Создан файл:", path)
}

func readFile(path string) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		if err == io.EOF {
			break
		}

		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	fmt.Println("Содержимое файла:", path)
	fmt.Println(string(text))
}

func deleteFile(path string) {
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("Удалён файл:", path)
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return err != nil
}
