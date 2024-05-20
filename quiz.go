package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type question struct {
	question string
	answer   string
}

func quiz(fileName string, mixQuestions bool) {
	if fileName == "" {
		fileName = "problems.csv"
	}

	var quiz []question
	var score int
	score = 0
	quiz = fetchQuestions(fileName, mixQuestions)
	var answer string

	for i, questions := range quiz {
		fmt.Printf("Вопрос %d: %s? ", i+1, questions.question)
		fmt.Scan(&answer)

		// избавляемся от чувствительности к регистру
		if strings.ToLower(strings.Trim(answer, " ")) == strings.ToLower(strings.Trim(questions.answer, " ")) {
			score++
		}
	}

	fmt.Printf("Вы ответили корректно на %d вопросов из %d вопросов всего.\n", score, len(quiz))
}

func fetchQuestions(fileName string, mixQuestions bool) []question {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка: CSV-файл не может быть прочитан.")
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	questions, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Ошибка: Вопросы из csv-файла не могут быть прочитаны.")
		panic(err)
	}

	var quiz []question

	for _, line := range questions {
		qa := question{
			question: line[0],
			answer:   line[len(line)-1],
		}

		quiz = append(quiz, qa)
	}

	// перемешиваем ответы из файла
	if mixQuestions {
		r := rand.New(rand.NewSource(time.Now().Unix()))

		for n := len(quiz); n > 0; n-- {
			randIndex := r.Intn(n)
			quiz[n-1], quiz[randIndex] = quiz[randIndex], quiz[n-1]
		}
	}

	return quiz
}
