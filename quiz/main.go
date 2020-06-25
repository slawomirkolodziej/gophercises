package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// QuizItem is single part of quiz containing question and its answer
type QuizItem struct {
	question string
	answer   string
}

var problems = []QuizItem{QuizItem{"question1", "answer1"}}

func performQuiz(quiz []QuizItem, score *int) {
	var userAnswer string
	for _, problem := range quiz {
		fmt.Println(problem.question)
		fmt.Scanln(&userAnswer)
		if userAnswer == problem.answer {
			*score++
		}
	}
}

func printScore(score int, total int) {
	fmt.Println("Your correct answers:", score)
	fmt.Println("Total questions", total)
}

func readFromCsv(path string) [][]string {
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Can't read from file", err)
	}

	stringReader := strings.NewReader(string(fileData))
	csvReader := csv.NewReader(stringReader)
	csv, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Can't parse csv", err)
	}

	return csv
}

func createQuizFromCsvData(csvData [][]string) []QuizItem {
	var quiz = make([]QuizItem, len(csvData))

	for index, csvRow := range csvData {
		quiz[index] = QuizItem{csvRow[0], csvRow[1]}
	}

	return quiz
}

func setupTimer(timeLimit int, score *int, total *int) {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	go func() {
		<-timer.C
		printScore(*score, *total)
		os.Exit(1)
	}()
}

func main() {
	quizPath := flag.String("quizPath", "problems.csv", "Path to csv with quiz")
	timeLimit := flag.Int("timeLimit", 10, "Time limit of quiz")
	flag.Parse()

	csv := readFromCsv(*quizPath)
	quiz := createQuizFromCsvData(csv)
	total := len(quiz)
	score := 0

	setupTimer(*timeLimit, &score, &total)
	performQuiz(quiz, &score)
	printScore(score, total)
}
