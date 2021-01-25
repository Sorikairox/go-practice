package quizz

import (
	"container/list"
	"quizz-game/question"
)

type Quizz struct {
	questionParser  *question.Parser
	questionPrinter *question.Printer
	questions       *list.List
	answers         *list.List
}

func NewQuizz(parser *question.Parser, printer *question.Printer) *Quizz {
	quizz := Quizz{
		questionParser:  parser,
		questionPrinter: printer,
		questions:       list.New(),
	}
	quizz.questions.Init()
	return &quizz
}

func (quizz *Quizz) Play() {
	actualQuestion := quizz.questionParser.GetNextQuestion()
	for actualQuestion != nil {
		quizz.questions.PushBack(actualQuestion)
		quizz.questionPrinter.PrintQuestion(actualQuestion)
		quizz.answers.PushBack(quizz.questionParser.GetAnswer())
		actualQuestion = quizz.questionParser.GetNextQuestion()
	}

	for q := quizz.questions.Front(); q != nil; q = q.Next() {
		quizz.questionPrinter.PrintQuestion(q.Value.(*question.Question))
	}
}
