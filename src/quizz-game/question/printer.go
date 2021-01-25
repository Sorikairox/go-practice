package question

import (
	"fmt"
)

type Printer struct {
}

func (p *Printer) PrintQuestion(question *Question) {
	fmt.Printf("%s\n", question.Label)
}
