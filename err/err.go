package err

import "fmt"

type Error struct {
	ErrorID   int
	ErrorCode string
	Message   string
}

func (err Error) Error() string {
	return fmt.Sprintf("Error errID: %d, code: %s | msg: %s |", err.ErrorID, err.ErrorCode, err.Message)
}

func New(mgs string, code string, id int) Error {
	return Error{}
}

func Message(mgs string) Error {
	return Error{
		Message: mgs,
	}
}
