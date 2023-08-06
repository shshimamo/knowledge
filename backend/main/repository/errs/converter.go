package errs

import "database/sql"

type RecordNotFoundError struct {
	Message string
}

func NewRecordNotFoundError() *RecordNotFoundError {
	return &RecordNotFoundError{
		Message: "Record not found",
	}
}

func (e *RecordNotFoundError) Error() string {
	return e.Message
}

func ConvertSqlError(err error) error {
	if err == nil {
		return nil
	}
	if err == sql.ErrNoRows {
		return NewRecordNotFoundError()
	}
	return err
}
