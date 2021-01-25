package question

type Reader interface {
	Read() ([]string, error)
	ReadString() (string, error)
}
