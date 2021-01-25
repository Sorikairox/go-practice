package question

type Question struct {
	Label  string
	Answer string
}

func NewQuestion() *Question {
	return &Question{}
}
