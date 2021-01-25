package main

import (
	"encoding/csv"
	"os"
	"quizz-game/question"
	"quizz-game/quizz"
)

func main() {
	fd, err := os.Open("./quizz.csv")
	if err != nil {
		os.Exit(1)
	}

	reader := csv.NewReader(fd)
	questionParser := question.NewQuestionParser(reader)
	questionPrinter := question.Printer{}

	quizz := quizz.NewQuizz(questionParser, &questionPrinter)
	quizz.Play()
}
