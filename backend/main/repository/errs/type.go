package errs

type RecordNotFoundError struct {
	Message string
}

func (e *RecordNotFoundError) Error() string {
	return e.Message
}

func NewRecordNotFoundError() *RecordNotFoundError {
	return &RecordNotFoundError{
		Message: "Record not found",
	}
}
