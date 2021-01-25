package question

import (
	"io"
)

type Parser struct {
	questionReader Reader
	answerReader   Reader
	ActualQuestion *Question
}

func NewQuestionParser(reader Reader) *Parser {
	parser := Parser{questionReader: reader}
	parser.ActualQuestion = NewQuestion()
	return &parser
}

func (questionParser *Parser) GetNextQuestion() *Question {
	record, err := questionParser.questionReader.Read()
	if err == io.EOF {
		return nil
	}
	questionParser.ActualQuestion.Label = record[0]
	questionParser.ActualQuestion.Answer = record[1]
	return questionParser.ActualQuestion
}

func (questionParser *Parser) GetAnswer() *string {
	answer, err := questionParser.answerReader.ReadString()
	if err != nil {
		return nil
	}
	return &answer
}
