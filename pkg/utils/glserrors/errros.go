package glserrors

import "fmt"

type GLSError struct {
	Code    string
	Message string
	Err     error
}

func (e GLSError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s, Err: %s", e.Code, e.Message, e.Err)
}

func PassExecutionErr(glsError *GLSError, err error) *GLSError {
	e := *glsError
	e.Err = err

	return &e
}
